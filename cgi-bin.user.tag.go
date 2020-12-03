package wxapi

import "reflect"

type cgiBinUserTag struct {
	AccessToken string
}

func (t *cgiBinUserTag) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Get 获取标签下粉丝列表
// GET https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=ACCESS_TOKEN
func (t *cgiBinUserTag) Get() {

}
