package oauth

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randString() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func randString2(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func escape(s string) string {
	t := make([]byte, 0, 3*len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isEscapable(c) {
			t = append(t, '%')
			t = append(t, "0123456789ABCDEF"[c>>4])
			t = append(t, "0123456789ABCDEF"[c&15])
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isEscapable(b byte) bool {
	return !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' || '0' <= b && b <= '9' || b == '-' || b == '.' || b == '_' || b == '~')
}

func canonicalizeURL(u *url.URL) string {
	var buf bytes.Buffer
	buf.WriteString(u.Scheme)
	buf.WriteString("://")
	buf.WriteString(u.Host)
	buf.WriteString(u.Path)

	return buf.String()
}

func sortValue(params url.Values, format string) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	paramsList := make([]string, 0, len(params))
	for _, key := range keys {
		// `%s="%s"` or "%s=%s"
		paramsList = append(paramsList, fmt.Sprintf(format, escape(key), escape(params.Get(key))))
	}
	return strings.Join(paramsList, "&")
}
