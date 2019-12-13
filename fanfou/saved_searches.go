package fanfou

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

// SavedSearchesService SavedSearchesService
type SavedSearchesService struct {
	debug  bool
	client *http.Client
}

// NewSavedSearchesService NewSavedSearchesService
func NewSavedSearchesService(c *http.Client, debug bool) *SavedSearchesService {
	return &SavedSearchesService{debug, c}
}

// Create 收藏搜索关键字
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/saved-searches.create
func (s *SavedSearchesService) Create(ctx context.Context, query []string, opts ...Option) (*SavedSearches, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "query")
	params["query"] = strings.Join(query, "|")

	output := &SavedSearches{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `saved_searches/create.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Destroy 删除收藏的搜索关键字
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/saved-searches.destroy
func (s *SavedSearchesService) Destroy(ctx context.Context, id int, opts ...Option) (*SavedSearches, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = strconv.Itoa(id)

	output := &SavedSearches{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `saved_searches/destroy.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// Show 返回搜索关键字的详细信息
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/saved-searches.show
func (s *SavedSearchesService) Show(ctx context.Context, id int, opts ...Option) (*SavedSearches, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = strconv.Itoa(id)

	output := &SavedSearches{}
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `saved_searches/show.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}

// List 列出登录用户保存的搜索关键字
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/saved-searches.list
func (s *SavedSearchesService) List(ctx context.Context, opts ...Option) ([]SavedSearches, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]SavedSearches, 0)
	req := &request{
		Debug:      s.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `saved_searches/list.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(s.client)
}
