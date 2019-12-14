package fanfou

import (
	"context"
	"net/http"
)

// DirectMessagesService DirectMessagesService
type DirectMessagesService struct {
	debug  bool
	client *http.Client
}

// NewDirectMessagesService NewDirectMessagesService
func NewDirectMessagesService(c *http.Client, debug bool) *DirectMessagesService {
	return &DirectMessagesService{debug, c}
}

// Destroy 删除某条私信
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.destroy
func (d *DirectMessagesService) Destroy(ctx context.Context, id string, opts ...Option) (*DirectMessages, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := &DirectMessages{}
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `direct_messages/destroy.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}

// Conversation 以对话的形式返回当前用户与某用户的私信
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.conversation
func (d *DirectMessagesService) Conversation(ctx context.Context, id string, opts ...Option) ([]DirectMessages, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	params["id"] = id

	output := make([]DirectMessages, 0)
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `direct_messages/conversation.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}

// New 发送私信
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.new
func (d *DirectMessagesService) New(ctx context.Context, id, text string, opts ...Option) (*DirectMessages, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}
	delete(params, "id")
	delete(params, "text")
	params["id"] = id
	params["text"] = text

	output := &DirectMessages{}
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `direct_messages/new.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}

// DirectMessagesConversationListOutput DirectMessagesConversationListOutput
type DirectMessagesConversationListOutput struct {
	OtherID string          `json:"otherid"`
	MsgNum  int             `json:"msg_num"`
	NewConv bool            `json:"new_conv"`
	DM      *DirectMessages `json:"dm"`
}

// ConversationList 以对话的形式返回当前用户的私信列表
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.conversation-list
func (d *DirectMessagesService) ConversationList(ctx context.Context, opts ...Option) (*DirectMessagesConversationListOutput, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &DirectMessagesConversationListOutput{}
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `direct_messages/conversation_list.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}

// Inbox 显示20条收件箱中的私信
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.inbox
func (d *DirectMessagesService) Inbox(ctx context.Context, opts ...Option) ([]DirectMessages, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]DirectMessages, 0)
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `direct_messages/inbox.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}

// Sent 显示发件箱中的私信
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/direct-messages.sent
func (d *DirectMessagesService) Sent(ctx context.Context, opts ...Option) ([]DirectMessages, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := make([]DirectMessages, 0)
	req := &request{
		Debug:      d.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `direct_messages/sent.json`,
		Input:      params,
		Output:     &output,
		Context:    ctx,
	}

	return output, req.send(d.client)
}
