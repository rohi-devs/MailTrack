package Utlis

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(str string) (hashed string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	hashed = string(hashedPassword)
	return hashed, err
}

func CompareHashAndPassword(hashedPwd string, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
