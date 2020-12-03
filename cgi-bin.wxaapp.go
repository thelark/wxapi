package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinWXAApp struct {
	AccessToken string
}

func (t *cgiBinWXAApp) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
type cgiBinWXAAppCreateWXAQRCode struct {
	*ErrorReturn
	Buffer []byte `json:"buffer"` // 返回的图片 Buffer
}

// CreateWXAQRCode 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
// POST https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=ACCESS_TOKEN
func (t *cgiBinWXAApp) CreateWXAQRCode(body struct {
	Path  string `json:"path"`            // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width int    `json:"width,omitempty"` // 二维码的宽度，单位 px。最小 280px，最大 1280px
}) (*cgiBinWXAAppCreateWXAQRCode, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinWXAAppCreateWXAQRCode)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/createwxaqrcode", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithBody(buf),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
