package util

import (
	"crypto/sha1"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(data []byte) string {
	h := sha1.New()
	h.Write(data)
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return sha1Hash
}

func HashPassword(password []byte) []byte {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		SLogger.Fatalf("unable to hash password: %v", err)
	}

	return hashedPassword
}

func PasswordHashMatch(hashedPassword []byte, password []byte) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		SLogger.Errorf("unable to compare password hashes: %v", err)
		return false
	}

	return true
}
