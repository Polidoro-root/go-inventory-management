package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGivenAValidUser_WhenICallNewUser_ThenIShouldReceiveANewUser(t *testing.T) {
	user, err := NewUser(
		"1",
		"John Doe",
		"technician",
		"email@mail.com",
		"5515999999999",
		"password",
		time.Now(),
		time.Now(),
	)

	assert.Nil(t, err)

	assert.Equal(t, user.ID, "1")
	assert.Equal(t, user.Name, "John Doe")
	assert.Equal(t, user.Role, Role("technician"))
	assert.Equal(t, user.Email, "email@mail.com")
	assert.Equal(t, user.PhoneNumber, "5515999999999")
	assert.Equal(t, user.Password, "password")
	assert.IsType(t, user.CreatedAt, time.Time{})
	assert.IsType(t, user.UpdatedAt, time.Time{})
}

func TestGivenAnInvalidRole_WhenICallNewUser_ThenIShouldReceiveAnInvalidRoleError(t *testing.T) {
	user, err := NewUser(
		"1",
		"Name",
		"supplier",
		"email@mail.com",
		"5515999999999",
		"password",
		time.Now(),
		time.Now(),
	)

	assert.Nil(t, user)
	assert.EqualError(t, err, "role 'supplier' is invalid")
}

func TestGivenAnInvalidPhoneNumber_WhenICallNewUser_ThenIShouldReceiveAnInvalidPhoneNumberError(t *testing.T) {
	user, err := NewUser(
		"1",
		"Name",
		"guest",
		"email@mail.com",
		"15999999999",
		"password",
		time.Now(),
		time.Now(),
	)

	assert.Nil(t, user)
	assert.EqualError(t, err, "user's phone number must have 13 digits")

}
