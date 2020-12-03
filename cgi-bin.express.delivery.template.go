package wxapi

import "reflect"

type cgiBinExpressDeliveryTemplate struct {
	AccessToken string
}

func (t *cgiBinExpressDeliveryTemplate) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Preview 预览面单模板。用于调试面单模板使用。
// POST https://api.weixin.qq.com/cgi-bin/express/delivery/template/preview?access_token=ACCESS_TOKEN
func (t *cgiBinExpressDeliveryTemplate) Preview() {

}
