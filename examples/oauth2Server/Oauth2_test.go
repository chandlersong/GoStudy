package oauth2Server

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v3"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
	"io/ioutil"
	"testing"
)

func TestOauth2Server(t *testing.T) {

	t.Run("simplest example", func(t *testing.T) {

		manager := manage.NewDefaultManager()
		manager.MustTokenStorage(store.NewMemoryTokenStore())

		// client memory store
		clientStore := store.NewClientStore()
		clientId := "000000"
		_ = clientStore.Set(clientId, &models.Client{
			ID:     clientId,
			Secret: "999999",
			Domain: "http://localhost",
		})
		manager.MapClientStorage(clientStore)
		client, _ := manager.GetClient(clientId)
		t.Logf("client informaiton is: %v", client)

		tgr := &oauth2.TokenGenerateRequest{
			ClientID:    clientId,
			UserID:      "userId",
			Scope:       "all",
			RedirectURI: "http://localhost",
		}

		cti, _ := manager.GenerateAuthToken(oauth2.Token, tgr)

		t.Logf("clt is %v", cti)

		b, _ := json.Marshal(cti)

		t.Logf("json is %v", string(b))
	})

	t.Run("jwt token", func(t *testing.T) {
		keyData, err := ioutil.ReadFile("rsa/rsa")

		if err != nil {
			t.Fatal("not found key file")
		}

		key, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(keyData, "test")
		t.Logf("key is %v", key)

		manager := manage.NewDefaultManager()
		manager.MustTokenStorage(store.NewMemoryTokenStore())
		manager.MapAccessGenerate(generates.NewJWTAccessGenerate(key.D.Bytes(), jwt.SigningMethodHS512))

		clientId := "000000"
		clientStore := store.NewClientStore()
		_ = clientStore.Set(clientId, &models.Client{
			ID:     clientId,
			Secret: "999999",
			Domain: "http://localhost",
		})
		manager.MapClientStorage(clientStore)

		tgr := &oauth2.TokenGenerateRequest{
			ClientID:     clientId,
			ClientSecret: "999999",
			UserID:       "userId",
			Scope:        "all",
			RedirectURI:  "http://localhost",
		}

		cti, _ := manager.GenerateAuthToken(oauth2.Token, tgr)

		b, _ := json.Marshal(cti)

		t.Logf("json is %v", string(b))

		token, err := manager.GenerateAccessToken(oauth2.ClientCredentials, tgr)

		if err != nil {
			t.Fatal(err)
		}

		jwtToken, _ := json.Marshal(token)
		t.Logf("token is %v", string(jwtToken))

	})
}
