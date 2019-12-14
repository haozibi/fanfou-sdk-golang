package oauth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestToken(t *testing.T) {

	const (
		consumerkey             = "test_consumer_key"
		consumerSecret          = "test_consumer_secret"
		requestOAuthToken       = "8ldIZyxQeVrFZXFOZH5tAwj6vzJYuLQpl0WUEYtWc"
		requestOAuthTokenSecret = "x6qpRnlEmW9JbQn4PQVVeVG8ZLPEx6A0TOebgwcuA"
		accessToken             = "891212-3MdIZyxPeVsFZXFOZj5tAwj6vzJYuLQplzWUmYtWd"
		accessTokenSecret       = "x6qpRnlEmW9JbQn4PQVVeVG8ZLPEx6A0TOebgwcuA"
		requestTokenURI         = "oauth/request_token"
		authorizeTokenURI       = "oauth/authorize"
		accessTokenURI          = "oauth/access_token"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := strings.TrimPrefix(r.URL.Path, "/")

		switch path {
		case requestTokenURI:
			fmt.Fprintf(w, "oauth_token=%s&oauth_token_secret=%s", requestOAuthToken, requestOAuthTokenSecret)
		case authorizeTokenURI:
		case accessTokenURI:
			fmt.Fprintf(w, "oauth_token=%s&oauth_token_secret=%s", accessToken, accessTokenSecret)
		}

	}))
	defer ts.Close()
	AuthBaseURL := ts.URL + "/"

	oa := NewOAuthWithClient(
		consumerkey,
		consumerSecret,
		ServiceProvider{
			RequestTokenURL:   AuthBaseURL + requestTokenURI,
			AuthorizeTokenURL: AuthBaseURL + authorizeTokenURI,
			AccessTokenURL:    AuthBaseURL + accessTokenURI,
			CallBackURL:       "",
		},
		ts.Client(),
	)

	token, err := oa.GetRequestToken(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
	}

	assert.Equal(t, requestOAuthToken, token.Token)
	assert.Equal(t, requestOAuthTokenSecret, token.Secret)

	uri, err := oa.GetAuthorizationURL(token)
	if err != nil {
		t.Errorf("%+v", err)
	}

	t.Log("Jump URL:", uri)

	acToken, err := oa.GetAccessToken(context.Background(),
		RequestToken{
			Token:  token.Token,
			Secret: token.Secret,
		}, "", nil)

	assert.Equal(t, accessToken, acToken.Token)
	assert.Equal(t, accessTokenSecret, acToken.Secret)

	// TODO: more auth test

	// build new client
	cc := oa.BuildClient(acToken)

	cc.Get(ts.URL)
	cc.Post(ts.URL, "application/x-www-form-urlencoded", strings.NewReader("a=1"))
}
