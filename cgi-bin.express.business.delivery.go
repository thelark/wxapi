package wxapi

import "reflect"

type cgiBinExpressBusinessDelivery struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessDelivery) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// GetAll 获取支持的快递公司列表
// GET https://api.weixin.qq.com/cgi-bin/express/business/delivery/getall?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessDelivery) GetAll() {

}
