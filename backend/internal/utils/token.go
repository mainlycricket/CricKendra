package utils

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomJwtClaims struct {
	UserId uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func GetSignedToken(userId uint, role string) (string, error) {
	claims := CustomJwtClaims{
		userId, role, jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	return signedString, err
}

func NewTokenCookie(token string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "access_token",
		Value:    token,
		Expires:  time.Now().Add(30 * 24 * time.Hour), // expires in 30 days
		HttpOnly: true,
		SameSite: http.SameSiteDefaultMode,
		Path:     "/",
		Secure:   true,
	}

	return cookie
}

func ValidateSignedToken(signedToken string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(signedToken, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, errors.New("failed to parse claims")
}
