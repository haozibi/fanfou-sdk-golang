package fanfou

import (
	"context"
	"net/http"
)

// StatusesService Statuses
type StatusesService struct {
	debug  bool
	client *http.Client
}

// NewStatusesService NewStatusesService
func NewStatusesService(c *http.Client, debug bool) *StatusesService {
	return &StatusesService{debug, c}
}

// Destroy 删除指定的消息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.destroy
func (s *StatusesService) Destroy(ctx context.Context, msgID string, opts ...Option) (*Status, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = msgID

	output := &Status{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `statuses/destroy.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// HomeTimeline 显示指定用户及其好友的消息(未设置隐私用户和登录用户好友的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.home-timeline
func (s *StatusesService) HomeTimeline(ctx context.Context, opts ...Option) ([]Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/home_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// PublicTimeline 显示20条随便看看的消息(未设置隐私用户的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.public-timeline
func (s *StatusesService) PublicTimeline(ctx context.Context, opts ...Option) ([]Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/public_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Replies 显示回复当前用户的20条消息(未设置隐私用户和登录用户好友的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.replies
func (s *StatusesService) Replies(ctx context.Context, opts ...Option) ([]Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/replies.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Followers 返回用户的前100个关注者
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.followers
func (s *StatusesService) Followers(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/followers.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Update 发送消息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.update
func (s *StatusesService) Update(ctx context.Context, status string, opts ...Option) (*Status, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "status")
	params["status"] = status

	output := &Status{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `statuses/update.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// UserTimeline 浏览指定用户已发送消息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.user-timeline
func (s *StatusesService) UserTimeline(ctx context.Context, opts ...Option) ([]Status, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/user_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Friends 返回用户好友
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.friends
func (s *StatusesService) Friends(ctx context.Context, opts ...Option) ([]User, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/friends.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// ContextTimeline 按照时间先后顺序显示消息上下文(好友和未设置隐私用户的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.context-timeline
func (s *StatusesService) ContextTimeline(ctx context.Context, opts ...Option) ([]Status, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/context_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Mentions 显示回复/提到当前用户的20条消息(未设置隐私用户和登录用户好友的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.mentions
func (s *StatusesService) Mentions(ctx context.Context, opts ...Option) ([]Status, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/mentions.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Show 返回好友或未设置隐私用户的某条消息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/statuses.show
func (s *StatusesService) Show(ctx context.Context, msgID string, opts ...Option) ([]Status, error) {
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = msgID

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `statuses/show.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}
