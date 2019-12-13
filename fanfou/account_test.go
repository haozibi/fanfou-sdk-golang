package fanfou

import (
	"context"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestAccountUpdateProfileImage(t *testing.T) {

	f := initFanfou(t)

	file, err := os.Open("123.jpg")
	if err != nil {
		t.Error(err)
	}

	out, err := f.AccountService.UpdateProfileImage(context.Background(), file)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}

func TestAccountUpdateNotifyNum(t *testing.T) {

	f := initFanfou(t)

	out, err := f.AccountService.UpdateNotifyNum(context.Background(), 9)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}

func TestAccountNotifyNumOutput(t *testing.T) {

	f := initFanfou(t)

	out, err := f.AccountService.NotifyNum(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}

func TestAccountVerifyCredentials(t *testing.T) {

	f := initFanfou(t)

	out, err := f.AccountService.VerifyCredentials(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}
