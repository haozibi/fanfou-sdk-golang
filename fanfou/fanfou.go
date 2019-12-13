package fanfou

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/haozibi/fanfou-sdk-golang/oauth"
	"github.com/haozibi/fanfou-sdk-golang/utils"

	"github.com/pkg/errors"
)

var (
	userAgent = "github.com/haozibi/fanfou-sdk-golang"

	// BaseURL represents Fanfou API base URL
	BaseURL = "https://api.fanfou.com/"

	// AuthBaseURL represents Fanfou API authorization base URL
	AuthBaseURL = "https://fanfou.com/"

	requestTokenURI   = "oauth/request_token"
	authorizeTokenURI = "oauth/authorize"
	accessTokenURI    = "oauth/access_token"
)

const (
	fileRequestKey = "sdk_file_request_key"
)

// Fanfou fanfou
type Fanfou struct {
	debug          bool
	consumerkey    string
	consumerSecret string

	SearchService         *SearchService
	BlocksService         *BlocksService
	UsersService          *UsersService
	AccountService        *AccountService
	SavedSearchesService  *SavedSearchesService
	PhotosService         *PhotosService
	TrendsService         *TrendsService
	FollowersService      *FollowersService
	FavoritesService      *FavoritesService
	FriendshipsService    *FriendshipsService
	FriendsService        *FriendsService
	StatusesService       *StatusesService
	DirectMessagesService *DirectMessagesService

	client *http.Client
	oauth  *oauth.OAuth
}

// NewFanFou new fanfou
func NewFanFou(consumerkey, consumerSecret string) *Fanfou {
	return NewFanFouWithClient(consumerkey, consumerSecret, &http.Client{})
}

// NewFanFouWithClient new fanfou with client
func NewFanFouWithClient(consumerkey, consumerSecret string, client *http.Client) *Fanfou {
	f := &Fanfou{
		debug:          utils.IsDebug(),
		consumerkey:    consumerkey,
		consumerSecret: consumerSecret,
		client:         client,
		oauth: oauth.NewOAuthWithClient(
			consumerkey,
			consumerSecret,
			oauth.ServiceProvider{
				RequestTokenURL:   AuthBaseURL + requestTokenURI,
				AuthorizeTokenURL: AuthBaseURL + authorizeTokenURI,
				AccessTokenURL:    AuthBaseURL + accessTokenURI,
			},
			client,
		),
	}

	return f
}

func (f *Fanfou) initService() {
	f.SearchService = NewSearchService(f.client, f.debug)
	f.BlocksService = NewBlocksService(f.client, f.debug)
	f.UsersService = NewUsersService(f.client, f.debug)
	f.AccountService = NewAccountService(f.client, f.debug)
	f.SavedSearchesService = NewSavedSearchesService(f.client, f.debug)
	f.PhotosService = NewPhotosService(f.client, f.debug)
	f.TrendsService = NewTrendsService(f.client, f.debug)
	f.FollowersService = NewFollowersService(f.client, f.debug)
	f.FavoritesService = NewFavoritesService(f.client, f.debug)
	f.FriendshipsService = NewFriendshipsService(f.client, f.debug)
	f.FriendsService = NewFriendsService(f.client, f.debug)
	f.StatusesService = NewStatusesService(f.client, f.debug)
	f.DirectMessagesService = NewDirectMessagesService(f.client, f.debug)
}

func (f *Fanfou) initOAuthConsumer() {

}

// RequestToken 获取未授权的Request Token
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Oauth#%E8%8E%B7%E5%8F%96%E6%9C%AA%E6%8E%88%E6%9D%83%E7%9A%84request-token
func (f *Fanfou) RequestToken(ctx context.Context, callback string) (*oauth.RequestToken, error) {
	f.oauth.ServiceProvider.CallBackURL = callback
	return f.oauth.GetRequestToken(ctx)
}

// AuthorizationURL 请求用户授权Request Token
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Oauth#%E8%AF%B7%E6%B1%82%E7%94%A8%E6%88%B7%E6%8E%88%E6%9D%83request-token
func (f *Fanfou) AuthorizationURL(token *oauth.RequestToken) (string, error) {
	return f.oauth.GetAuthorizationURL(token)
}

// AccessToken 使用授权后的 Request Token 换取 Access Token
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Oauth#%E4%BD%BF%E7%94%A8%E6%8E%88%E6%9D%83%E5%90%8E%E7%9A%84request-token%E6%8D%A2%E5%8F%96access-token
func (f *Fanfou) AccessToken(ctx context.Context, token oauth.RequestToken, verificationCode string, additionalParams map[string]string) (*oauth.AccessToken, error) {
	return f.oauth.GetAccessToken(ctx, token, verificationCode, additionalParams)
}

// OAuth 使用已经生成的 AccessToken 认证
func (f *Fanfou) OAuth(accessToken *oauth.AccessToken) error {
	if accessToken == nil ||
		accessToken.Token == "" ||
		accessToken.Secret == "" {
		return errors.WithStack(fmt.Errorf("accessToken error"))
	}
	f.client = f.oauth.BuildClient(accessToken)
	f.initService()
	return nil
}

// XAuth 为移动设备简化的 OAuth 认证
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Xauth
func (f *Fanfou) XAuth(username, password string) error {

	accessToken, err := f.oauth.GetAccessToken(
		context.Background(),
		oauth.RequestToken{},
		"",
		map[string]string{
			"x_auth_username": username,
			"x_auth_password": password,
			"x_auth_mode":     "client_auth",
		},
	)
	if err != nil {
		return err
	}

	if f.debug {
		utils.ShowInfomation("Fanfou AccessToken", accessToken)
	}

	return f.OAuth(accessToken)
}

type request struct {
	Debug      bool
	HTTPMethod string
	HTTPPath   string
	Input      map[string]string
	Output     interface{} // must be ptr
	Context    context.Context
	Reader     io.Reader
	req        *http.Request
}

func (r *request) buildRequest() error {
	if r.HTTPMethod == "" {
		r.HTTPMethod = http.MethodGet
	}

	var (
		req    *http.Request
		err    error
		uri    = BaseURL + r.HTTPPath
		params = buildQuery(r.Input)
	)

	if r.HTTPMethod == http.MethodPost {
		req, err = http.NewRequestWithContext(r.Context, r.HTTPMethod, uri, strings.NewReader(params.Encode()))
		if err != nil {
			return errors.WithStack(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequestWithContext(r.Context, r.HTTPMethod, uri, nil)
		if err != nil {
			return errors.WithStack(err)
		}
		req.URL.RawQuery = params.Encode()
	}

	req.Header.Add("User-Agent", userAgent)
	r.req = req
	return nil
}

func (r *request) buildFileRequest(reader io.Reader) error {
	if r.HTTPMethod == "" {
		r.HTTPMethod = http.MethodPost
	}

	if reader == nil {
		return errors.WithStack(fmt.Errorf("miss file content"))
	}

	var (
		uri    = BaseURL + r.HTTPPath
		buf    = new(bytes.Buffer)
		writer = multipart.NewWriter(buf)
		params = buildQuery(r.Input)
	)

	fileName := "photo"
	if r.Context.Value(fileRequestKey) != nil {
		if t, ok := r.Context.Value(fileRequestKey).(string); ok {
			fileName = t
		}
	}

	for k := range params {
		writer.WriteField(k, params.Get(k))
	}
	formFile, _ := writer.CreateFormFile(fileName, "photo.jpg")
	io.Copy(formFile, reader)
	contentType := writer.FormDataContentType()
	writer.Close()

	req, err := http.NewRequestWithContext(r.Context, r.HTTPMethod, uri, buf)
	if err != nil {
		return errors.WithStack(err)
	}

	req.Header.Set("Content-Type", contentType)
	req.Header.Add("User-Agent", userAgent)
	r.req = req
	return nil
}

func (r *request) send(c *http.Client) error {

	if c == nil {
		return errors.WithStack(fmt.Errorf("miss http client"))
	}

	cleanList := cleanOption(r.Input, r.HTTPPath)
	if len(cleanList) != 0 {
		log.Printf("[fanfou-sdk] warning, request %s not support parmas %v, and delete them\n", r.HTTPPath, cleanList)
	}

	var err error
	if r.HTTPMethod == http.MethodPost &&
		r.Reader != nil {
		err = r.buildFileRequest(r.Reader)
	} else {
		err = r.buildRequest()
	}

	if err != nil {
		return err
	}

	if r.Debug {
		utils.ShowInfomation("Fanfou Request", r.req)
	}

	resp, err := c.Do(r.req)
	if err != nil {
		return errors.WithStack(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	if r.Debug {
		utils.ShowInfomation("Fanfou Response", string(body))
	}

	if err := r.handleError(resp, body); err != nil {
		return err
	}

	err = json.Unmarshal(body, r.Output)
	return errors.WithStack(err)
}

func (r *request) handleError(resp *http.Response, body []byte) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	e := &utils.ErrorResponse{}
	err := json.Unmarshal(body, e)
	if err != nil {
		return errors.WithStack(err)
	}
	e.StatusCode = resp.StatusCode
	return errors.WithStack(e)
}

func buildQuery(input map[string]string) url.Values {
	q := url.Values{}
	if len(input) == 0 {
		return q
	}

	for k, v := range input {
		if v != "" {
			q.Add(k, v)
		}
	}

	return q
}

func buildQuery2(input interface{}) url.Values {

	q := url.Values{}

	if input == nil {
		return q
	}

	t := reflect.TypeOf(input).Elem()
	v := reflect.ValueOf(input).Elem()
	if !v.IsValid() {
		return q
	}

	for i := 0; i < v.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		if v.Field(i).IsValid() &&
			!v.Field(i).IsZero() {
			q.Add(key, fmt.Sprintf("%v", v.Field(i)))
		}
	}
	return q
}

// cleanOption 保证 request 不包含多余参数
func cleanOption(params map[string]string, path string) []string {
	if len(params) == 0 {
		return nil
	}

	var f = func(path string) string {
		if strings.Index(path, "favorites/") == -1 {
			return path
		}

		paths := strings.Split(path, "/")
		if len(paths) == 0 {
			return path
		}
		path = strings.TrimSuffix(path, paths[len(paths)-1])
		return path
	}

	var (
		allow = allowParams[f(path)]
		tmp   = make(map[string]bool)
		clean = make([]string, 0, len(params))
	)

	for _, v := range allow {
		tmp[v] = true
	}

	for k := range params {
		if !tmp[k] {
			delete(params, k)
			clean = append(clean, k)
		}
	}
	return clean
}

// Option Option
type Option interface {
	apply(map[string]string)
}

type optionFunc func(map[string]string)

func (f optionFunc) apply(o map[string]string) {
	f(o)
}

// WithID 目标 id
func WithID(id string) Option {
	return optionFunc(func(m map[string]string) {
		m["id"] = id
	})
}

// WithSource 消息来源
func WithSource(source string) Option {
	return optionFunc(func(m map[string]string) {
		m["source"] = source
	})
}

// WithLocation 发布消息的地点
//
// 最多30个字符, 使用'地点名称' 或 '一个半角逗号分隔的经纬度坐标'。如：北京市海淀区 或者 39.9594049,116.298419
func WithLocation(localtion string) Option {
	return optionFunc(func(m map[string]string) {
		m["location"] = localtion
	})
}

// WithLiteMode 用户信息不包含用户自定义 profile
func WithLiteMode() Option {
	return optionFunc(func(m map[string]string) {
		m["mode"] = "lite"
	})
}

// WithHTMLFormat 返回消息中 @ 提到的用户,网址等输出 html 链接
func WithHTMLFormat() Option {
	return optionFunc(func(m map[string]string) {
		m["format"] = "html"
	})
}

// WithCount 指定一次请求返回的数量
func WithCount(count int) Option {
	return optionFunc(func(m map[string]string) {
		m["count"] = strconv.Itoa(count)
	})
}

// WithSinceID 只返回消息 id 大于 sinceID 的消息
func WithSinceID(sinceID int64) Option {
	return optionFunc(func(m map[string]string) {
		m["since_id"] = strconv.FormatInt(sinceID, 10)
	})
}

// WithMaxID 只返回消息 id 小于等于 max_id 的消息
func WithMaxID(maxID int64) Option {
	return optionFunc(func(m map[string]string) {
		m["max_id"] = strconv.FormatInt(maxID, 10)
	})
}

// WithPage 指定返回结果的页码
func WithPage(pageID int) Option {
	return optionFunc(func(m map[string]string) {
		m["page"] = strconv.Itoa(pageID)
	})
}

// WithStatus 描述
func WithStatus(status string) Option {
	return optionFunc(func(m map[string]string) {
		m["status"] = status
	})
}

// WithURL 指定自定义网址
func WithURL(url string) Option {
	return optionFunc(func(m map[string]string) {
		m["url"] = url
	})
}

// WithDescription 指定自述
func WithDescription(description string) Option {
	return optionFunc(func(m map[string]string) {
		m["description"] = description
	})
}

// WithName 指定姓名
func WithName(name string) Option {
	return optionFunc(func(m map[string]string) {
		m["name"] = name
	})
}

// WithEmail 指定电子邮件
func WithEmail(email string) Option {
	return optionFunc(func(m map[string]string) {
		m["email"] = email
	})
}

// WithInReplyToStatusID 回复的消息 id
func WithInReplyToStatusID(msgID string) Option {
	return optionFunc(func(m map[string]string) {
		m["in_reply_to_status_id"] = msgID
	})
}

// WithInReplyToUserID 回复的用户 id
func WithInReplyToUserID(msgID string) Option {
	return optionFunc(func(m map[string]string) {
		m["in_reply_to_user_id"] = msgID
	})
}

// WithRepostStatusID 转发的消息 id
func WithRepostStatusID(msgID string) Option {
	return optionFunc(func(m map[string]string) {
		m["repost_status_id"] = msgID
	})
}

// WithInReplyToID 回复的私信 id
func WithInReplyToID(msgID string) Option {
	return optionFunc(func(m map[string]string) {
		m["in_reply_to_id"] = msgID
	})
}

// WithCallback 【舍弃】当使用 json 格式时,生成的 json 对象将作为参数传给指定的 javascript 函数
func WithCallback(name string) Option {
	return optionFunc(func(m map[string]string) {
		// m["callback"] = name
	})
}
