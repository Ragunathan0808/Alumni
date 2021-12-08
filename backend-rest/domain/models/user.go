package models

type UserRepo interface {
	Create(*User) (*User, error)

	DeleteByID(uint) error

	FindByID(uint) (*User, error)
	FindByMail(string) (*User, error)
	FindByDesignation(string) ([]*User, error)
	FindByBatch(uint) ([]*User, error)

	UpdatePassword(uint, string) error
	UpdateVerify(uint, bool) error
	UpdateActive(uint, bool) error
	UpdateName(uint, string) error
}
