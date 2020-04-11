package first_app

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	t.Run("test how to sign a jwt", func(t *testing.T) {
		keyData, _ := ioutil.ReadFile("aws-lambda/rsa/rsa")

		key, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(keyData, "test")
		token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
			"foo": "bar",
			"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		})

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(key)
		fmt.Println(tokenString, err)
	})
}
