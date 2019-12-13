package fanfou

import (
	"context"
	"net/http"
)

// FollowersService FollowersService
type FollowersService struct {
	debug  bool
	client *http.Client
}

// NewFollowersService NewFollowersService
func NewFollowersService(c *http.Client, debug bool) *FollowersService {
	return &FollowersService{debug, c}
}

// Ids 返回用户关注者的id列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/followers.ids
func (f *FollowersService) Ids(ctx context.Context, opts ...Option) ([]string, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]string, 0)
	req := &request{
		Debug:      f.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `followers/ids.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(f.client)
}
