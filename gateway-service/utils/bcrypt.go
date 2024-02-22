package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// func HashPassword(pass *string) {
// 	bytePass := []byte(*pass)
// 	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
// 	*pass = string(hPass)
// }

func HashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass := sha256.Sum256(bytePass)
	datapass := fmt.Sprintf("%x", hPass)
	*pass = string(datapass)
}

// func ComparePassword(dbPass, pass string) bool {
// 	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
// }

func ComparePassword(dbPass, pass string) bool {
	verify := false

	if dbPass == pass {
		verify = true
	} else {
		verify = false
	}

	return verify

}

// GenerateToken -> generates token
func GenerateToken(username string) string {
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 3).Unix(),
		"iat":      time.Now().Unix(),
		"username": username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	//t, _ := token.SignedString([]byte("{*aKV|~ACD$Gm-Sk"))
	return t

}

func GenerateSignature(data string) string {
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
		"iat":  time.Now().Unix(),
		"data": data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t

}

// ValidateToken --> validate the given token
func ValidateToken(token string) (*jwt.Token, error) {

	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
		//return []byte("{*aKV|~ACD$Gm-Sk"), nil
	})
}

func ValidasiUser(username string) string {
	return username
}
