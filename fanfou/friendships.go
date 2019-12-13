package fanfou

import (
	"context"
	"net/http"
)

// FriendshipsService FriendshipsService
type FriendshipsService struct {
	debug  bool
	client *http.Client
}

// NewFriendshipsService NewFriendshipsService
func NewFriendshipsService(c *http.Client, debug bool) *FriendshipsService {
	return &FriendshipsService{debug, c}
}

// Create 添加用户为好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.create
func (a *FriendshipsService) Create(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `friendships/create.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// Destroy 取消关注好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.destroy
func (a *FriendshipsService) Destroy(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `friendships/destroy.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// Requests 查询Follow请求
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.requests
func (a *FriendshipsService) Requests(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `friendships/requests.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// Deny 拒绝好友请求
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.deny
func (a *FriendshipsService) Deny(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `friendships/deny.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// Exists 查询两个用户之间是否有follow关系 如果user_a关注user_b则返回 true, 否则返回false.
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.exists
func (a *FriendshipsService) Exists(ctx context.Context, idA, idB string) (bool, error) {

	params := make(map[string]string)
	params["user_a"] = idA
	params["user_b"] = idB

	var output bool
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `friendships/exists.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// Accept 接受好友请求
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.accept
func (a *FriendshipsService) Accept(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `friendships/accept.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// FriendshipsShowOutput FriendshipsShowOutput
type FriendshipsShowOutput struct {
	Relationship *struct {
		Source *struct {
			ID                   string `json:"id"`
			ScreenName           string `json:"screen_name"`
			Following            string `json:"following"`
			FollowedBy           string `json:"followed_by"`
			NotificationsEnabled string `json:"notifications_enabled"`
			Blocking             string `json:"blocking"`
		} `json:"source"`
		Target *struct {
			ID         string `json:"id"`
			ScreenName string `json:"screen_name"`
			Following  string `json:"following"`
			FollowedBy string `json:"followed_by"`
		} `json:"target"`
	} `json:"relationship"`
}

// Show 返回两个用户之间follow关系的详细信息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/friendships.show
func (a *FriendshipsService) Show(ctx context.Context, sourceLoginName, sourceID, targetLoginName, targetID string) (*User, error) {

	params := make(map[string]string)
	if sourceLoginName != "" {
		params["source_login_name"] = sourceLoginName
	}

	if sourceID != "" {
		params["source_id"] = sourceID
	}

	if targetLoginName != "" {
		params["target_login_name"] = targetLoginName
	}

	if targetID != "" {
		params["target_id"] = targetID
	}

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `friendships/show.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}
