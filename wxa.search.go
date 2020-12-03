package wxapi

import "reflect"

type wxaSearch struct {
	AccessToken string
}

func (t *wxaSearch) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// WXAApiSubmitPages 小程序开发者可以通过本接口提交小程序页面url及参数信息，让微信可以更及时的收录到小程序的页面信息，开发者提交的页面信息将可能被用于小程序搜索结果展示。
// POST https://api.weixin.qq.com/wxa/search/wxaapi_submitpages?access_token=TOKEN
func (t *wxaSearch) WXAApiSubmitPages() {

}
