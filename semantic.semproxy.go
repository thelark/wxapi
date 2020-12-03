package wxapi

import "reflect"

type semanticSemProxy struct {
	AccessToken string
}

func (t *semanticSemProxy) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Search 语义理解(微信开放平台语义理解接口调用（http请求）简单方便，用户无需掌握语义理解及相关技术，只需根据自己的产品特点，选择相应的服务即可搭建一套智能语义服务。)
// POST https://api.weixin.qq.com/semantic/semproxy/search?access_token=YOUR_ACCESS_TOKEN
func (t *semanticSemProxy) Search() {

}
