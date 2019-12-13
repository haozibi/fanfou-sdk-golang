package fanfou

import (
	"context"
	"net/http"
)

// BlocksService BlocksService
type BlocksService struct {
	debug  bool
	client *http.Client
}

// NewBlocksService NewBlocksService
func NewBlocksService(c *http.Client, debug bool) *BlocksService {
	return &BlocksService{debug, c}
}

// Ids 获取用户黑名单id列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/blocks.ids
func (b *BlocksService) Ids(ctx context.Context) ([]string, error) {

	output := make([]string, 0)
	req := &request{
		Debug:      b.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `blocks/ids.json`,
		Input:      nil,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(b.client)
}

// Blocking 获取黑名单上用户资料
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/blocks.blocking
func (b *BlocksService) Blocking(ctx context.Context, opts ...Option) ([]User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]User, 0)
	req := &request{
		Debug:      b.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `blocks/blocking.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(b.client)
}

// Create 把指定 id 用户加入黑名单，指定目标用户的 user_id 或者 loginname
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/blocks.create
func (b *BlocksService) Create(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      b.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `blocks/create.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(b.client)
}

// Exists 检查用户是否被加入了黑名单，指定目标用户的 user_id 或者 loginname
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/blocks.exists
func (b *BlocksService) Exists(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      b.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `blocks/exists.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(b.client)
}

// Destroy 将指定id用户解除黑名单
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/blocks.destroy
func (b *BlocksService) Destroy(ctx context.Context, id string, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &User{}
	req := &request{
		Debug:      b.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `blocks/destroy.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(b.client)
}
