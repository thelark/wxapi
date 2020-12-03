package wxapi

import "reflect"

type wxaApiUserLog struct {
	AccessToken string
}

func (t *wxaApiUserLog) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// UserLogSearch 实时日志查询
// GET https://api.weixin.qq.com/wxaapi/userlog/userlog_search?access_token=ACCESS_TOKEN
func (t *wxaApiUserLog) UserLogSearch() {

}
