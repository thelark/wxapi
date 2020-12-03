package wxapi

import "reflect"

type cgiBinMessageWxOpenTemplate struct {
	AccessToken string
}

func (t *cgiBinMessageWxOpenTemplate) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
type cgiBinMessageWxOpenTemplateUniformSend struct {
}

// UniformSend 下发小程序和公众号统一的服务消息
// POST https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageWxOpenTemplate) UniformSend() {

}
