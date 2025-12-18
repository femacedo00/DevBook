package security

import "golang.org/x/crypto/bcrypt"

// Hash receives a string and returns a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// ValidatePassword checks if the propvided passoword matches the given hash
func ValidatePassword(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}
