package wxapi

import "reflect"

type cgiBinMessageWxOpenUpDataBleMsg struct {
	AccessToken string
}

func (t *cgiBinMessageWxOpenUpDataBleMsg) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
type cgiBinMessageWxOpenUpDataBleMsgSend struct {
}

// Send 修改被分享的动态消息。详见动态消息。
// POST https://api.weixin.qq.com/cgi-bin/message/wxopen/updatablemsg/send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageWxOpenUpDataBleMsg) Send() {

}
