package wxapi

import "reflect"

type customServiceKFSession struct {
	AccessToken string
}

func (t *customServiceKFSession) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Create 创建会话(此接口在客服和用户之间创建一个会话，如果该客服和用户会话已存在，则直接返回0。指定的客服帐号必须已经绑定微信号且在线。)
// POST https://api.weixin.qq.com/customservice/kfsession/create?access_token=ACCESS_TOKEN
func (t *customServiceKFSession) Create() {

}

//      --------------------------------------------------------------------
// Close 关闭会话
// POST https://api.weixin.qq.com/customservice/kfsession/close?access_token=ACCESS_TOKEN
func (t *customServiceKFSession) Close() {

}

//      --------------------------------------------------------------------
// GetSession 获取客户会话状态
// GET https://api.weixin.qq.com/customservice/kfsession/getsession?access_token=ACCESS_TOKEN&openid=OPENID
func (t *customServiceKFSession) GetSession() {

}

//      --------------------------------------------------------------------
// GetSessionList 获取客服会话列表
// GET https://api.weixin.qq.com/customservice/kfsession/getsessionlist?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
func (t *customServiceKFSession) GetSessionList() {

}

//      --------------------------------------------------------------------
// GetWaitCase 获取未接入会话列表
// GET https://api.weixin.qq.com/customservice/kfsession/getwaitcase?access_token=ACCESS_TOKEN
func (t *customServiceKFSession) GetWaitCase() {

}
