package database

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
)

type UserRepository struct {
	*Queries
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		Queries: New(db),
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	exist, err := r.Queries.UserEmailExist(
		ctx,
		user.Email,
	)

	if err != nil {
		return err
	}

	if exist {
		return errors.New("email already exists")
	}

	exist, err = r.Queries.UserPhoneNumberExist(
		ctx,
		user.PhoneNumber,
	)

	if err != nil {
		return err
	}

	if exist {
		return errors.New("phone number already exists")
	}

	return r.Queries.SaveUser(ctx, SaveUserParams{
		ID:          user.ID,
		Name:        user.Name,
		Role:        string(user.Role),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	})
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	user, err := r.Queries.FindUserByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:          user.ID,
		Name:        user.Name,
		Role:        entity.Role(user.Role),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := r.Queries.FindUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:          user.ID,
		Name:        user.Name,
		Role:        entity.Role(user.Role),
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    user.Password,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}, nil
}
