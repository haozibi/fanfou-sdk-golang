package oauth

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestGetRequestToken(t *testing.T) {

	oa := NewOAuth(
		"",
		"",
		ServiceProvider{
			RequestTokenURL:   "https://fanfou.com/oauth/request_token",
			AuthorizeTokenURL: "https://fanfou.com/oauth/authorize",
			AccessTokenURL:    "https://fanfou.com/oauth/access_token",
			CallBackURL:       "",
		},
	)

	reqToken, err := oa.GetRequestToken(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(reqToken)

	u, err := oa.GetAuthorizationURL(reqToken)
	fmt.Println("+>", u, err)

	time.Sleep(7 * time.Second)

	acToken, err := oa.GetAccessToken(context.Background(), *reqToken, "", nil)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(acToken)
}
