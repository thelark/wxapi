package wxapi

import "reflect"

type cgiBinExpressBusinessQuota struct {
	AccessToken string
}

func (t *cgiBinExpressBusinessQuota) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Get 获取电子面单余额。仅在使用加盟类快递公司时，才可以调用。
// POST https://api.weixin.qq.com/cgi-bin/express/business/quota/get?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusinessQuota) Get() {

}
