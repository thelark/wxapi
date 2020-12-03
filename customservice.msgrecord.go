package wxapi

import "reflect"

type customServiceMsgRecord struct {
	AccessToken string
}

func (t *customServiceMsgRecord) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// GetMsgList 获取聊天记录
// POST https://api.weixin.qq.com/customservice/msgrecord/getmsglist?access_token=ACCESS_TOKEN
func (t *customServiceMsgRecord) GetMsgList() {

}
