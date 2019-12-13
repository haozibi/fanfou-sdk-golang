package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSavedSearchesCreate(t *testing.T) {

	f := initFanfou(t)

	out, err := f.SavedSearchesService.Create(context.Background(), []string{"美", "abc"})
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestSavedSearchesDestroy(t *testing.T) {

	f := initFanfou(t)

	out, err := f.SavedSearchesService.Destroy(context.Background(), 95406)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestSavedSearchesShow(t *testing.T) {

	f := initFanfou(t)

	out, err := f.SavedSearchesService.Show(context.Background(), 95407)
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}

func TestSavedSearchesList(t *testing.T) {

	f := initFanfou(t)

	out, err := f.SavedSearchesService.List(context.Background())
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}
