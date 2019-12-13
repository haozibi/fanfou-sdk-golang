package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFriendshipsExists(t *testing.T) {

	f := initFanfou(t)

	out, err := f.FriendshipsService.Exists(context.Background(), "wangxing", "wangxing")
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}
