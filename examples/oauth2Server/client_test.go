package oauth2Server

import (
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
			ClientID:     "YOUR_CLIENT_ID",
			ClientSecret: "YOUR_CLIENT_SECRET",
			RedirectURL:  "REDIRECT_URL",
			Scopes:       []string{"SCOPE1", "SCOPE2"},
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
		t.Logf("token is %v", token)
	})
}
