package wxapi

import "reflect"

type cgiBinCustomService struct {
	AccessToken string
}

func (t *cgiBinCustomService) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// GetKFList 获取所有客服账号
// GET https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN
func (t *cgiBinCustomService) GetKFList() {

}

//      --------------------------------------------------------------------

// GetOnlineKFList
// GET https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN
func (t *cgiBinCustomService) GetOnlineKFList() {

}
