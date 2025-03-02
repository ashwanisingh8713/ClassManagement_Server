package database

import (
	"context"
	"fmt"
	"time"
)

// User represents the Users table
type User struct {
	UserID       int       `json:"user_id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// IsUserExist Check if user exists by email
func IsUserExist(email string) (bool, error) {
	var isUserNotExist bool = false
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE email=$1)", TableUsers)
	err := db.QueryRow(context.Background(), query, email).Scan(&isUserNotExist)
	if err != nil {
		return !isUserNotExist, err
	}
	return isUserNotExist, nil
}

// Create User
func CreateUser(username, passwordHash, role, email string) (bool, error) {
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash, role, email) VALUES ($1, $2, $3, $4)", TableUsers)
	_, err := db.Exec(context.Background(), query, username, passwordHash, role, email)
	return err == nil, err
}

// CreateUser inserts a new user into the database
func CreateUserByStruct(user User) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (username, password_hash, role, email)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id, created_at, updated_at`, TableUsers)
	return db.QueryRow(context.Background(), query, user.Username, user.PasswordHash, user.Role, user.Email).
		Scan(&user.UserID, &user.CreatedAt, &user.UpdatedAt)
}

// GetUserByEmailAndPasswordHash retrieves a user by email and password hash
func GetUserByEmailAndPasswordHash(email, passwordHash string) (User, error) {
	var user User
	query := fmt.Sprintf(`
		SELECT user_id, username, role, email, created_at, updated_at
		FROM %s
		WHERE email = $1 AND password_hash = $2`, TableUsers)
	err := db.QueryRow(context.Background(), query, email, passwordHash).
		Scan(&user.UserID, &user.Username, &user.Role, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return User{}, fmt.Errorf("error retrieving user: %v", err)
	}
	return user, nil
}

func GetUserByID(userID int) (string, string, string, string, error) {
	query := fmt.Sprintf(`SELECT username, password_hash, role, email FROM %s WHERE user_id = $1`, TableUsers)
	row := db.QueryRow(context.Background(), query, userID)
	var username, passwordHash, role, email string
	err := row.Scan(&username, &passwordHash, &role, &email)
	return username, passwordHash, role, email, err
}

func UpdateUserField(userID int, username, role, email string) error {
	query := fmt.Sprintf(`UPDATE %s SET username = $1, role = $2, email = $3, updated_at = CURRENT_TIMESTAMP WHERE user_id = $4`, TableUsers)
	_, err := db.Exec(context.Background(), query, username, role, email, userID)
	return err
}

// GetUsers that queries the database
func GetUsers() map[int]interface{} {
	var users = make(map[int]interface{}, 0)
	query := fmt.Sprintf("SELECT id, name FROM %s", TableUsers)
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return users
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return users
		}
		users[id] = name
	}
	return users
}

////////////////////////////////
// ====================== Users Table CRUD Operations ======================

// GetUser retrieves a user by ID
func GetUser(userID int) (User, error) {
	var user User
	query := fmt.Sprintf(`
		SELECT user_id, username, password_hash, role, email, created_at, updated_at
		FROM %s
		WHERE user_id = $1`, TableUsers)
	err := db.QueryRow(context.Background(), query, userID).
		Scan(&user.UserID, &user.Username, &user.PasswordHash, &user.Role, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

// UpdateUser updates an existing user
func UpdateUser(user User) error {
	query := fmt.Sprintf(`
		UPDATE %s
		SET username = $1, password_hash = $2, role = $3, email = $4, updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $5`, TableUsers)
	_, err := db.Exec(context.Background(), query, user.Username, user.PasswordHash, user.Role, user.Email, user.UserID)
	return err
}

// DeleteUser deletes a user by ID
func DeleteUser(userID int) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE user_id = $1`, TableUsers)
	_, err := db.Exec(context.Background(), query, userID)
	return err
}
