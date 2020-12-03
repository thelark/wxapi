package wxapi

import "reflect"

type cgiBinMessageWxOpenActivityID struct {
	AccessToken string
}

func (t *cgiBinMessageWxOpenActivityID) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
type cgiBinMessageWxOpenActivityIDCreate struct {
}

// Create 创建被分享动态消息的 activity_id。详见动态消息。
// GET https://api.weixin.qq.com/cgi-bin/message/wxopen/activityid/create?access_token=ACCESS_TOKEN
func (t *cgiBinMessageWxOpenActivityID) Create() {

}
