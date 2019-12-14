package fanfou

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/haozibi/fanfou-sdk-golang/oauth"
	"github.com/haozibi/fanfou-sdk-golang/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var testDebug bool

func TestRequestBuildRequest(t *testing.T) {

	BaseURL = ""
	buildInput := func(q url.Values) map[string]string {
		m := make(map[string]string, len(q))

		for k := range q {
			m[k] = q.Get(k)
		}
		return m
	}

	tests := []struct {
		method string
		path   string
		query  url.Values
		reader io.Reader
		ctx    context.Context
		err    error
	}{
		{
			"GET",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			nil,
			context.Background(),
			nil,
		},
		{
			"",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			nil,
			context.Background(),
			nil,
		},
		{
			"POST",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			nil,
			context.Background(),
			nil,
		},
		{
			"POST",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			strings.NewReader("123"),
			context.Background(),
			nil,
		},
		{
			"POST",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			strings.NewReader("123"),
			context.WithValue(context.Background(), fileRequestKey, "abc.jpg"),
			nil,
		},
		{
			"POST",
			"/abc",
			url.Values{
				"a": []string{"1"},
				"b": []string{"2"},
			},
			strings.NewReader("123"),
			context.WithValue(context.Background(), fileRequestKey, 123),
			nil,
		},
	}

	for _, v := range tests {

		req := &request{
			HTTPMethod: v.method,
			HTTPPath:   v.path,
			Input:      buildInput(v.query),
			Context:    v.ctx,
			Reader:     v.reader,
		}

		var r *http.Request
		var err error
		if v.reader == nil {
			err = req.buildRequest()
			r = req.req
			r.ParseForm()
			assert.Equal(t, r.Form, v.query)
		} else {
			err = req.buildFileRequest(req.Reader)
			r = req.req
			r.ParseMultipartForm(1 << 20)
			assert.Equal(t, r.MultipartForm.Value, map[string][]string(v.query))
		}

		assert.Equal(t, err, v.err)

		if v.method == "" {
			assert.Equal(t, r.Method, http.MethodGet)
		} else {
			assert.Equal(t, r.Method, v.method)
		}

		assert.Equal(t, r.URL.Path, v.path)
	}

}

func TestRequestSend(t *testing.T) {

	tests := []struct {
		desc    string
		handler http.HandlerFunc
		req     *request
		err     error
	}{
		{
			"错误: miss context",
			func(w http.ResponseWriter, r *http.Request) {
				return
			},
			&request{
				Debug:      testDebug,
				HTTPMethod: "GET",
			},
			fmt.Errorf("miss context.Context"),
		},
		{
			"错误: 返回错误",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, `{"request":"search/public_timeline.json","error":"参数错误"}`)
				return
			},
			&request{
				Debug:      testDebug,
				HTTPMethod: "GET",
				HTTPPath:   "search/public_timeline.json",
				Input:      map[string]string{"a": "1"},
				Context:    context.Background(),
			},
			&utils.ErrorResponse{
				StatusCode: 500,
				Request:    "search/public_timeline.json", ErrorMsg: "参数错误",
			},
		},
		{
			"正确: 不带参数",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `[{"created_at":"Sat Dec 14 08:35:42 +0000 2019"},{"created_at":"Sat Dec 14 08:35:42 +0000 2019"}]`)
				return
			},
			&request{
				Debug:      testDebug,
				HTTPMethod: "GET",
				HTTPPath:   "search/public_timeline.json",
				Input:      map[string]string{"a": "1"},
				Context:    context.Background(),
				Output: func() *[]Status {
					s := make([]Status, 0)
					return &s
				}(),
			},
			nil,
		},
		{
			"正确: 带参数",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, `[{"created_at":"Sat Dec 14 08:35:42 +0000 2019"},{"created_at":"Sat Dec 14 08:35:42 +0000 2019"}]`)
				return
			},
			&request{
				Debug:      testDebug,
				HTTPMethod: "GET",
				HTTPPath:   "search/public_timeline.json",
				Input:      map[string]string{"a": "1", "format": "html"},
				Context:    context.Background(),
				Output: func() *[]Status {
					s := make([]Status, 0)
					return &s
				}(),
			},
			nil,
		},
	}

	for k, v := range tests {
		ts := httptest.NewServer(v.handler)
		BaseURL = ts.URL + "/"
		err := v.req.send(ts.Client())
		assert.Equalf(t, errors.Cause(err), v.err, "desc: %s, id: %d", v.desc, k)
		ts.Close()
	}

}

func TestCleanOption(t *testing.T) {

	tests := []struct {
		params map[string]string
		path   string
		clean  []string
	}{
		{
			map[string]string{"a": "1", "b": "2"},
			"search/public_timeline.json",
			[]string{"a", "b"},
		},
		{
			map[string]string{},
			"search/public_timeline.json",
			nil,
		},
		{
			map[string]string{"a": "1", "b": "2"},
			"favorites/destroy/abcaAFAGi.json",
			[]string{"a", "b"},
		},
	}

	for k, v := range tests {

		got := cleanOption(v.params, v.path)

		sort.Strings(got)
		sort.Strings(v.clean)

		assert.Equalf(t, v.clean, got, "id: %d, path: %s", k, v.path)
	}
}

func TestFanfouRequestToken(t *testing.T) {

	const (
		consumerkey             = "test_consumer_key"
		consumerSecret          = "test_consumer_secret"
		requestOAuthToken       = "8ldIZyxQeVrFZXFOZH5tAwj6vzJYuLQpl0WUEYtWc"
		requestOAuthTokenSecret = "x6qpRnlEmW9JbQn4PQVVeVG8ZLPEx6A0TOebgwcuA"
		accessToken             = "891212-3MdIZyxPeVsFZXFOZj5tAwj6vzJYuLQplzWUmYtWd"
		accessTokenSecret       = "x6qpRnlEmW9JbQn4PQVVeVG8ZLPEx6A0TOebgwcuA"
	)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := strings.TrimPrefix(r.URL.Path, "/")

		switch path {
		case requestTokenURI:
			fmt.Fprintf(w, "oauth_token=%s&oauth_token_secret=%s", requestOAuthToken, requestOAuthTokenSecret)
		case authorizeTokenURI:
		case accessTokenURI:
			fmt.Fprintf(w, "oauth_token=%s&oauth_token_secret=%s", accessToken, accessTokenSecret)
		}

	}))
	defer ts.Close()
	AuthBaseURL = ts.URL + "/"

	f := NewFanFouWithClient(consumerkey, consumerSecret, ts.Client())

	// oAuth

	token, err := f.RequestToken(context.Background(), "")
	if err != nil {
		t.Errorf("%+v", err)
	}

	assert.Equal(t, requestOAuthToken, token.Token)
	assert.Equal(t, requestOAuthTokenSecret, token.Secret)

	uri, err := f.AuthorizationURL(token)
	if err != nil {
		t.Errorf("%+v", err)
	}

	t.Log("Jump URL:", uri)

	acToken, err := f.AccessToken(context.Background(),
		oauth.RequestToken{
			Token:  token.Token,
			Secret: token.Secret,
		}, "", nil)

	assert.Equal(t, accessToken, acToken.Token)
	assert.Equal(t, accessTokenSecret, acToken.Secret)

	f.OAuth(acToken)

	// xAuth
	f.XAuth("user", "pass")
}

func TestOption(t *testing.T) {

	m := make(map[string]string)

	opts := []Option{
		WithID("id_abc"),
		WithSource("source_abc"),
		WithLocation("location_abc"),
		WithLiteMode(),
		WithHTMLFormat(),
		WithCount(1),
		WithSinceID(2),
		WithMaxID(3),
		WithPage(4),
		WithStatus("status_abc"),
		WithURL("url_abc"),
		WithDescription("desc_abc"),
		WithName("name_abc"),
		WithEmail("email_abc"),
		WithInReplyToStatusID("in_reply_to_status_id_abc"),
		WithInReplyToUserID("in_reply_to_user_id_abc"),
		WithRepostStatusID("repost_status_id_abc"),
		WithInReplyToID("in_reply_to_id_abc"),
	}

	for _, v := range opts {
		v.apply(m)
	}

	assert.Equal(t, m["id"], "id_abc")
	assert.Equal(t, m["source"], "source_abc")
	assert.Equal(t, m["location"], "location_abc")
	assert.Equal(t, m["mode"], "lite")
	assert.Equal(t, m["format"], "html")
	assert.Equal(t, m["count"], "1")
	assert.Equal(t, m["since_id"], "2")
	assert.Equal(t, m["max_id"], "3")
	assert.Equal(t, m["page"], "4")
	assert.Equal(t, m["status"], "status_abc")
	assert.Equal(t, m["url"], "url_abc")
	assert.Equal(t, m["description"], "desc_abc")
	assert.Equal(t, m["name"], "name_abc")
	assert.Equal(t, m["email"], "email_abc")
	assert.Equal(t, m["in_reply_to_status_id"], "in_reply_to_status_id_abc")
	assert.Equal(t, m["in_reply_to_user_id"], "in_reply_to_user_id_abc")
	assert.Equal(t, m["repost_status_id"], "repost_status_id_abc")
	assert.Equal(t, m["in_reply_to_id"], "in_reply_to_id_abc")

	assert.Equal(t, len(opts), len(m))

	// just test

	WithCallback("haozibi")
}
