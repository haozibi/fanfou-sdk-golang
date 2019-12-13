package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestTrendsList(t *testing.T) {

	f := initFanfou(t)
	output, err := f.TrendsService.List(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(output)
}
