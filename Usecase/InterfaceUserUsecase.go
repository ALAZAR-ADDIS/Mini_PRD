package Usecase

import "Mini_PRD/Domain"

type IUserRepo interface {
	CreateUser(user Domain.User) error
	FindUserByEmail(email string) (*Domain.User, error)
	FindUserByID(id string) (*Domain.User, error)
	UpdateUserTokens(userID string, tokens []Domain.Token) error
	DeleteUserTokens(userID string) error
	FindUserByToken(token string) (*Domain.User, error)
}

type IPasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type IJWTService interface {
	GenerateTokens(userID string, email string, role string) (*Domain.Token, error)
	ValidateToken(tokenString string) (string, error)
	RefreshToken(refreshToken string) (*Domain.Token, error)
}
