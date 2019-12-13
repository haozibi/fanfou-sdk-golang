package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/haozibi/fanfou-sdk-golang/fanfou"
	"github.com/haozibi/fanfou-sdk-golang/oauth"
)

var (
	consumerKey    string
	consumerSecret string
)

func init() {
	flag.StringVar(&consumerKey, "consumerKey", "", "consumer key from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&consumerSecret, "consumerSecret", "", "consumer secret from Fanfou. https://fanfou.com/apps")
	flag.Parse()
}

func main() {
	if consumerKey == "" ||
		consumerSecret == "" {
		flag.PrintDefaults()
		return
	}

	f := fanfou.NewFanFou(consumerKey, consumerSecret)

	token, err := f.RequestToken(context.Background(), "oob")
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
	uri, err := f.AuthorizationURL(token)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	fmt.Println("(1) Go to: " + uri)
	fmt.Println("(2) Grant access, you should get back a verification code.")
	fmt.Println("(3) Enter that verification code here: ")

	verificationCode := ""
	_, err = fmt.Scanln(&verificationCode)
	if err != nil {
		log.Fatalln(err)
	}

	accessToken, err := f.AccessToken(context.Background(), oauth.RequestToken{
		Token:  token.Token,
		Secret: token.Secret,
	}, verificationCode, nil)

	f.OAuth(accessToken)

	fmt.Println("=>", accessToken)

	status, err := f.StatusesService.Update(context.Background(), "abc")
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(os.Stdout).Encode(status)
}
