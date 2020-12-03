package wxapi

import "reflect"

type cgiBinMedia struct {
	AccessToken string
}

func (t *cgiBinMedia) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinMedia) Voice(opts ...option) *cgiBinMediaVoice {
	self := &cgiBinMediaVoice{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

// UploadNews 上传图文消息素材【订阅号与服务号认证后均可用】
// POST https://api.weixin.qq.com/cgi-bin/media/uploadnews?access_token=ACCESS_TOKEN
func (t *cgiBinMedia) UploadNews() {

}

//      --------------------------------------------------------------------
// UploadImg 上传图文消息内的图片获取URL(本接口所上传的图片不占用公众号的素材库中图片数量的100000个的限制。图片仅支持jpg/png格式，大小必须在1MB以下。)
// POST https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=ACCESS_TOKEN
func (t *cgiBinMedia) UploadImg() {

}

//      --------------------------------------------------------------------
// Upload 把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
// POST https://api.weixin.qq.com/cgi-bin/media/upload?access_token=ACCESS_TOKEN&type=TYPE
func (t *cgiBinMedia) Upload() {

}

//      --------------------------------------------------------------------
// Get 获取客服消息内的临时素材。即下载临时的多媒体文件。目前小程序仅支持下载图片文件。
// GET https://api.weixin.qq.com/cgi-bin/media/get?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func (t *cgiBinMedia) Get() {

}

//      --------------------------------------------------------------------
// UploadVideo 把媒体文件上传到微信服务器。目前仅支持图片。用于发送客服消息或被动回复用户消息。
// POST https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token=ACCESS_TOKEN
func (t *cgiBinMedia) UploadVideo() {

}

//      --------------------------------------------------------------------
// GetJSSdk 高清语音素材获取接口(公众号可以使用本接口获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。)
// GET https://api.weixin.qq.com/cgi-bin/media/get/jssdk?access_token=ACCESS_TOKEN&media_id=MEDIA_ID
func (t *cgiBinMedia) GetJSSdk() {

}
