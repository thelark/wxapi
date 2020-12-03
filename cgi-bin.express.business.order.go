package wxapi

import "reflect"

type cgiBinExpressBusinessOrder struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessOrder) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Add 生成运单
// POST https://api.weixin.qq.com/cgi-bin/express/business/order/add?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessOrder) Add() {

}

//      --------------------------------------------------------------------
// BatchGet 批量获取运单数据
// POST https://api.weixin.qq.com/cgi-bin/express/business/order/batchget?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessOrder) BatchGet() {

}

//      --------------------------------------------------------------------
// Cancel 取消运单
// POST https://api.weixin.qq.com/cgi-bin/express/business/order/cancel?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessOrder) Cancel() {

}

//      --------------------------------------------------------------------
// Get 获取运单数据
// POST https://api.weixin.qq.com/cgi-bin/express/business/order/get?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessOrder) Get() {

}
