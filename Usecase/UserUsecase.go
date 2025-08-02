package Usecase

import (
	"Mini_PRD/Domain"
	"errors"
	"time"
)

type UserUsecase struct {
	UserRepo        IUserRepo
	PasswordService IPasswordService
	JWTService      IJWTService
}

func NewUserUscase(userRepo IUserRepo, passwordService IPasswordService, jwtService IJWTService) *UserUsecase {
	return &UserUsecase{
		UserRepo:        userRepo,
		PasswordService: passwordService,
		JWTService:      jwtService,
	}
}

// Signup handles user registration
func (u *UserUsecase) Signup(request Domain.SignupRequest) (*Domain.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := u.UserRepo.FindUserByEmail(request.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := u.PasswordService.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	// Create new user
	user := Domain.User{
		Email:       request.Email,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		Password:    hashedPassword,
		Role:        "user", // Default role
		Tokens:      []Domain.Token{},
		UserProfile: Domain.Profile{},
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Save user to database
	err = u.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	// Generate tokens
	tokens, err := u.JWTService.GenerateTokens(user.Id.Hex(), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	// Update user with tokens
	err = u.UserRepo.UpdateUserTokens(user.Id.Hex(), []Domain.Token{*tokens})
	if err != nil {
		return nil, err
	}

	// Return response
	user.Password = "" // Don't return password
	return &Domain.AuthResponse{
		User:         user,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Message:      "User registered successfully",
	}, nil
}

// Signin handles user login
func (u *UserUsecase) Signin(request Domain.SigninRequest) (*Domain.AuthResponse, error) {
	// Find user by email
	user, err := u.UserRepo.FindUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	err = u.PasswordService.ComparePassword(user.Password, request.Password)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate new tokens
	tokens, err := u.JWTService.GenerateTokens(user.Id.Hex(), user.Email, user.Role)
	if err != nil {
		return nil, err
	}

	// Update user tokens
	err = u.UserRepo.UpdateUserTokens(user.Id.Hex(), []Domain.Token{*tokens})
	if err != nil {
		return nil, err
	}

	// Return response
	user.Password = "" // Don't return password
	return &Domain.AuthResponse{
		User:         *user,
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		Message:      "Login successful",
	}, nil
}

// Signout handles user logout
func (u *UserUsecase) Signout(request Domain.SignoutRequest) (*Domain.SignoutResponse, error) {
	// Validate token and get user ID
	userID, err := u.JWTService.ValidateToken(request.AccessToken)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Find user to ensure they exist
	_, err = u.UserRepo.FindUserByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Delete all tokens for the user
	err = u.UserRepo.DeleteUserTokens(userID)
	if err != nil {
		return nil, err
	}

	return &Domain.SignoutResponse{
		Message: "Logout successful",
	}, nil
}
