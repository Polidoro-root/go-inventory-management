package usecase_test

import (
	"context"

	"github.com/Polidoro-root/go-inventory-management/internal/entity"
)

type userRepositoryMock struct {
	saveFn        func(ctx context.Context, user *entity.User) error
	findByIDFn    func(ctx context.Context, id string) (*entity.User, error)
	findByEmailFn func(ctx context.Context, email string) (*entity.User, error)
}

func (m *userRepositoryMock) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	if m != nil && m.findByEmailFn != nil {
		return m.findByEmailFn(ctx, email)
	}

	return nil, nil
}

func (m *userRepositoryMock) Save(ctx context.Context, user *entity.User) error {
	if m != nil && m.saveFn != nil {
		return m.saveFn(ctx, user)
	}

	return nil
}

func (m *userRepositoryMock) FindByID(ctx context.Context, id string) (*entity.User, error) {
	if m != nil && m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}
