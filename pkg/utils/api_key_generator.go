package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/google/uuid"
)

// Function to generate a modified API key
func GenerateAPIKey() string {
	// Generate a UUID
	id := uuid.New()

	// Convert the UUID to a string
	uuidStr := id.String()

	// Hash the UUID using MD5
	hash := md5.Sum([]byte(uuidStr))

	// Convert the hash to a hexadecimal string
	hashStr := hex.EncodeToString(hash[:])

	// Join the UUID and hash strings
	apiKey := strings.Join([]string{uuidStr, hashStr}, "-")

	return apiKey
}
