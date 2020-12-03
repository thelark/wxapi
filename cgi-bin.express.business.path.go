package wxapi

import "reflect"

type cgiBinExpressBusinessPath struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessPath) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Get 查询运单轨迹
// POST https://api.weixin.qq.com/cgi-bin/express/business/path/get?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessPath) Get() {

}
