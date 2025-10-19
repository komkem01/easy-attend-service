package jwt

import (
	"context"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/komkem01/easy-attend-service/model"
)

// Claims structure for JWT tokens
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// RefreshClaims structure for refresh tokens
type RefreshClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// GetAccessTokenExpiry returns access token expiry duration in seconds
func GetAccessTokenExpiry() int64 {
	return 24 * 60 * 60 // 24 hours in seconds
}

// GetRefreshTokenExpiry returns refresh token expiry duration in seconds
func GetRefreshTokenExpiry() int64 {
	return 7 * 24 * 60 * 60 // 7 days in seconds
}

// GenerateToken creates a new access token
func GenerateToken(userID, email, role string) (string, error) {
	godotenv.Load()

	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(GetAccessTokenExpiry()) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	return token.SignedString(secret)
}

// GenerateRefreshToken creates a new refresh token
func GenerateRefreshToken(userID string) (string, error) {
	godotenv.Load()

	claims := &RefreshClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(GetRefreshTokenExpiry()) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(os.Getenv("JWT_SECRET"))
	return token.SignedString(secret)
}

// ValidateToken validates an access token and returns claims
func ValidateToken(tokenString string) (*Claims, error) {
	godotenv.Load()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// ValidateRefreshToken validates a refresh token and returns claims
func ValidateRefreshToken(tokenString string) (*RefreshClaims, error) {
	godotenv.Load()

	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		secret := []byte(os.Getenv("JWT_SECRET"))
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid refresh token")
}

// VerifyToken - legacy function for compatibility
func VerifyToken(raw string) (map[string]any, error) {
	claims, err := ValidateToken(raw)
	if err != nil {
		return nil, err
	}

	result := map[string]any{
		"user_id": claims.UserID,
		"email":   claims.Email,
		"role":    claims.Role,
		"exp":     claims.ExpiresAt.Unix(),
	}
	return result, nil
}

// GenerateTokenUser - legacy function for compatibility
func GenerateTokenUser(ctx context.Context, user *model.Users) (string, error) {
	return GenerateToken(user.ID.String(), user.Email, string(user.Role))
}
