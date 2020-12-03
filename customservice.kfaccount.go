package wxapi

import "reflect"

type customServiceKFAccount struct {
	AccessToken string
}

func (t *customServiceKFAccount) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// Add 添加客服帐号
// POST https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN
func (t *customServiceKFAccount) Add() {

}

//      --------------------------------------------------------------------
// Update 修改客服帐号
// POST https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN
func (t *customServiceKFAccount) Update() {

}

//      --------------------------------------------------------------------
// Del 删除客服帐号
// POST https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN
func (t *customServiceKFAccount) Del() {

}

//      --------------------------------------------------------------------
// UploadHeadImg 设置客服帐号的头像
// POST https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=KFACCOUNT
func (t *customServiceKFAccount) UploadHeadImg() {

}

//      --------------------------------------------------------------------
// InviteWorker 邀请绑定客服帐号
// POST https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN
func (t *customServiceKFAccount) InviteWorker() {

}
