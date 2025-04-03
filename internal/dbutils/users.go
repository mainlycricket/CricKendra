package dbutils

import (
	"context"

	"github.com/mainlycricket/CricKendra/internal/models"
)

// doesn't hash the password
func InsertUser(ctx context.Context, db DB_Exec, user *models.User) (int64, error) {
	var userId int64 = -1

	query := `INSERT INTO USERS (name, email, password, role) VALUES($1, $2, $3, $4) RETURNING id`

	err := db.QueryRow(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&userId)

	return userId, err
}

// doesn't hash the password
func UpsertUser(ctx context.Context, db DB_Exec, user *models.User) (int64, error) {
	var userId int64 = -1

	query := `
		INSERT INTO USERS (name, email, password, role) VALUES($1, $2, $3, $4) 
		ON CONFLICT (email)
		DO UPDATE SET name = $1, email = $2, password = $3, role = $4
		RETURNING id`

	err := db.QueryRow(ctx, query, user.Name, user.Email, user.Password, user.Role).Scan(&userId)

	return userId, err
}

type loginOutput struct {
	UserId         uint
	HashedPassword string
	Role           string
}

func LoginUser(ctx context.Context, db DB_Exec, email string) (loginOutput, error) {
	var output loginOutput

	query := `SELECT id, password, role FROM users WHERE email = $1`

	err := db.QueryRow(ctx, query, email).Scan(&output.UserId, &output.HashedPassword, &output.Role)

	return output, err
}
