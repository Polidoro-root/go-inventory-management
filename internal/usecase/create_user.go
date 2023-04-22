package usecase

import (
	"context"
	"time"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
	"github.com/google/uuid"
)

type CreateUserInputDTO struct {
	AdminID     string `json:"admin_id"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type CreateUserOutputDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

type CreateUserUseCase struct {
	UserRepository entity.UserRepositoryInterface
}

func NewCreateUserUseCase(
	userRepository entity.UserRepositoryInterface,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		UserRepository: userRepository,
	}
}

func (c *CreateUserUseCase) Execute(input CreateUserInputDTO) (*CreateUserOutputDTO, error) {
	ctx := context.Background()

	newUser, err := entity.NewUser(
		uuid.New().String(),
		input.Name,
		entity.Role(input.Role),
		input.Email,
		input.PhoneNumber,
		input.Password,
	)

	if err != nil {
		return nil, err
	}

	err = c.UserRepository.Save(ctx, newUser)

	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDTO{
		ID:          newUser.ID,
		Name:        newUser.Name,
		Role:        string(newUser.Role),
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,
		CreatedAt:   newUser.CreatedAt,
	}, nil
}
