package wxapi

import "reflect"

type cgiBinExpressBusinessPrinter struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessPrinter) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// GetAll 获取打印员。若需要使用微信打单 PC 软件，才需要调用。
// GET https://api.weixin.qq.com/cgi-bin/express/business/printer/getall?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessPrinter) GetAll() {

}

//      --------------------------------------------------------------------
// Update 配置面单打印员，可以设置多个，若需要使用微信打单 PC 软件，才需要调用。
// POST https://api.weixin.qq.com/cgi-bin/express/business/printer/update?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessPrinter) Update() {

}
