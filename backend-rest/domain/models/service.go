package models

type AuthService interface {
	LogIn(*User) (string, error)
	SignUp(*User) error
	RemoveAccount(*User) error
}

type TokenService interface {
	GenerateToken(*User) (string, error)
	DecodeToken(string) (*User, error)
}
