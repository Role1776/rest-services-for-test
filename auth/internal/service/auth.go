package service

import (
	"app"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt      = "20876BHJKLFSDF9789786b0-346rnklwd=1=3-sejor90234"
	signature = "2mHB87THI900ud02we78HJ9Y9D078e=3-sejor9fesdf0234"
)

type NewAuthorization struct{}

func (s *NewAuthorization) CreateUser(user app.User) (app.User, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	return user, nil
}

func (s *NewAuthorization) ParseToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(signature), nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed to parse token: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("failed to get claims")
	}

	id, ok := claims["user"].(float64)
	if !ok {
		return 0, fmt.Errorf("user id not found in token")
	}
	
	return int(id), nil
}

func (s *NewAuthorization) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *NewAuthorization) GenerateJWT(id int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(signature))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
