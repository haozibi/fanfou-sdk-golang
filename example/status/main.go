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
	status         string
)

func init() {
	flag.StringVar(&consumerKey, "consumerKey", "", "consumer key from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&consumerSecret, "consumerSecret", "", "consumer secret from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&username, "username", "", "username for Fanfou.")
	flag.StringVar(&password, "password", "", "password for Fanfou.")

	flag.StringVar(&status, "status", "", "upload status")

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

	if status == "" {
		status = "Hello World"
	}
	status += `By "github.com/haozibi/fanfou-sdk-golang"`

	s, err := f.StatusesService.Update(
		context.Background(),
		status,
		fanfou.WithLocation("Earth"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("update status success, id: %s\n", s.ID)
}
