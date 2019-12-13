package fanfou

import (
	"context"
	"net/http"
)

// FriendsService Friends
type FriendsService struct {
	debug  bool
	client *http.Client
}

// NewFriendsService NewFriendsService
func NewFriendsService(c *http.Client, debug bool) *FriendsService { return &FriendsService{debug, c} }

// Ids 返回用户好友的id列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friends.ids
func (f *FriendsService) Ids(ctx context.Context, opts ...Option) ([]string, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]string, 0)
	req := &request{
		Debug:      f.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `friends/ids.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(f.client)
}
