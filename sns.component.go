package wxapi

import (
	"fmt"
	"reflect"

	"github.com/thelark/request"
)

type snsComponent struct {
	AppID                string // 公众号的 appid
	ComponentAppID       string // 服务开发方的 appid
	ComponentAccessToken string // 服务开发方的 access_token
}

func (t *snsComponent) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type snsComponentJsCode2Session struct {
	*ErrorReturn
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
}

// JsCode2Session 小程序登录
// GET https://api.weixin.qq.com/sns/component/jscode2session?appid=APPID&js_code=JSCODE&grant_type=authorization_code&component_appid=COMPONENT_APPID&component_access_token=COMPONENT_ACCESS_TOKEN
func (t *snsComponent) JsCode2Session(code string) (*snsComponentJsCode2Session, error) {
	rsp := new(snsComponentJsCode2Session)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/jscode2session", getBasePath()),
		request.WithParam("appid", t.AppID),
		request.WithParam("js_code", code),
		request.WithParam("grant_type", "authorization_code"),
		request.WithParam("component_appid", t.ComponentAppID),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
