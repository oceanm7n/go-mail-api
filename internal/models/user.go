package models

import (
	"github.com/google/uuid"

	"mailganer-task/internal/db"
)

// Create a user model

// Path: internal\models\user.go

// User model
type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate string    `json:"birth_date"`
	Email     string    `json:"email"`
}

// New user
func NewUser(firstName string, lastName string, birthDate string, email string) *User {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		Email:     email,
	}
}

// Get all users from the database

// Path: internal\models\user.go

// GetAllUsers gets all users from the database
func GetAllUsers() ([]User, error) {
	var users []User
	var user User

	rows, err := db.DB.Query("SELECT * FROM Users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.BirthDate, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) Save() error {
	_, err := db.DB.Exec("INSERT INTO Users (id, first_name, last_name, birth_date, email) VALUES ($1, $2, $3, $4, $5)", u.ID, u.FirstName, u.LastName, u.BirthDate, u.Email)
	if err != nil {
		return err
	}

	return nil
}
