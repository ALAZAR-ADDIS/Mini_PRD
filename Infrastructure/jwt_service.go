package Infrastructure

import (
	"Mini_PRD/Domain"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	JWTSecret string
}

func NewJWTService() *JWTService {
	return &JWTService{
		JWTSecret: getJWTSecret(),
	}
}

func getJWTSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key-change-in-production"
	}
	return secret
}

func (s *JWTService) GenerateTokens(userID string, email string, role string) (*Domain.Token, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(15 * time.Minute).Unix(),
		"type":    "access",
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"role":    role,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"type":    "refresh",
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &Domain.Token{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresAt:    time.Now().Add(15 * time.Minute),
		CreatedAt:    time.Now(),
	}, nil
}

func (s *JWTService) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.JWTSecret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, exists := claims["type"]; !exists || tokenType != "access" {
			return "", errors.New("invalid token type")
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return "", errors.New("invalid user_id in token")
		}
		return userID, nil
	}

	return "", errors.New("invalid token")
}

func (s *JWTService) RefreshToken(refreshTokenString string) (*Domain.Token, error) {
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if tokenType, exists := claims["type"]; !exists || tokenType != "refresh" {
			return nil, errors.New("invalid token type")
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, errors.New("invalid user_id in token")
		}

		email, ok := claims["email"].(string)
		if !ok {
			return nil, errors.New("invalid email in token")
		}

		role, ok := claims["role"].(string)
		if !ok {
			return nil, errors.New("invalid role in token")
		}

		return s.GenerateTokens(userID, email, role)
	}

	return nil, errors.New("invalid refresh token")
}
