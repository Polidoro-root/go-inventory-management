package usecase

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	entity "github.com/Polidoro-root/go-inventory-management/internal/entity/user"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Save(user *entity.User) (entity.User, error) {
	m.Called(user)

	newUser, err := entity.NewUser(
		user.ID,
		user.Name,
		user.Role,
		user.Email,
		user.PhoneNumber,
		user.Password,
		user.CreatedAt,
		time.Time{},
	)

	return *newUser, err
}

func (m *UserRepositoryMock) FindAll() ([]entity.User, error) {
	return make([]entity.User, 0), nil
}

func (m *UserRepositoryMock) FindByID(id string) (*entity.User, error) {
	return &entity.User{}, nil
}

type CreateUserUseCaseTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *CreateUserUseCaseTestSuite) SetupSuite() {

}

func TestGivenAValidUserInput_WhenIExecuteCreateUser_ThenIShouldReceiveTheCreatedUser(t *testing.T) {

	repository := &UserRepositoryMock{}

	input := CreateUserInputDTO{
		AdminID:     uuid.New().String(),
		Name:        "John Doe",
		Role:        "technician",
		Email:       "john@doe.com",
		PhoneNumber: "5515999999999",
		Password:    "password",
	}

	expectedOutput := CreateUserOutputDTO{

		Name:        "John Doe",
		Role:        "technician",
		Email:       "john@doe.com",
		PhoneNumber: "5515999999999",
	}

	uc := NewCreateUserUseCase(repository)

	repository.On("Save", mock.Anything).Return(expectedOutput)

	output, err := uc.Execute(input)

	repository.AssertNumberOfCalls(t, "Save", 1)
	repository.AssertExpectations(t)

	assert.Nil(t, err)

	assert.NotEmpty(t, output.ID)
	assert.Equal(t, expectedOutput.Name, output.Name)
	assert.Equal(t, expectedOutput.Email, output.Email)
	assert.Equal(t, expectedOutput.PhoneNumber, output.PhoneNumber)
	assert.Equal(t, expectedOutput.Role, output.Role)
	assert.NotNil(t, output.CreatedAt)

}
