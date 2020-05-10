package grammar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestBcrypt(t *testing.T) {

	t.Run("helloWorld for bcrypt", func(t *testing.T) {
		hash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

		if err != nil {
			fmt.Printf("error is %v", err)
		}

		fmt.Printf("after bcrypt encode is %v", string(hash))
	})

	t.Run("bcrypt different const", func(t *testing.T) {
		hash1, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		hash2, _ := bcrypt.GenerateFromPassword([]byte("password"), 11)

		fmt.Printf("hash1 is %v", string(hash1))
		fmt.Printf("hash2 is %v", string(hash2))

		assert.Nil(t, bcrypt.CompareHashAndPassword(hash1, []byte("password")))
	})
}
