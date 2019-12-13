package fanfou

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func initFanfou(t *testing.T) *Fanfou {

	body, err := ioutil.ReadFile("../example/1.json")
	if err != nil {
		t.Error(err)
	}

	var m map[string]string
	json.Unmarshal(body, &m)

	var (
		consumerKey    = m["consumer_key"]
		consumerSecret = m["consumer_secret"]
		username       = m["username"]
		password       = m["password"]
	)

	f := NewFanFou(consumerKey, consumerSecret)

	if err := f.XAuth(username, password); err != nil {
		t.Errorf("%+v\n", err)
	}
	return f
}
