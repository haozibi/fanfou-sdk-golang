# fanfou-sdk-golang

![logo](logo.png)

[![CircleCI](https://circleci.com/gh/haozibi/fanfou-sdk-golang.svg?style=svg)](https://circleci.com/gh/haozibi/fanfou-sdk-golang) [![Coverage Status](https://coveralls.io/repos/github/haozibi/fanfou-sdk-golang/badge.svg)](https://coveralls.io/github/haozibi/fanfou-sdk-golang) [![GoDoc](https://godoc.org/github.com/haozibi/fanfou-sdk-golang?status.svg)](https://godoc.org/github.com/haozibi/fanfou-sdk-golang) [![pkg.go.dev](https://img.shields.io/badge/pkg.go.dev-haozibi%2Ffanfou--sdk--golang-blue)](https://pkg.go.dev/github.com/haozibi/fanfou-sdk-golang) [![Go Report Card](https://goreportcard.com/badge/github.com/haozibi/fanfou-sdk-golang)](https://goreportcard.com/report/github.com/haozibi/fanfou-sdk-golang) [![](https://img.shields.io/github/license/haozibi/fanfou-sdk-golang)](https://github.com/haozibi/fanfou-sdk-golang/LICENSE)


Fanfou SDK for Golang 

[README.md](README.md) | [中文 README.md](README_zh.md)

## Features

- Simple API and Error wrapping
- Supports OAuth and XAuth
- Use secure HTTPS protocol

**[Fanfou API Docs](https://github.com/FanfouAPI/FanFouAPIDoc/wiki)**

## Installation

```console
$ go get -u github.com/haozibi/fanfou-sdk-golang
```

## Example


Use Environment `FANFOU_SDK_DEBUG` to enable debug mode:

```console
$ FANFOU_SDK_DEBUG=true go run examples/status/main.go
```

### OAuth Authentication

```console
$ go run example/oauth/main.go\
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>
```

### OAuth oob Authentication

```console
$ go run example/oauth_oob/main.go\
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>
```

### Send Status With XAuth Authentication

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>\
    -status "在这里，无人知晓你是谁"
```

### Send Picture Status With XAuth Authentication

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>\
    -status "在这里，无人知晓你是谁"\
    -picture https://i.loli.net/2019/12/12/gYRoF4exsyNCc5J.jpg
```

### PublicTimeline With XAuth

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>
```

*More Example* **[Example](example/)**

## Error Handling

- All Error Wrap By "github.com/pkg/errors"
- API Response Error use `ErrorResponse` struct

```go
import pkgErr "github.com/pkg/errors"

func TestError(t *testing.T) {

	var f fanfou.Fanfou

	_, err := f.AccountService.VerifyCredentials(context.Background())
	if err != nil {
		err = pkgErr.Cause(err)
		var e *utils.ErrorResponse
		if errors.As(err, &e) {
			fmt.Println(e.GetStatusCode())
			fmt.Println(e.GetErrorMsg())
			fmt.Println(e.GetRequest())
		} else {
			fmt.Println(e)
		}
	}
}
```

## Important

- **Do not** support *callback*
- Only API v1

## License

MIT
