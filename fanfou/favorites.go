package fanfou

import (
	"context"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// FavoritesService FavoritesService
type FavoritesService struct {
	debug  bool
	client *http.Client
}

// NewFavoritesService NewFavoritesService
func NewFavoritesService(c *http.Client, debug bool) *FavoritesService {
	return &FavoritesService{debug, c}
}

// Destroy 取消收藏指定消息(当前用户的收藏)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/favorites.destroy
func (f *FavoritesService) Destroy(ctx context.Context, msgID string, opts ...Option) (*Status, error) {

	if msgID == "" {
		return nil, errors.WithStack(fmt.Errorf("miss msg id"))
	}

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	path := fmt.Sprintf("favorites/destroy/%s.json", msgID)

	output := &Status{}
	req := &request{
		Debug:      f.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   path,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(f.client)
}

// Favorites 浏览指定用户收藏消息(未设置隐私用户或登录用户好友)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/favorites
func (f *FavoritesService) Favorites(ctx context.Context, opts ...Option) ([]Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      f.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `favorites/id.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(f.client)
}

// Create 收藏消息(当前用户关注者和未设置隐私用户发出的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/favorites.create
func (f *FavoritesService) Create(ctx context.Context, msgID string, opts ...Option) (*Status, error) {

	if msgID == "" {
		return nil, errors.WithStack(fmt.Errorf("miss msg id"))
	}

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	path := fmt.Sprintf("favorites/create/%s.json", msgID)

	output := &Status{}
	req := &request{
		Debug:      f.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   path,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(f.client)
}
