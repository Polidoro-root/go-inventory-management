package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
	"github.com/Polidoro-root/go-inventory-management/internal/usecase"
)

type CreateUserUseCaseTestSuite struct {
}

func TestGivenAValidUserInput_WhenIExecuteCreateUser_ThenIShouldReceiveTheCreatedUser(t *testing.T) {

	repository := &userRepositoryMock{
		saveFn: func(ctx context.Context, user *entity.User) error {
			return nil
		},
	}

	input := usecase.CreateUserInputDTO{
		AdminID:     uuid.New().String(),
		Name:        "John Doe",
		Role:        "technician",
		Email:       "john@doe.com",
		PhoneNumber: "5515999999999",
		Password:    "password",
	}

	expectedUser := entity.User{
		ID:          uuid.New().String(),
		Name:        "John Doe",
		Role:        "technician",
		Email:       "john@doe.com",
		PhoneNumber: "5515999999999",
	}

	uc := usecase.NewCreateUserUseCase(repository)

	output, err := uc.Execute(input)

	table := []struct {
		input, expected string
	}{
		{expectedUser.Name, output.Name},
		{expectedUser.Email, output.Email},
		{expectedUser.PhoneNumber, output.PhoneNumber},
		{string(expectedUser.Role), output.Role},
	}

	if err != nil {
		t.Error(err)
	}

	if output.ID == "" {
		t.Error("user must have an ID")
	}

	for _, item := range table {
		if item.input != item.expected {
			t.Errorf("Expected %s but got %s", item.expected, item.input)
		}
	}

}
