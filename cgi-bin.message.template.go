package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinMessageTemplate struct {
	AccessToken string
}

func (t *cgiBinMessageTemplate) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

type MiniProgram struct {
	AppId    string `json:"appid"`    // 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	PagePath string `json:"pagepath"` // 所需跳转到小程序的具体页面路径，支持带参数,（示例index?foo=bar），要求该小程序已发布，暂不支持小游戏
}
type Data struct {
	Value string `json:"value"`
	Color string `json:"color,omitempty"`
}

// 方法 --------------------------------------------------------------------

type cgiBinMessageTemplateSend struct {
	*ErrorReturn
	MsgId int `json:"msgid"`
}

// Send 发送模板消息
// POST https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageTemplate) Send(req struct {
	ToUser      string          `json:"touser"`                // 接收者openid
	TemplateId  string          `json:"template_id"`           // 模板ID
	Url         string          `json:"url,omitempty"`         // 模板跳转链接（海外帐号没有跳转能力）
	MiniProgram *MiniProgram    `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        map[string]Data `json:"data"`                  // 模板数据
}) (*cgiBinMessageTemplateSend, error) {
	rsp := new(cgiBinMessageTemplateSend)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Post(
		fmt.Sprintf("%s/send", getBasePath()),
		request.WithBody(body),
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

//      --------------------------------------------------------------------

type cgiBinMessageTemplateSubscribe struct {
	*ErrorReturn
	MsgId int `json:"msgid"`
}

// Subscribe 通过API推送订阅模板消息给到授权微信用户
// POST https://api.weixin.qq.com/cgi-bin/message/template/subscribe?access_token=ACCESS_TOKEN
func (t *cgiBinMessageTemplate) Subscribe(req struct {
	ToUser      string          `json:"touser"`                // 填接收消息的用户openid
	TemplateId  string          `json:"template_id"`           // 订阅消息模板ID
	Url         string          `json:"url,omitempty"`         // 点击消息跳转的链接，需要有ICP备案
	MiniProgram *MiniProgram    `json:"miniprogram,omitempty"` // 跳小程序所需数据，不需跳小程序可不用传该数据
	Scene       string          `json:"scene"`                 // 订阅场景值
	Title       string          `json:"title"`                 // 消息标题，15字以内
	Data        map[string]Data `json:"data"`                  // 消息正文，value为消息内容文本（200字以内），没有固定格式，可用\n换行，color为整段消息内容的字体颜色（目前仅支持整段消息为一种颜色）
}) (*cgiBinMessageTemplateSubscribe, error) {
	rsp := new(cgiBinMessageTemplateSubscribe)
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Post(
		fmt.Sprintf("%s/subscribe", getBasePath()),
		request.WithBody(body),
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
