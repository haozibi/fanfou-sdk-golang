package fanfou

import (
	"context"
	"io"
	"net/http"
)

// PhotosService Photos
type PhotosService struct {
	debug  bool
	client *http.Client
}

// NewPhotosService NewPhotosService
func NewPhotosService(c *http.Client, debug bool) *PhotosService { return &PhotosService{debug, c} }

// UserTimeline 浏览指定用户的图片
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/photos.user-timeline
func (p *PhotosService) UserTimeline(ctx context.Context, opts ...Option) ([]Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      p.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `photos/user_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(p.client)
}

// Upload 上传图片
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/photos.upload
func (p *PhotosService) Upload(ctx context.Context, reader io.Reader, opts ...Option) (*Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	ctx = context.WithValue(ctx, fileRequestKey, "photo")

	output := &Status{}
	req := &request{
		Debug:      p.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `photos/upload.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
		Reader:     reader,
	}

	return output, req.send(p.client)
}
