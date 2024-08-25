package service

import "golang.org/x/crypto/bcrypt"

func Encrypt(str string) (string, error) {
	bytes := []byte(str)

	bytes, err := bcrypt.GenerateFromPassword(bytes, 10)
	if err != nil {
		return "", err
	}

	str = string(bytes)

	return str, nil
}

func Compare(str string, hash string) bool {
	strBytes := []byte(str)
	hashBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(hashBytes, strBytes)
	return err == nil
}
