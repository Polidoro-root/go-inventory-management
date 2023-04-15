package user

type UserRepositoryInterface interface {
	Save(user *User) (User, error)

	FindByID(id string) (*User, error)
	FindAll() ([]User, error)
}
