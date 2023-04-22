package entity

import (
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Role string

const (
	Salesman   Role = "salesman"
	Technician Role = "technician"
	Admin      Role = "admin"
)

type User struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Role        Role      `json:"role"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewUser(
	id string,
	name string,
	role Role,
	email string,
	phoneNumber string,
	password string,
) (*User, error) {

	err := validateRole(role)

	if err != nil {
		return nil, err
	}

	err = validatePhoneNumber(phoneNumber)

	if err != nil {
		return nil, err
	}

	if password == "" {
		return nil, errors.New("user must have a password")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:          id,
		Name:        name,
		Role:        role,
		Email:       email,
		PhoneNumber: phoneNumber,
		Password:    string(hash),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}, nil
}

func (u *User) ValidatePassword(password string) error {

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func validateRole(role Role) error {
	isRoleValid :=
		role == Salesman ||
			role == Technician ||
			role == Admin

	if !isRoleValid {
		return fmt.Errorf("role '%s' is invalid", role)
	}

	return nil
}

func validatePhoneNumber(phoneNumber string) error {
	requiredLength := 13

	length := len(phoneNumber)

	if length != requiredLength {
		return errors.New("user's phone number must have 13 digits")
	}

	return nil
}
