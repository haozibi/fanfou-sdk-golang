package fanfou

import (
	"context"
	"net/http"
)

// TrendsService TrendsService
type TrendsService struct {
	debug  bool
	client *http.Client
}

// NewTrendsService NewTrendsService
func NewTrendsService(c *http.Client, debug bool) *TrendsService { return &TrendsService{debug, c} }

// TrendsListOutput TrendsListOutput
type TrendsListOutput struct {
	AsOf   string   `json:"as_of"`
	Trends []Trends `json:"trends"`
}

// List 列出饭否热门话题
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/trends.list
func (t *TrendsService) List(ctx context.Context, opts ...Option) (*TrendsListOutput, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &TrendsListOutput{}
	req := &request{
		Debug:      t.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `trends/list.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(t.client)
}
