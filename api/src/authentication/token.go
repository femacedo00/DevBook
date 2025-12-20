package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CreateToken returns a token signed with the user permissions
func CreateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

// ValidateToken checks if the request token is validated
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, returnKeyVerification)
	if error != nil {
		return error
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}

// ExtractUserId returns the userID stored in token
func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, error := jwt.Parse(tokenString, returnKeyVerification)
	if error != nil {
		return 0, error
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, errors.New("Invalid Token")
	}

	userID, ok := permissions["userId"].(float64)
	if !ok {
		return 0, errors.New("Invalid Token")
	}

	return uint64(userID), nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	split := strings.Split(token, " ")

	if len(split) == 2 {
		return split[1]
	}

	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
