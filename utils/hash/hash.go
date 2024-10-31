package crypto

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

var (
	secret = viper.GetString("jwt_secret")
)

//HashText generates hased password
func HashText(password string) string {
	salt := 14
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes)
}

//CheckTextHash compares hash and user plain password
func CheckTextHash(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	if err != nil {
		return err == nil
	}
	return true
}

//Jwt generates json web tokens for stateless authentication
func Jwt(claim interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": claim,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString
}
