package entity_test

import (
	"testing"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
)

func TestGivenAValidUser_WhenICallNewUser_ThenIShouldReceiveANewUser(t *testing.T) {
	password := "password"

	user, err := entity.NewUser(
		"1",
		"John Doe",
		"technician",
		"email@mail.com",
		"5515999999999",
		password,
	)

	tableEquals := []struct{ input, expected interface{} }{
		{user.ID, "1"},
		{user.Name, "John Doe"},
		{user.Role, entity.Role("technician")},
		{user.Email, "email@mail.com"},
		{user.PhoneNumber, "5515999999999"},
	}

	if err != nil {
		t.Error(err)
	}

	if password == user.Password {
		t.Error("Password must be encrypted")
	}

	for _, item := range tableEquals {
		if item.expected != item.input {
			t.Errorf("Expected %s but got %s", item.expected, item.input)
		}
	}

}

func TestGivenAValidUser_WhenICallValidatePasswordWithTheRightPassword_ThenIShouldNotReceiveAnInvalidPasswordError(t *testing.T) {
	user, err := entity.NewUser(
		"1",
		"Name",
		"technician",
		"email@mail.com",
		"5515999999999",
		"password",
	)

	if err != nil {
		t.Error(err)
	}

	if user == nil {
		t.Error("user should have been created")
	}

	if err := user.ValidatePassword("password"); err != nil {
		t.Error(err)
	}

}

func TestGivenAnInvalidRole_WhenICallNewUser_ThenIShouldReceiveAnInvalidRoleError(t *testing.T) {
	user, err := entity.NewUser(
		"1",
		"Name",
		"supplier",
		"email@mail.com",
		"5515999999999",
		"password",
	)

	if user != nil {
		t.Error("user should not be created")
	}

	if err.Error() != "role 'supplier' is invalid" {
		t.Fail()
	}

}

func TestGivenAnInvalidPhoneNumber_WhenICallNewUser_ThenIShouldReceiveAnInvalidPhoneNumberError(t *testing.T) {
	user, err := entity.NewUser(
		"1",
		"Name",
		"technician",
		"email@mail.com",
		"15999999999",
		"password",
	)

	if user != nil {
		t.Error("user should not be created")
	}

	if err.Error() != "user's phone number must have 13 digits" {
		t.Fail()
	}
}

func TestGivenAnInvalidPassword_WhenICallNewUser_ThenIShouldReceiveAnInvalidPasswordError(t *testing.T) {
	user, err := entity.NewUser(
		"1",
		"Name",
		"technician",
		"email@mail.com",
		"5515999999999",
		"",
	)

	if user != nil {
		t.Error("user should not be created")
	}

	if err.Error() != "user must have a password" {
		t.Fail()
	}
}
