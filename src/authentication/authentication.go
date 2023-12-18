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

func CreateToken(ID string, roleType string) (string, error) {
	permissions := jwt.MapClaims{}

	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = ID
	permissions["role"] = roleType

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signature method! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func ExtractUserID(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%s", permissions["userId"]), nil
	}

	return "", errors.New("invalid token")
}

func ExtractRoleFromToken(r *http.Request) (string, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return "", err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return fmt.Sprintf("%s", permissions["role"]), nil
	}

	return "", errors.New("Invalid token")
}
