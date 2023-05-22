package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// Load dotenv
func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Validate if jwt secret key in .env
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("No JWT_SECRET_KEY provided in .env file")
	}
}

// GenerateToken generates a JWT token with the given claims.
func GenerateToken(subject string, expiration time.Duration) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = subject
	claims["exp"] = time.Now().Add(expiration).Unix()

	// Sign the token with the secret key
	secretKey := os.Getenv("JWT_SECRET_KEY")

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}

		secretKey := os.Getenv("JWT_SECRET_KEY")

		// Return the secret key used for signing
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	// Check the token expiration
	expiration := claims["exp"].(float64)
	expirationTime := time.Unix(int64(expiration), 0)
	if time.Now().After(expirationTime) {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
