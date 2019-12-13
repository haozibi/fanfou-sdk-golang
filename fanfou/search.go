package fanfou

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// SearchService SearchService
type SearchService struct {
	debug  bool
	client *http.Client
}

// NewSearchService new SearchService
func NewSearchService(c *http.Client, debug bool) *SearchService {
	return &SearchService{debug, c}
}

// PublicTimeline 搜索全站消息(未设置隐私用户的消息)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/search.public-timeline
func (s *SearchService) PublicTimeline(ctx context.Context, keywords []string, opts ...Option) ([]Status, error) {

	if len(keywords) == 0 {
		return nil, errors.WithStack(fmt.Errorf("miss keywords"))
	}
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "q")
	params["q"] = strings.Join(keywords, "|")

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `search/public_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// SearchUsersOutput SearchUsersOutput
type SearchUsersOutput struct {
	TotalNumber int    `json:"total_number"`
	Users       []User `json:"users"`
}

// Users 【maybe 失效】搜索全站用户(只返回其中未被ban掉的用户)
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/search.users
func (s *SearchService) Users(ctx context.Context, keywords []string, opts ...Option) (*SearchUsersOutput, error) {
	if len(keywords) == 0 {
		return nil, errors.WithStack(fmt.Errorf("miss keywords"))
	}

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "q")
	params["q"] = strings.Join(keywords, "|")

	output := &SearchUsersOutput{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `search/users.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// UserTimeline 搜索指定用户消息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/search.user-timeline
func (s *SearchService) UserTimeline(ctx context.Context, keywords []string, opts ...Option) ([]Status, error) {

	if len(keywords) == 0 {
		return nil, errors.WithStack(fmt.Errorf("miss keywords"))
	}
	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "q")
	params["q"] = strings.Join(keywords, "|")

	output := make([]Status, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `search/user_timeline.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}
