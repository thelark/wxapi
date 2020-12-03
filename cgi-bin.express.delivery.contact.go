package wxapi

import "reflect"

type cgiBinExpressDeliveryContact struct {
	AccessToken string
}

func (t *cgiBinExpressDeliveryContact) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Get 获取面单联系人信息
// POST https://api.weixin.qq.com/cgi-bin/express/delivery/contact/get?access_token=ACCESS_TOKEN
func (t *cgiBinExpressDeliveryContact) Get() {

}
