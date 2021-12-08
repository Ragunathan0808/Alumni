package models

type UserRepo interface {
	Create(*User) (*User, error)

	DeleteUserByID(uint) error

	FindByID(uint) (*User, error)
	FindByMail(uint) (*User, error)
	FindByDesignation(uint) ([]*User, error)
	FindByBatch(uint) ([]*User, error)

	UpdatePassword(uint, string) error
	UpdateVerify(uint, bool) error
	UpdateActive(uint, bool) error
	UpdateName(uint, string) error
}
