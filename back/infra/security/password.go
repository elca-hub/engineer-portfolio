package security

import (
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(res[:])
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
