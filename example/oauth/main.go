package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/haozibi/fanfou-sdk-golang/fanfou"
	"github.com/haozibi/fanfou-sdk-golang/oauth"
)

var (
	consumerKey    string
	consumerSecret string
	port           int
)

func init() {
	tokens = make(map[string]*oauth.RequestToken)
	flag.StringVar(&consumerKey, "consumerKey", "", "consumer key from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&consumerSecret, "consumerSecret", "", "consumer secret from Fanfou. https://fanfou.com/apps")
	flag.IntVar(&port, "port", 8080, "http port")
	flag.Parse()
}

func main() {

	if consumerKey == "" ||
		consumerSecret == "" {
		flag.PrintDefaults()
		return
	}

	f := &fan{
		client: fanfou.NewFanFou(consumerKey, consumerSecret),
	}

	http.HandleFunc("/", f.Login)
	http.HandleFunc("/callback", f.GetAccessToken)

	log.Println("listen port:", port)
	log.Printf("You can visit http://localhost:%d to start the authorization process\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalln(err)
	}
}

var tokens map[string]*oauth.RequestToken

type fan struct {
	client *fanfou.Fanfou
}

func (f *fan) Login(w http.ResponseWriter, r *http.Request) {

	callback := fmt.Sprintf("http://%s/callback", r.Host)

	// 1.获取未授权的 Request Token
	token, err := f.client.RequestToken(context.Background(), callback)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%+v", err)
		return
	}

	tokens[token.Token] = token

	// 2.请求用户授权 Request Token
	uri, err := f.client.AuthorizationURL(token)
	if err != nil {
		fmt.Fprintf(w, "%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, uri, http.StatusTemporaryRedirect)
}

func (f *fan) GetAccessToken(w http.ResponseWriter, r *http.Request) {

	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	token := tokens[tokenKey]

	// 3. 使用授权后的 Request Token 换取 Access Token
	accessToken, err := f.client.AccessToken(context.Background(), oauth.RequestToken{
		Token:  token.Token,
		Secret: token.Secret,
	}, verificationCode, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%+v", err)
		return
	}

	fmt.Println("accessToken", accessToken)

	// 4.使用 Access Token 访问饭否 API
	f.client.OAuth(accessToken)

	user, err := f.client.AccountService.VerifyCredentials(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%+v", err)
		return
	}

	json.NewEncoder(w).Encode(user)
}
