package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/Girbons/effective-potato/pkg/config"
	log "github.com/sirupsen/logrus"
)

func GetToken(username string) (string, error) {
	conf, err := config.GetConf()

	if err != nil {
		log.Error(err)
	}

	signingKey := []byte(conf.JWTSigningKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * conf.JWTExpiration).Unix(),
	})

	tokenString, err := token.SignedString(signingKey)
	return tokenString, err
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	conf, err := config.GetConf()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	signingKey := []byte(conf.JWTSigningKey)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return token.Claims, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
