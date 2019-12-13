package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFriendsIds(t *testing.T) {

	f := initFanfou(t)

	out, err := f.FriendsService.Ids(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}
