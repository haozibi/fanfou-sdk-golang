package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestStatusesUpdate(t *testing.T) {

	f := initFanfou(t)
	out, err := f.StatusesService.Update(context.Background(), "(╥╯^╰╥)", WithLiteMode())
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusDestroy(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.Destroy(context.Background(), "", nil)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusHomeTimeline(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.HomeTimeline(context.Background(), nil)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusReplies(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.Replies(context.Background(), nil)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusFollowers(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.Followers(context.Background(), nil)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusUserTimeline(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.UserTimeline(context.Background(), WithCount(1))
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestStatusFriends(t *testing.T) {

	f := initFanfou(t)

	out, err := f.StatusesService.Friends(context.Background(), WithCount(1))
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}
