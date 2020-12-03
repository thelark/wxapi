package wxapi

import "reflect"

type cgiBinExpressBusinessAccount struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessAccount) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Bind 绑定、解绑物流账号
// POST https://api.weixin.qq.com/cgi-bin/express/business/account/bind?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessAccount) Bind() {

}

//      --------------------------------------------------------------------
// GetAll 获取所有绑定的物流账号
// GET https://api.weixin.qq.com/cgi-bin/express/business/account/getall?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessAccount) GetAll() {

}
