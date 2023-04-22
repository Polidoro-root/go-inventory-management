package usecase_test

import (
	"context"
	"testing"

	"github.com/google/uuid"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
	"github.com/Polidoro-root/go-inventory-management/internal/usecase"
)

func TestGivenAValidUserInput_WhenIExecuteUserSignIn_ThenIShouldReceive(t *testing.T) {

	user := &entity.User{
		ID:       uuid.NewString(),
		Email:    "email@mail.com",
		Password: "$2a$10$3KWResahBU6x3wZis9uHhO5kwPQQSwcr3XMreufXimubaEx.csyXO",
	}

	repository := &userRepositoryMock{
		findByEmailFn: func(ctx context.Context, email string) (*entity.User, error) {
			return user, nil
		},
	}

	uc := usecase.NewUserSignInUseCase(repository)

	input := usecase.UserSignInInputDTO{
		Email:    "email@mail.com",
		Password: "password",
	}

	output, err := uc.Execute(input)

	if err != nil {
		t.Error(err)
	}

	if user.ID != output.UserID {
		t.Errorf("Expected %s but got %s", output.UserID, user.ID)
	}

}
