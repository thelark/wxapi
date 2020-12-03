package wxapi

import "reflect"

type cgiBinMediaVoice struct {
	AccessToken string
}

func (t *cgiBinMediaVoice) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// AddVoiceToRECOForText 提交语音
// POST http://api.weixin.qq.com/cgi-bin/media/voice/addvoicetorecofortext?access_token=ACCESS_TOKEN&format=&voice_id=xxxxxx&lang=zh_CN
func (t *cgiBinMediaVoice) AddVoiceToRECOForText() {

}

//      --------------------------------------------------------------------
// QueryRECOResultForText 获取语音识别结果
// POST http://api.weixin.qq.com/cgi-bin/media/voice/queryrecoresultfortext?access_token=ACCESS_TOKEN&voice_id=xxxxxx&lang=zh_CN
func (t *cgiBinMediaVoice) QueryRECOResultForText() {

}

//      --------------------------------------------------------------------
// TranslateContent 微信翻译
// POST http://api.weixin.qq.com/cgi-bin/media/voice/translatecontent?access_token=ACCESS_TOKEN&lfrom=xxx&lto=xxx
func (t *cgiBinMediaVoice) TranslateContent() {

}
