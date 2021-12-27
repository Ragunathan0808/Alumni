package models

type AuthService interface {
	LogIn(*User) (string, error)
	SignUp(*User) error
	RemoveAccount(*User) error

	UpdateActive(*User, bool) error
}

type TokenService interface {
	GenerateToken(*User) (string, error)
	DecodeToken(string) (*User, error)
}

type UserService interface {
	GetByID(uint) (*User, error)
	GetByEmail(string) (*User, error)
	GetByDesignation(string) ([]*User, error)
	GetByBatch(uint) ([]*User, error)
	GetUser(*User) ([]*User, error)
}

type EmailService interface {
	Send(string, string, string, string) error
}

type OTPService interface {
	GenerateOTP(*User) error
	VerifyOTP(string) error
}
