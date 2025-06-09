package utils

import (
	"context"
	"errors"
	"net/http"
	"slices"

	"github.com/mainlycricket/CricKendra/backend/internal/responses"
)

type ContextKey string

func AuthorizationMiddleware(allowedRoles []string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtInfo, err := AuthorizeRequest(r, allowedRoles)
			if err != nil {
				responses.WriteJsonResponse(w, responses.ApiResponse{Success: false, Message: "unauthorized request", Data: err}, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ContextKey("token_info"), jwtInfo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func AuthorizeRequest(r *http.Request, allowedRoles []string) (CustomJwtClaims, error) {
	cookie, err := r.Cookie("access_token")
	if err != nil {
		return CustomJwtClaims{}, errors.New("access token not found")
	}

	claims, err := ValidateSignedToken(cookie.Value)
	if err != nil {
		return CustomJwtClaims{}, errors.New("invalid access token")
	}

	jwtInfo := CustomJwtClaims{
		Role:   claims["role"].(string),
		UserId: uint(claims["user_id"].(float64)),
	}

	if len(allowedRoles) > 0 && !slices.Contains(allowedRoles, jwtInfo.Role) {
		return CustomJwtClaims{}, errors.New("unauthorized role")
	}

	return jwtInfo, nil
}
