package wxapi

import "reflect"

type cgiBinExpressDeliveryPath struct {
	AccessToken string
}

func (t *cgiBinExpressDeliveryPath) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Update 更新运单轨迹
// POST https://api.weixin.qq.com/cgi-bin/express/delivery/path/update?access_token=ACCESS_TOKEN
func (t *cgiBinExpressDeliveryPath) Update() {

}
