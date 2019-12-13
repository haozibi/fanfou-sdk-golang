package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/haozibi/fanfou-sdk-golang/fanfou"
)

var (
	consumerKey    string
	consumerSecret string
	username       string
	password       string
	status         string
	picture        string
)

func init() {
	flag.StringVar(&consumerKey, "consumerKey", "", "consumer key from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&consumerSecret, "consumerSecret", "", "consumer secret from Fanfou. https://fanfou.com/apps")
	flag.StringVar(&username, "username", "", "username for Fanfou.")
	flag.StringVar(&password, "password", "", "password for Fanfou.")

	flag.StringVar(&status, "status", "", "upload status")
	flag.StringVar(&picture, "picture", "", "upload picture")

	flag.Parse()
}

func getReader(pic string) (io.Reader, error) {

	if strings.HasPrefix(pic, "https://") ||
		strings.HasPrefix(pic, "http://") {
		resp, err := http.Get(pic)
		if err != nil {
			return nil, err
		}

		buf := &bytes.Buffer{}
		io.Copy(buf, resp.Body)
		resp.Body.Close()
		return buf, nil
	}
	body, err := ioutil.ReadFile(pic)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(body), nil
}

func main() {

	if consumerKey == "" ||
		consumerSecret == "" ||
		username == "" ||
		password == "" {
		flag.PrintDefaults()
		return
	}

	reader, err := getReader(picture)
	if err != nil {
		log.Fatalln(err)
	}

	f := fanfou.NewFanFou(consumerKey, consumerSecret)
	if err := f.XAuth(username, password); err != nil {
		log.Fatalln(err)
	}

	if status == "" {
		status = "Hello World"
	}
	status += `By "github.com/haozibi/fanfou-sdk-golang"`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	s, err := f.PhotosService.Upload(
		ctx,
		reader,
		fanfou.WithStatus(status),
		fanfou.WithLocation("Earth"),
	)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("update status success, id: %s\n", s.ID)
}
