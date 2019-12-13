package fanfou

import (
	"context"
	"io"
	"net/http"
	"strconv"
)

// AccountService AccountService
type AccountService struct {
	debug  bool
	client *http.Client
}

// NewAccountService NewAccountService
func NewAccountService(c *http.Client, debug bool) *AccountService { return &AccountService{debug, c} }

// VerifyCredentials 检查用户名密码是否正确
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.verify-credentials
func (a *AccountService) VerifyCredentials(ctx context.Context, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `account/verify_credentials.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// UpdateProfileImage 通过 API 更新用户头像
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.update-profile-image
func (a *AccountService) UpdateProfileImage(ctx context.Context, reader io.Reader, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	ctx = context.WithValue(ctx, fileRequestKey, "image")

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `account/update_profile_image.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
		Reader:     reader,
	}

	return output, req.send(a.client)
}

// RateLimitStatusOutput RateLimitStatusResp
type RateLimitStatusOutput struct {
	ResetTime          string `json:"reset_time"`
	RemainingHits      int    `json:"remaining_hits"`
	HourlyLimit        int    `json:"hourly_limit"`
	ResetTimeInSeconds int64  `json:"reset_time_in_seconds"`
}

// RateLimitStatus 获取 API 限制
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.rate-limit-status
func (a *AccountService) RateLimitStatus(ctx context.Context, opts ...Option) (*RateLimitStatusOutput, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &RateLimitStatusOutput{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `account/rate_limit_status.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// UpdateProfile 通过 API 更新用户资料
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.update-profile
func (a *AccountService) UpdateProfile(ctx context.Context, opts ...Option) (*User, error) {

	params := make(map[string]string)
	for _, o := range opts {
		o.apply(params)
	}

	output := &User{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `account/update_profile.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// NotificationOutput NotificationOutput
type NotificationOutput struct {
	Mentions       int `json:"mentions"`
	DirectMessages int `json:"direct_messages"`
	FriendRequests int `json:"friend_requests"`
}

// Notification 返回未读的mentions, direct message 以及关注请求数量
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.notification
func (a *AccountService) Notification(ctx context.Context) (*NotificationOutput, error) {

	output := &NotificationOutput{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `account/notification.json`,
		Input:      nil,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// UpdateNotifyNumOutput NotifyNumOutput
type UpdateNotifyNumOutput struct {
	Result    string `json:"result"`
	NotifyNum string `json:"notify_num"`
}

// UpdateNotifyNum 向饭否更新当前app上的新提醒数量
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.update-notify-num
func (a *AccountService) UpdateNotifyNum(ctx context.Context, notifyNum int) (*UpdateNotifyNumOutput, error) {

	params := make(map[string]string)
	params["notify_num"] = strconv.Itoa(notifyNum)

	output := &UpdateNotifyNumOutput{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodPost,
		HTTPPath:   `account/update_notify_num.json`,
		Input:      params,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}

// NotifyNumOutput NotifyNumOutput
type NotifyNumOutput struct {
	Result    string `json:"result"`
	NotifyNum int    `json:"notify_num"`
}

// NotifyNum 获取当前app上的新提醒数量
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/account.notify-num
func (a *AccountService) NotifyNum(ctx context.Context) (*NotifyNumOutput, error) {

	output := &NotifyNumOutput{}
	req := &request{
		Debug:      a.debug,
		HTTPMethod: http.MethodGet,
		HTTPPath:   `account/notify_num.json`,
		Input:      nil,
		Output:     output,
		Context:    ctx,
	}

	return output, req.send(a.client)
}
