package fanfou

import (
	"context"
	"net/http"
)

// UsersService Users
type UsersService struct {
	debug  bool
	client *http.Client
}

// NewUsersService NewUsersService
func NewUsersService(c *http.Client, debug bool) *UsersService {
	return &UsersService{debug, c}
}

// Tagged 返回标记为指定标签的用户列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.tagged
func (u *UsersService) Tagged(ctx context.Context, tag string, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "tag")
	params["tag"] = tag

	output := make([]User, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `users/tagged.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// Show 返回好友或未设置隐私用户的信息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.show
func (u *UsersService) Show(ctx context.Context, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &User{}
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `users/show.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// TagList 获取用户标签列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.tag-list
func (u *UsersService) TagList(ctx context.Context, opts ...Option) ([]string, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]string, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `users/tag_list.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// Followers 返回用户的最近登录的关注者
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.followers
func (u *UsersService) Followers(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `users/followers.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// Recommendation 返回系统推荐的好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.recommendation
func (u *UsersService) Recommendation(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `2/users/recommendation.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// CancelRecommendation 忽略系统推荐的好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.cancel-recommendation
func (u *UsersService) CancelRecommendation(ctx context.Context, id string, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := make([]User, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `/2/users/cancel_recommendation.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}

// Friends 返回最近登录的用户好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/users.friends
func (u *UsersService) Friends(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      u.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `users/friends.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(u.client)
}
