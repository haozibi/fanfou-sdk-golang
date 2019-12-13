package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestUsersRecommendation(t *testing.T) {

	f := initFanfou(t)

	out, err := f.UsersService.Recommendation(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}
