package util

import (
	"fmt"
	"log"

	"github.com/aldy505/phc-crypto/argon2"
)

// HashPassword returns the bcrypt hash of the password
func HashPassword(password string) (string, error) {
	hashedPassword, err := argon2.Hash(password, argon2.Config{
		Parallelism: 3,
		Variant:     argon2.I,
	})
	if err != nil {
		return " ", fmt.Errorf("failed to hash password: %w", err)
	}
	log.Println(hashedPassword)
	return hashedPassword, nil
}

// CheckPassword checks if the provided password is correct or not
func CheckPassword(password string, hashedPassword string) (bool, error) {
	return argon2.Verify(hashedPassword, password)
}
