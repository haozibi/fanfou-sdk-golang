# fanfou-sdk-golang

![logo](logo.png)

[![GoDoc](https://godoc.org/github.com/haozibi/fanfou-sdk-golang?status.svg)](https://godoc.org/github.com/haozibi/fanfou-sdk-golang) [![pkg.go.dev](https://img.shields.io/badge/pkg.go.dev-haozibi%2Ffanfou--sdk--golang-blue)](https://pkg.go.dev/github.com/haozibi/fanfou-sdk-golang) [![](https://img.shields.io/github/license/haozibi/fanfou-sdk-golang)](https://github.com/haozibi/fanfou-sdk-golang/LICENSE) 

Fanfou SDK for Golang

## 特点

- 优雅的接口设计和错误包装
- 支持 OAuth 和 XAuth 两种验证方式
- 全部接口使用安全的 HTTPS 协议


饭否 API 文档: [Fanfou API Docs](https://github.com/FanfouAPI/FanFouAPIDoc/wiki)

## 安装

```console
$ go get -u github.com/haozibi/fanfou-sdk-golang
```

## 示例

使用环境变量 `FANFOU_SDK_DEBUG` 开启调试模式，例如:

```console
$ FANFOU_SDK_DEBUG=true go run examples/status/main.go
```

### OAuth 认证

```console
$ go run example/oauth/main.go\
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>
```

### OAuth oob 认证

```console
$ go run example/oauth_oob/main.go\
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>
```

### 使用 XAuth 发送 Status 

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>\
    -status "在这里，无人知晓你是谁"
```

### 使用 XAuth 发送带图片的 Status

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>\
    -status "在这里，无人知晓你是谁"\
    -picture https://i.loli.net/2019/12/12/gYRoF4exsyNCc5J.jpg
```

### 使用 XAuth 访问“随便看看”

```console
$ go run examples/status/main.go \
    -consumerKey <consumerKey>\
    -consumerSecret <consumerSecret>\
    -password <password>\
    -username <username>
```

*更多使用示例请参考* **[Example](example/)**

## 错误处理

- 所有错误都使用 "github.com/pkg/errors" 包装
- API 响应错误都是 `ErrorResponse` 结构

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

## 注意:

- 所有接口都不支持 callback 参数
- 只实现了稳定的 v1 接口

## License

MIT
