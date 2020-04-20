package oauth2Server

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"testing"
)

const (
	authServerURL = "http://localhost:9096"
)

func TestClient(t *testing.T) {

	ctx := context.Background()

	t.Run("login with password", func(t *testing.T) {
		conf := &oauth2.Config{
			ClientID:     "222222",
			ClientSecret: "22222222",
			RedirectURL:  "REDIRECT_URL",
			Scopes:       []string{"all"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  authServerURL + "/auth",
				TokenURL: authServerURL + "/token",
			},
		}
		token, err := conf.PasswordCredentialsToken(ctx, "hello", "hello")

		if err != nil {
			t.Logf("err is %v", err)
			return
		}
		t.Logf("refresh token is %v", token.RefreshToken)
		t.Logf("access token is %v", token.AccessToken)
		t.Logf("token type is %v", token.TokenType)

		segment, err := jwt.DecodeSegment(token.AccessToken)

		claims := jwt.MapClaims{}
		jwtToken, err := jwt.ParseWithClaims(token.AccessToken, claims, func(token *jwt.Token) (interface{}, error) {
			t.Logf("token is function method is %v", token)
			return []byte("<YOUR VERIFICATION KEY>"), nil
		})
		t.Logf("segment is %v", string(segment))

		t.Logf("token is %v", jwtToken)
	})
}
