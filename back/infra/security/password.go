package security

import (
	"devport/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *model.RawPassword) *model.HashedPassword {
	res, err := bcrypt.GenerateFromPassword([]byte(password.RawPassword()), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return model.NewHashedPassword(string(res))
}

func CheckPasswordHash(password *model.RawPassword, hash *model.HashedPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash.HashedPassword()), []byte(password.RawPassword()))

	return err == nil
}
