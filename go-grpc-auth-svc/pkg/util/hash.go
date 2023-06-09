package util

import (
	"golang.org/x/crypto/bcrypt"
)


func HashPassowrd(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	
	return string(bytes)
}

func CheckPassowrdHash(hashed, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plainText))
	return err == nil 
}