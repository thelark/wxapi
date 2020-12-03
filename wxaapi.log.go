package wxapi

import "reflect"

type wxaApiLog struct {
	AccessToken string
}

func (t *wxaApiLog) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// GetClientVersion 获取客户端版本
// GET https://api.weixin.qq.com/wxaapi/log/get_client_version?access_token=ACCESS_TOKEN
func (t *wxaApiLog) GetClientVersion() {

}

//      --------------------------------------------------------------------

// GetPerformance 性能监控
// POST https://api.weixin.qq.com/wxaapi/log/get_performance?access_token=ACCESS_TOKEN
func (t *wxaApiLog) GetPerformance() {

}

//      --------------------------------------------------------------------

// GetScene 获取访问来源
// GET https://api.weixin.qq.com/wxaapi/log/get_scene?access_token=ACCESS_TOKEN
func (t *wxaApiLog) GetScene() {

}

//      --------------------------------------------------------------------

// JsErrSearch 错误查询
// POST https://api.weixin.qq.com/wxaapi/log/jserr_search?access_token=ACCESS_TOKEN
func (t *wxaApiLog) JsErrSearch() {

}
