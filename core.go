package wxapi

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/thelark/request"
)

const (
	// 通用域名 - 使用该域名将访问官方指定就近的接入点
	wxApiUrl = "api.weixin.qq.com"
	// 通用异地容灾域名 - 当上述域名不可访问时可改访问此域名
	wxApi2Url = "api2.weixin.qq.com"
	// 上海域名 - 使用该域名将访问上海的接入点
	wxShApiUrl = "sh.api.weixin.qq.com"
	// 深圳域名 - 使用该域名将访问深圳的接入点
	wxSzApiUrl = "sz.api.weixin.qq.com"
	// 香港域名 - 使用该域名将访问香港的接入点
	wxXgApiUrl = "hk.api.weixin.qq.com"
)

var wxRequest = request.New(wxApiUrl)

type response interface {
	Error() error
}

type ErrorReturn struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (err *ErrorReturn) Error() error {
	if err != nil {
		if err.ErrCode != 0 {
			return fmt.Errorf(err.ErrMsg)
		}
	}
	return nil
}
func checkError(rsp response) error {
	return rsp.Error()
}

type api interface {
	set(k, v string)
}

type option func(api)

func WithAppID(appID string) option {
	return func(self api) {
		self.set("AppID", appID)
	}
}
func WithAppSecret(appSecret string) option {
	return func(self api) {
		self.set("AppSecret", appSecret)
	}
}
func WithAccessToken(accessToken string) option {
	return func(self api) {
		self.set("AccessToken", accessToken)
	}
}
func WithComponentAppID(componentAppID string) option {
	return func(self api) {
		self.set("ComponentAppID", componentAppID)
	}
}
func WithComponentAccessToken(componentAccessToken string) option {
	return func(self api) {
		self.set("ComponentAccessToken", componentAccessToken)
	}
}

// 根据文件名称获取请求路由
func getBasePath() string {
	_, file, _, ok := runtime.Caller(1)
	if ok {
		path := strings.TrimSuffix(path.Base(file), path.Ext(file))
		path = strings.ReplaceAll(path, ".", "/")
		return path
	}
	return ""
}
