package usecase

import (
	"context"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
)

type UserSignInUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

type UserSignInInputDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInOutputDTO struct {
	UserID string `json:"user_id"`
}

func NewUserSignInUseCase(
	userRepository entity.UserRepositoryInterface,

) *UserSignInUseCase {
	return &UserSignInUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserSignInUseCase) Execute(input UserSignInInputDTO) (*UserSignInOutputDTO, error) {
	ctx := context.Background()

	user, err := uc.UserRepository.FindByEmail(ctx, input.Email)

	if err != nil {
		return nil, err
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		return nil, err
	}

	return &UserSignInOutputDTO{
		UserID: user.ID,
	}, nil
}
