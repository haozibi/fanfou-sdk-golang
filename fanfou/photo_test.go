package fanfou

import (
	"context"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPhotosServiceUpload(t *testing.T) {

	f := initFanfou(t)

	file, err := os.Open("abc.jpg")
	if err != nil {
		t.Error(err)
	}

	out, err := f.PhotosService.Upload(context.Background(), file, WithStatus("Hello World"))
	if err != nil {
		t.Errorf("%+v", err)
	}

	spew.Dump(out)
}
