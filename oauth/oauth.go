// Package oauth only oAuth 1.0
package oauth

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/haozibi/fanfou-sdk-golang/utils"

	"github.com/pkg/errors"
)

const (
	OAUTH_VERSION          = "1.0"
	HTTP_AUTH_HEADER       = "Authorization"
	OAUTH_HEADER           = "OAuth "
	BODY_HASH_PARAM        = "oauth_body_hash"
	CALLBACK_PARAM         = "oauth_callback"
	CONSUMER_KEY_PARAM     = "oauth_consumer_key"
	NONCE_PARAM            = "oauth_nonce"
	SESSION_HANDLE_PARAM   = "oauth_session_handle"
	SIGNATURE_METHOD_PARAM = "oauth_signature_method"
	SIGNATURE_PARAM        = "oauth_signature"
	TIMESTAMP_PARAM        = "oauth_timestamp"
	TOKEN_PARAM            = "oauth_token"
	TOKEN_SECRET_PARAM     = "oauth_token_secret"
	VERIFIER_PARAM         = "oauth_verifier"
	VERSION_PARAM          = "oauth_version"
)

// OAuth OAuth
type OAuth struct {
	debug          bool
	consumerKey    string
	consumerSecret string
	realm          string

	ServiceProvider *ServiceProvider
	client          *http.Client
}

// ServiceProvider ServiceProvider
type ServiceProvider struct {
	RequestTokenURL   string
	AuthorizeTokenURL string
	AccessTokenURL    string
	CallBackURL       string
}

// NewOAuth NewOAuth
func NewOAuth(consumerKey, consumerSecret string, serviceProvider ServiceProvider) *OAuth {
	return NewOAuthWithClient(consumerKey, consumerSecret, serviceProvider, &http.Client{})
}

// NewOAuthWithClient NewOAuthWithClient
func NewOAuthWithClient(consumerKey, consumerSecret string, serviceProvider ServiceProvider, client *http.Client) *OAuth {
	return &OAuth{
		debug:           utils.IsDebug(),
		consumerKey:     consumerKey,
		consumerSecret:  consumerSecret,
		ServiceProvider: &serviceProvider,
		client:          &http.Client{},
	}
}

// GetRequestToken 获取未授权的Request Token
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Oauth#%E8%8E%B7%E5%8F%96%E6%9C%AA%E6%8E%88%E6%9D%83%E7%9A%84request-token
func (o *OAuth) GetRequestToken(ctx context.Context) (*RequestToken, error) {

	params := o.baseParams()

	req := &request{
		debug:       o.debug,
		method:      http.MethodGet,
		url:         o.ServiceProvider.RequestTokenURL,
		oauthParams: &params,
		ctx:         ctx,
	}

	if o.ServiceProvider.CallBackURL != "" {
		req.oauthParams.Set(CALLBACK_PARAM, o.ServiceProvider.CallBackURL)
	}

	if o.debug {
		utils.ShowInfomation("GetRequestToken Request", req)
	}

	if err := o.signRequest(req, ""); err != nil {
		return nil, err
	}

	body, err := req.Send(o.client)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &RequestToken{
		Token:  values.Get("oauth_token"),
		Secret: values.Get("oauth_token_secret"),
	}, nil
}

// GetAuthorizationURL 请求用户授权Request Token
//
// https://github.com/FanfouAPI/FanFouAPIDoc/wiki/Oauth#%E8%AF%B7%E6%B1%82%E7%94%A8%E6%88%B7%E6%8E%88%E6%9D%83request-token
func (o *OAuth) GetAuthorizationURL(token *RequestToken) (string, error) {
	if token == nil || token.Token == "" {
		return "", errors.WithStack(fmt.Errorf("miss request token"))
	}

	q := url.Values{}
	q.Set(TOKEN_PARAM, token.Token)
	q.Set(CALLBACK_PARAM, o.ServiceProvider.CallBackURL)

	return o.ServiceProvider.AuthorizeTokenURL + "?" + q.Encode(), nil
}

// GetAccessToken token 只是授权此 requestToken，需要用此 token 换取 AccessToken
func (o *OAuth) GetAccessToken(ctx context.Context, token RequestToken, verificationCode string, additionalParams map[string]string) (*AccessToken, error) {
	params := o.baseParams()
	req := &request{
		debug:       o.debug,
		method:      http.MethodGet,
		url:         o.ServiceProvider.AccessTokenURL,
		oauthParams: &params,
		ctx:         ctx,
	}

	for k, v := range additionalParams {
		req.oauthParams.Set(k, v)
	}

	req.oauthParams.Set(TOKEN_PARAM, token.Token)
	if verificationCode != "" {
		req.oauthParams.Set(VERIFIER_PARAM, verificationCode)
	}

	if o.debug {
		utils.ShowInfomation("GetAccessToken Request", req)
	}

	if err := o.signRequest(req, token.Secret); err != nil {
		return nil, err
	}

	body, err := req.Send(o.client)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &AccessToken{
		Token:  values.Get("oauth_token"),
		Secret: values.Get("oauth_token_secret"),
	}, nil
}

func (o *OAuth) baseParams() url.Values {

	q := url.Values{}
	q.Set(VERSION_PARAM, OAUTH_VERSION)
	q.Set(SIGNATURE_METHOD_PARAM, "HMAC-SHA1")
	q.Set(TIMESTAMP_PARAM, strconv.FormatInt(time.Now().Unix(), 10))
	q.Set(NONCE_PARAM, randString2(32))
	q.Set(CONSUMER_KEY_PARAM, o.consumerKey)
	return q
}

func baseString(method string, uri string, q url.Values) string {

	uri = strings.Replace(uri, "https://api.fanfou.com", "http://api.fanfou.com", -1)
	uri = strings.Replace(uri, "https://fanfou.com", "http://fanfou.com", -1)
	method = strings.ToUpper(method)
	paramsString := sortValue(q, "%s=%s")

	return strings.Join([]string{method, escape(uri), escape(paramsString)}, "&")
}

type request struct {
	debug       bool
	method      string
	url         string
	oauthParams *url.Values
	userParams  *url.Values
	ctx         context.Context
}

func (r *request) Send(c *http.Client) (string, error) {
	if c == nil {
		return "", errors.WithStack(fmt.Errorf("miss http client"))
	}

	rep, err := r.buildRequest()
	if err != nil {
		return "", err
	}

	if r.debug {
		utils.ShowInfomation("OAuth Request", rep)
	}

	resp, err := c.Do(rep)
	if err != nil {
		return "", errors.WithStack(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.WithStack(err)
	}
	defer resp.Body.Close()

	if r.debug {
		utils.ShowInfomation("OAuth Response Body", string(body))
	}

	return string(body), r.handleError(resp, body)
}

func (r *request) buildRequest() (*http.Request, error) {
	if r.method == "" {
		r.method = http.MethodGet
	}

	var (
		req *http.Request
		err error
	)

	if r.method == http.MethodPost {
		req, err = http.NewRequestWithContext(r.ctx, r.method, r.url, strings.NewReader(r.oauthParams.Encode()))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequestWithContext(r.ctx, r.method, r.url, nil)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		req.URL.RawQuery = r.oauthParams.Encode()
	}

	return req, nil
}

func (r *request) handleError(resp *http.Response, body []byte) error {
	if resp.StatusCode == http.StatusOK {
		return nil
	}

	e := &utils.ErrorResponse{}
	d := xml.NewDecoder(bytes.NewReader(body))
	d.Strict = false

	err := d.Decode(e)
	if err != nil {
		return errors.WithStack(err)
	}
	e.StatusCode = resp.StatusCode
	return errors.WithStack(e)
}

func (o *OAuth) signRequest(req *request, key string) error {

	bs := baseString(req.method, req.url, *req.oauthParams)
	signature := o.hmacSha1(key, bs)
	req.oauthParams.Set(SIGNATURE_PARAM, signature)

	if o.debug {
		utils.ShowInfomation("OAuth BaseString", bs)
	}

	return nil
}

func (o *OAuth) hmacSha1(key, message string) string {
	key = escape(o.consumerSecret) + "&" + escape(key)
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(message))
	rawSignature := h.Sum(nil)
	return base64.StdEncoding.EncodeToString(rawSignature)
}

// RequestToken RequestToken
type RequestToken struct {
	Token  string
	Secret string
}

// AccessToken AccessToken
type AccessToken struct {
	Token  string
	Secret string
}
