package entity

import "context"

type UserRepositoryInterface interface {
	Save(ctx context.Context, user *User) error

	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
