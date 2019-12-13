package fanfou

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestFavoritesDestroy(t *testing.T) {

	f := initFanfou(t)

	out, err := f.FavoritesService.Destroy(context.Background(), "")
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}

func TestFavoritesFavorites(t *testing.T) {

	f := initFanfou(t)

	out, err := f.FavoritesService.Favorites(context.Background())
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}

func TestFavoritesCreate(t *testing.T) {

	f := initFanfou(t)

	out, err := f.FavoritesService.Create(context.Background(), "")
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	spew.Dump(out)
}
