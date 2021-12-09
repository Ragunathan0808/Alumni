package models

type AuthService interface {
	LogIn(string, string) (string, error)
	SignUp(string, string) (string, error)
	RemoveAccount(uint) error
}

type TokenService interface {
	GenerateToken(*User) (string, error)
	DecodeToken(string) (*User, error)
}
