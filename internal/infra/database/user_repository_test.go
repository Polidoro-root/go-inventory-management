package database_test

import (
	"context"
	"testing"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
	"github.com/Polidoro-root/go-inventory-management/internal/infra/database"
	"github.com/Polidoro-root/go-inventory-management/internal/testutils"
	"github.com/google/uuid"
)

type UserRepositoryTestSuite struct {
	Repository *database.UserRepository
}

func setupTest(t *testing.T) *database.UserRepository {
	db := testutils.SetupDatabase(t)

	return database.NewUserRepository(db)
}

func TestGivenAValidUserInput_WhenICallSave_ThenIShouldNotReceiveError(t *testing.T) {
	repository := setupTest(t)

	newUser, err := entity.NewUser(
		uuid.NewString(),
		"John Doe",
		"technician",
		"john@doe.com",
		"5515999999999",
		"password",
	)

	if err != nil {
		t.Fatal(err)
	}

	if newUser == nil {
		t.Fatal("user should not be nil")
	}

	err = repository.Save(context.Background(), newUser)

	if err != nil {
		t.Fatal(err)
	}

	t.Run("TestGivenAStoredUser_WhenICallSaveWithSameEmail_ThenIShouldReceiveEmailAlreadyExistsError", func(t *testing.T) {
		newUser, err = entity.NewUser(
			uuid.NewString(),
			"John Doe",
			"technician",
			"john@doe.com",
			"5515999999999",
			"password",
		)

		if err != nil {
			t.Fatal(err)
		}

		if newUser == nil {
			t.Fatal("user should not be nil")
		}

		err = repository.Save(context.Background(), newUser)

		if err == nil || err.Error() != "email already exists" {
			t.Error("err should be 'email already exists'")
		}

	})

	t.Run("TestGivenAStoredUser_WhenICallSaveWithSamePhoneNumber_ThenIShouldReceivePhoneNumberAlreadyExistsError", func(t *testing.T) {
		newUser, err = entity.NewUser(
			uuid.NewString(),
			"John Doe",
			"technician",
			"john1@doe.com",
			"5515999999999",
			"password",
		)
		if err != nil {
			t.Fatal(err)
		}

		if newUser == nil {
			t.Fatal("user should not be nil")
		}

		err = repository.Save(context.Background(), newUser)

		if err == nil || err.Error() != "phone number already exists" {
			t.Error("err should be 'phone number already exists'")
		}

	})
}

func TestGivenAStoredUser_WhenICallFindByID_ThenIShouldReceiveUser(t *testing.T) {
	repository := setupTest(t)

	newUser, err := entity.NewUser(
		uuid.NewString(),
		"John Doe",
		"technician",
		"john1@doe.com",
		"5515199999999",
		"password",
	)

	if err != nil {
		t.Fatal(err)
	}

	if newUser == nil {
		t.Fatal("user should not be nil")

	}

	err = repository.Save(context.Background(), newUser)

	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.FindByID(context.Background(), newUser.ID)

	if err != nil {
		t.Fatal(err)
	}

	if user == nil {
		t.Fatal("user should not be nil")
	}

	if newUser.ID != user.ID {
		t.Errorf("Expected ID to be %s but got %s", newUser.ID, user.ID)
	}

}

func TestGivenAStoredUser_WhenICallFindByEmail_ThenIShouldReceiveUser(t *testing.T) {
	repository := setupTest(t)

	newUser, err := entity.NewUser(
		uuid.NewString(),
		"John Doe",
		"technician",
		"john2@doe.com",
		"5515299999999",
		"password",
	)

	if err != nil {
		t.Fatal(err)
	}

	err = repository.Save(context.Background(), newUser)

	if err != nil {
		t.Fatal(err)
	}

	user, err := repository.FindByEmail(context.Background(), newUser.Email)

	if err != nil {
		t.Fatal(err)

	}

	if user == nil {
		t.Fatal("user should not be nil")
	}

	if newUser.Email != user.Email {
		t.Errorf("Expected Email to be %s but got %s", newUser.Email, user.Email)
	}
}
