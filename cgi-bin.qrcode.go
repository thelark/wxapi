package wxapi

import "reflect"

type cgiBinQRCode struct {
	AccessToken string
}

func (t *cgiBinQRCode) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Create 创建二维码
// POST https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=TOKEN
func (t *cgiBinQRCode) Create() {

}
