package security

import "golang.org/x/crypto/bcrypt"

//Encrypt use golang.org/x/crypto/bcrypt generate password
func Encrypt(src string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//ValidatePassword use golang.org/x/crypto/bcrypt compare passwords for equality
func ValidatePassword(plaintext, ciphertext string) bool {
	if len(ciphertext) <= 0 {
		return false
	}
	err := bcrypt.CompareHashAndPassword([]byte(ciphertext), []byte(plaintext))
	if err != nil {
		return false
	}
	return true
}
