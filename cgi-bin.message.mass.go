package wxapi

import "reflect"

type cgiBinMessageMass struct {
	AccessToken string
}

func (t *cgiBinMessageMass) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// SendAll 根据标签进行群发【订阅号与服务号认证后均可用】
// POST https://api.weixin.qq.com/cgi-bin/message/mass/sendall?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) SendAll() {

}

//      --------------------------------------------------------------------
// Send 根据OpenID列表群发【订阅号不可用，服务号认证后可用】
// POST https://api.weixin.qq.com/cgi-bin/message/mass/send?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) Send() {

}

//      --------------------------------------------------------------------
// Delete 删除群发【订阅号与服务号认证后均可用】
// POST https://api.weixin.qq.com/cgi-bin/message/mass/delete?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) Delete() {

}

//      --------------------------------------------------------------------
// Preview 预览接口【订阅号与服务号认证后均可用】
// POST https://api.weixin.qq.com/cgi-bin/message/mass/preview?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) Preview() {

}

//      --------------------------------------------------------------------
// Get 查询群发消息发送状态【订阅号与服务号认证后均可用】
// POST https://api.weixin.qq.com/cgi-bin/message/mass/get?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) Get() {

}

//      --------------------------------------------------------------------
// SpeedGet 控制群发速度
// POST https://api.weixin.qq.com/cgi-bin/message/mass/speed/get?access_token=ACCESS_TOKEN
func (t *cgiBinMessageMass) SpeedGet() {

}
