package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinMessageSubscribe struct {
	AccessToken string
}

func (t *cgiBinMessageSubscribe) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinMessageSubscribeSend struct {
	*ErrorReturn
}

// Send 发送订阅消息(POST)
// POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageSubscribe) Send(body struct {
	ToUser     string `json:"touser"`         // 接收者openid
	TemplateID string `json:"template_id"`    // 模板ID
	Page       string `json:"page,omitempty"` // 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo=bar）。该字段不填则模板无跳转。
	Data       map[string]struct {
		Value string `json:"value"`
	} `json:"data"` // 模板数据
	MiniProgramState string `json:"miniprogram_state,omitempty"` // 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版
	Lang             string `json:"lang,omitempty"`              // 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN
}) (*cgiBinMessageSubscribeSend, error) {
	rsp := new(cgiBinMessageSubscribeSend)
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Post(
		fmt.Sprintf("%s/send", getBasePath()),
		request.WithBody(buf),
		request.WithParam("access_token", t.AccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
