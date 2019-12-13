package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/haozibi/fanfou-sdk-golang/fanfou"
)

var (
	consumerKey    string
	consumerSecret string
	username       string
	password       string
)

func init() {
	flag.StringVar(&consumerKey, "consumerKey", "", "consumer key from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&consumerSecret, "consumerSecret", "", "consumer secret from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&username, "username", "", "username for Fanfou.")
	flag.StringVar(&password, "password", "", "password for Fanfou.")

	flag.Parse()
}

func main() {

	if consumerKey == "" ||
		consumerSecret == "" ||
		username == "" ||
		password == "" {
		flag.PrintDefaults()
		return
	}

	f := fanfou.NewFanFou(consumerKey, consumerSecret)
	if err := f.XAuth(username, password); err != nil {
		log.Fatalln(err)
	}

	status, err := f.StatusesService.PublicTimeline(
		context.Background(),
		fanfou.WithCount(5),
	)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range status {
		fmt.Printf("user: %s \n\tsay: %s\n", v.User.Name, v.Text)
	}

}
