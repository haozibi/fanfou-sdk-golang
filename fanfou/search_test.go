package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSearchPublicTimeline(t *testing.T) {

	f := initFanfou(t)

	output, err := f.SearchService.PublicTimeline(
		context.Background(),
		[]string{"美"},
	)
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(output)
}
