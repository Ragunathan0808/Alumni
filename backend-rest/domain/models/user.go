package models

type UserRepo interface {
	Create(*User) (error, *User)

	DeleteUserByID(uint) error

	FindByID(uint) (error, *User)
	FindByMail(uint) (error, *User)
	FindByDesignation(uint) (error, []*User)
	FindByBatch(uint) (error, []*User)

	UpdatePassword(uint, string) error
	UpdateVerify(uint, bool) error
	UpdateActive(uint, bool) error
	UpdateName(uint, string) error
}
