package database

import (
	"database/sql"

	entity "github.com/Polidoro-root/go-inventory-management/internal/entity/user"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (r *UserRepository) Save(user *entity.User) error {
	return nil
}
