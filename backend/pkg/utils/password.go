package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// NormalizePassword func for a returning the users input as a byte slice.
func NormalizePassword(p string) []byte {
	return []byte(p)
}

// GeneratePassword func for a making hash & salt with user password.
func GeneratePassword(p string) []byte {
	bytePwd := NormalizePassword(p)

	hash, err := bcrypt.GenerateFromPassword(bytePwd, bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return hash
}

// ComparePasswords func for a comparing password.
func ComparePasswords(hashedPwd []byte, inputPwd string) bool {
	byteInput := NormalizePassword(inputPwd)

	// Return result.
	if err := bcrypt.CompareHashAndPassword(hashedPwd, byteInput); err != nil {
		return false
	}

	return true
}
