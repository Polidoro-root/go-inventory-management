package usecase

import (
	"time"

	entity "github.com/Polidoro-root/go-inventory-management/internal/entity/user"
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

func (c *CreateUserUseCase) Execute(input CreateUserInputDTO) (CreateUserOutputDTO, error) {
	newUser, err := entity.NewUser(
		uuid.New().String(),
		input.Name,
		entity.Role(input.Role),
		input.Email,
		input.PhoneNumber,
		input.Password,
		time.Now(),
		time.Time{},
	)

	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	output, err := c.UserRepository.Save(newUser)

	if err != nil {
		return CreateUserOutputDTO{}, err
	}

	return CreateUserOutputDTO{
		ID:          output.ID,
		Name:        output.Name,
		Role:        string(output.Role),
		Email:       output.Email,
		PhoneNumber: output.PhoneNumber,
		CreatedAt:   output.CreatedAt,
	}, nil
}
