package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"time"

	"github.com/prathamesh-88/go-boilerplate/pkg/database"
)

type User struct {
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Password        string    `json:"-"`
	Role            string    `json:"role"`
	IsEmailVerified bool      `json:"is_email_verified"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func CreateUser(user *User) (User, error) {

	passwordHashBytes := sha256.Sum256([]byte(user.Password))
	hashedPassword := hex.EncodeToString(passwordHashBytes[:])

	rows, err := database.Db.Query(`
		INSERT INTO users(username, email, password, role, is_email_verified, created_at, updated_at) VALUES
		($1, $2, $3, $4, $5, NOW(), NOW());
		`,
		user.Username,
		user.Email,
		hashedPassword,
		user.Role,
		user.IsEmailVerified,
	)

	if err != nil {
		return User{}, err
	}

	log.Println("The user entry was created successfully: ", rows)

	return *user, nil
}

func GetUserByEmail(email string) (*sql.Rows, error) {
	rows, err := database.Db.Query(`
		SELECT * FROM users WHERE email = $1;
	`, email)

	if err != nil {
		log.Fatalf("Error while fetching users")
		return nil, err
	}

	log.Println(rows)

	return rows, nil
}
