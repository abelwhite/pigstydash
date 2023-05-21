// Filename: internal/models/users.go
//this file contains all our data and sql we need for rooms

// Written by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
// Tested by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
// Debbuged by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalezpackage models
package models

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid Credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

// Create a user
type User struct {
	UserID    int64
	Name      string
	Email     string
	Password  []byte
	Activated bool
	CreatedAt time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password, confirmpassword string) error {
	//lets first hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}
	query := `
			INSERT INTO users(name, email, password_hash)
			VALUES($1, $2, $3)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, name, email, hashedPassword)
	if err != nil {
		switch {
		case err.Error() == `Error: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
func (m *UserModel) Authenticate(email, password string) (int, error) {
	//compare
	var id int
	var hashedPassword []byte

	//check if there is a row in the table for the email provided
	query := `
		SELECT user_id, password_hash
		FROM users 
		WHERE email = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	} //handling error
	//the user does exist
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	//password is correct
	return id, nil
}

func (m *UserModel) Read() ([]*User, error) {
	//create SQL statement
	statement := `
		SELECT *
		FROM users
	`
	rows, err := m.DB.Query(statement)
	if err != nil {
		return nil, err
	}
	//cleanup before we leave our read method
	defer rows.Close()

	users := []*User{} //this will contain the pointer to all quotes

	for rows.Next() {
		u := &User{}
		err = rows.Scan(&u.UserID, &u.Name, &u.Email, &u.Password, &u.Activated, &u.CreatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, u) //contain first row
	}
	//check to see if there were error generated

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil

}

func (m *UserModel) GetByID(id int) (*User, error) {
	// Create SQL statement
	statement := `
		SELECT *
		FROM users
		WHERE user_id = $1
	`

	// Query the database
	row := m.DB.QueryRow(statement, id)

	// Create a new User instance
	u := &User{}

	// Scan the row data into the User instance
	err := row.Scan(&u.UserID, &u.Name, &u.Email, &u.Password, &u.Activated, &u.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	// Return the User instance and no error
	return u, nil
}

func (m *UserModel) UserData(UserID int) (string, string, error) {
	var name string
	var email string
	query := `
		SELECT name, email
		FROM users
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, UserID).Scan(&name, &email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", "", err
		}
	}
	//we got the user name and email
	return name, email, nil

}
