package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type wxa struct {
	AccessToken string
}

func (t *wxa) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

type LineColor struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}

// 子节点 --------------------------------------------------------------------
func (t *wxa) Search(opts ...option) *wxaSearch {
	self := &wxaSearch{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

// AddNearbyPoi 添加地点
// POST https://api.weixin.qq.com/wxa/addnearbypoi?access_token=ACCESS_TOKEN
func (t *wxa) AddNearbyPoi() {

}

//      --------------------------------------------------------------------

// DelNearbyPoi 删除地点
// POST https://api.weixin.qq.com/wxa/delnearbypoi?access_token=ACCESS_TOKEN
func (t *wxa) DelNearbyPoi() {

}

//      --------------------------------------------------------------------

// DevPlugin 获取当前所有插件使用方(供插件开发者调用)|修改插件使用申请的状态（供插件开发者调用）
// POST https://api.weixin.qq.com/wxa/devplugin?access_token=TOKEN
func (t *wxa) DevPlugin() {

}

//      --------------------------------------------------------------------

// GetNearbyPoiList 查看地点列表
// GET https://api.weixin.qq.com/wxa/getnearbypoilist?page=1&page_rows=20&access_token=ACCESS_TOKEN
func (t *wxa) GetNearbyPoiList() {

}

//      --------------------------------------------------------------------

type wxaGetPaidUnionId struct {
	*ErrorReturn
	UnionID string `json:"unionid"` // 用户唯一标识，调用成功后返回
}

// GetPaidUnionId 用户支付完成后，获取该用户的 UnionId，无需用户授权。本接口支持第三方平台代理查询。
// GET https://api.weixin.qq.com/wxa/getpaidunionid?access_token=ACCESS_TOKEN&openid=OPENID
func (t *wxa) GetPaidUnionId(openid string) (*wxaGetPaidUnionId, error) {
	rsp := new(wxaGetPaidUnionId)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getpaidunionid", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("openid", openid),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

//      --------------------------------------------------------------------

type wxaGetWXACode struct {
	*ErrorReturn
	ContentType string `json:"contentType"` // 返回类型
	Buffer      []byte `json:"buffer"`      // 返回的图片 Buffer
}

// GetWXACode 获取小程序码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制，详见获取二维码。
// POST https://api.weixin.qq.com/wxa/getwxacode?access_token=ACCESS_TOKEN
func (t *wxa) GetWXACode(body struct {
	Path      string     `json:"path"`                 // 扫码进入的小程序页面路径，最大长度 128 字节，不能为空；对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"，即可在 wx.getLaunchOptionsSync 接口中的 query 参数获取到 {foo:"bar"}。
	Width     int        `json:"width,omitempty"`      // 二维码的宽度，单位 px。最小 280px，最大 1280px
	AutoColor bool       `json:"auto_color,omitempty"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调
	LineColor *LineColor `json:"line_color,omitempty"` // auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaLine bool       `json:"is_hyaline,omitempty"` // 是否需要透明底色，为 true 时，生成透明底色的小程序码
}) (*wxaGetWXACode, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaGetWXACode)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/getwxacode", getBasePath()),
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

//      --------------------------------------------------------------------
type wxaGetWXACodeUnLimit struct {
	*ErrorReturn
	Buffer []byte `json:"buffer"` // 返回的图片 Buffer
}

// GetWXACodeUnLimit 获取小程序码，适用于需要的码数量极多的业务场景。通过该接口生成的小程序码，永久有效，数量暂无限制。 更多用法详见 获取二维码。
// POST https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=ACCESS_TOKEN
func (t *wxa) GetWXACodeUnLimit(body struct {
	Scene     string     `json:"scene"`                // 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&'()*+,/:;=?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式）
	Page      string     `json:"page,omitempty"`       // 必须是已经发布的小程序存在的页面（否则报错），例如 pages/index/index, 根路径前不要填加 /,不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面
	Width     int        `json:"width,omitempty"`      // 二维码的宽度，单位 px，最小 280px，最大 1280px
	AutoColor bool       `json:"auto_color,omitempty"` // 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false
	LineColor *LineColor `json:"line_color,omitempty"` // auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {"r":"xxx","g":"xxx","b":"xxx"} 十进制表示
	IsHyaLine bool       `json:"is_hyaline,omitempty"` // 是否需要透明底色，为 true 时，生成透明底色的小程序
}) (*wxaGetWXACodeUnLimit, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaGetWXACodeUnLimit)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/getwxacodeunlimit", getBasePath()),
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

//      --------------------------------------------------------------------

// ImageSearch 本接口提供基于小程序的站内搜商品图片搜索能力
// POST https://api.weixin.qq.com/wxa/imagesearch?access_token=TOKEN
func (t *wxa) ImageSearch() {

}

//      --------------------------------------------------------------------

// ImgSecCheck 校验一张图片是否含有违法违规内容。
// POST https://api.weixin.qq.com/wxa/img_sec_check?access_token=ACCESS_TOKEN
func (t *wxa) ImgSecCheck() {

}

//      --------------------------------------------------------------------

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容。
// POST https://api.weixin.qq.com/wxa/media_check_async?access_token=ACCESS_TOKEN
func (t *wxa) MediaCheckAsync() {

}

//      --------------------------------------------------------------------

// MsgSecCheck 检查一段文本是否含有违法违规内容。
// POST https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN
func (t *wxa) MsgSecCheck() {

}

//      --------------------------------------------------------------------

// Plugin 向插件开发者发起使用插件的申请|查询已添加的插件|删除已添加的插件
// POST https://api.weixin.qq.com/wxa/plugin?access_token=TOKEN
func (t *wxa) Plugin() {

}

//      --------------------------------------------------------------------

// ServiceMarket 调用服务平台提供的服务
// POST https://api.weixin.qq.com/wxa/servicemarket?access_token=ACCESS_TOKEN
func (t *wxa) ServiceMarket() {

}

//      --------------------------------------------------------------------

// SetNearbyPoiShowStatus 展示/取消展示附近小程序
// POST https://api.weixin.qq.com/wxa/setnearbypoishowstatus?access_token=ACCESS_TOKEN
func (t *wxa) SetNearbyPoiShowStatus() {

}

//      --------------------------------------------------------------------

// SiteSearch 小程序内部搜索API提供针对页面的查询能力，小程序开发者输入搜索词后，将返回自身小程序和搜索词相关的页面。因此，利用该接口，开发者可以查看指定内容的页面被微信平台的收录情况；同时，该接口也可供开发者在小程序内应用，给小程序用户提供搜索能力。
// POST https://api.weixin.qq.com/wxa/sitesearch?access_token=TOKEN
func (t *wxa) SiteSearch() {

}
