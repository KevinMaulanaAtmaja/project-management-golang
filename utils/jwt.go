package utils

import (
	"time"

	"github.com/KevinMaulanaAtmaja/project-management-golang/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// generate token jwt
func GenerateToken(userID int64, role, email string, publicID uuid.UUID) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTExpire)

	claims := jwt.MapClaims{
		"user_id":   userID,
		"role":      role,
		"email":     email,
		"public_id": publicID,
		"exp":       time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// generate refresh token jwt
func GenerateRefreshToken(userID int64) (string, error) {
	secret := config.AppConfig.JWTSecret
	duration, _ := time.ParseDuration(config.AppConfig.JWTRefreshToken)

	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
