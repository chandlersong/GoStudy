package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"gopkg.in/oauth2.v3/generates"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	authServerURL = "http://localhost:9096"
)

var (
	config = oauth2.Config{
		ClientID:     "222222",
		ClientSecret: "22222222",
		Scopes:       []string{"all"},
		RedirectURL:  "http://localhost:9095/oauth2",
		Endpoint: oauth2.Endpoint{
			AuthURL:  authServerURL + "/authorize",
			TokenURL: authServerURL + "/token",
		},
	}
)

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	// token store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// generate jwt access token
	keyData, err := ioutil.ReadFile("rsa/rsa")

	if err != nil {
		fmt.Errorf("not found key file %v", "rsa/rsa")
	}

	key, _ := jwt.ParseRSAPrivateKeyFromPEMWithPassword(keyData, "test")
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate(key.D.Bytes(), jwt.SigningMethodHS512))

	clientStore := store.NewClientStore()
	clientStore.Set("222222", &models.Client{
		ID:     "222222",
		Secret: "22222222",
		Domain: "http://localhost:9094",
	})
	manager.MapClientStorage(clientStore)

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "hello" && password == "hello" {
			userID = "hello"
		}
		return
	})
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("auth is %v", r)

		_ = r.ParseForm()
		password := r.Form.Get("password")
		fmt.Printf("password is %v", password)
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("token is %v", r)
		_ = r.ParseForm()
		password := r.Form.Get("password")
		fmt.Printf("password is %v", password)

		err := srv.HandleTokenRequest(w, r)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})
	log.Fatal(http.ListenAndServe(":9096", nil))
}
