package oauth

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/haozibi/fanfou-sdk-golang/utils"

	"github.com/pkg/errors"
)

// RoundTripper Custom RoundTripper
type RoundTripper struct {
	debug bool
	oauth *OAuth
	token *AccessToken
}

// BuildClient build client with oauth
func (o *OAuth) BuildClient(accessToken *AccessToken) *http.Client {
	return &http.Client{
		Transport: &RoundTripper{
			debug: utils.IsDebug(),
			oauth: o,
			token: accessToken,
		},
	}
}

// RoundTrip RoundTrip
func (r *RoundTripper) RoundTrip(userRequest *http.Request) (*http.Response, error) {

	params := r.oauth.baseParams()
	oauthParams := params
	if r.token != nil && len(r.token.Token) > 0 {
		params.Set(TOKEN_PARAM, r.token.Token)
		oauthParams.Set(TOKEN_PARAM, r.token.Token)
	}

	otherParams, err := collectParameters(userRequest)
	if err != nil {
		return nil, err
	}

	for key := range otherParams {
		params.Set(key, otherParams.Get(key))
	}

	bs := baseString(userRequest.Method, canonicalizeURL(userRequest.URL), params)
	signature := r.oauth.hmacSha1(r.token.Secret, bs)

	if r.debug {
		utils.ShowInfomation("Fanfou Signed BaseString", bs)
	}

	oauthParams.Set(SIGNATURE_PARAM, signature)
	userRequest.Header.Set("Authorization", authHeaderValue(oauthParams))

	if r.debug {
		utils.ShowInfomation("Fanfou Signed Request", userRequest)
	}

	return r.oauth.client.Do(userRequest)
}

func collectParameters(req *http.Request) (url.Values, error) {

	params := url.Values{}

	query := req.URL.Query()
	for key := range query {
		params.Set(key, query.Get(key))
	}

	if req.Body != nil &&
		req.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		values, err := url.ParseQuery(string(b))
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for key := range values {
			params.Set(key, values.Get(key))
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	return params, nil
}

func authHeaderValue(params url.Values) string {

	paramsString := sortValue(params, `%s="%s"`)
	return "OAuth " + paramsString
}
