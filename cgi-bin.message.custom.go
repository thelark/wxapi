package wxapi

import "reflect"

type cgiBinMessageCustom struct {
	AccessToken string
}

func (t *cgiBinMessageCustom) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// Send 发送客服消息给用户。详细规则见 发送客服消息
// POST https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageCustom) Send() {

}

//      --------------------------------------------------------------------

// Typing 下发客服当前输入状态给用户。详见 客服消息输入状态
// POST https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=ACCESS_TOKEN
func (t *cgiBinMessageCustom) Typing() {

}
