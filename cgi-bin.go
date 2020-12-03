package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBin struct {
	AppID       string
	AppSecret   string
	AccessToken string
}

func (t *cgiBin) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBin) CallBack(opts ...option) *cgiBinCallBack {
	self := &cgiBinCallBack{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Comment(opts ...option) *cgiBinComment {
	self := &cgiBinComment{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) CustomService(opts ...option) *cgiBinCustomService {
	self := &cgiBinCustomService{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Express(opts ...option) *cgiBinExpress {
	self := &cgiBinExpress{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Material(opts ...option) *cgiBinMaterial {
	self := &cgiBinMaterial{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Media(opts ...option) *cgiBinMedia {
	self := &cgiBinMedia{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Menu(opts ...option) *cgiBinMenu {
	self := &cgiBinMenu{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Message(opts ...option) *cgiBinMessage {
	self := &cgiBinMessage{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) SoTer(opts ...option) *cgiBinSoTer {
	self := &cgiBinSoTer{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) QRCode(opts ...option) *cgiBinQRCode {
	self := &cgiBinQRCode{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Tags(opts ...option) *cgiBinTags {
	self := &cgiBinTags{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Template(opts ...option) *cgiBinTemplate {
	self := &cgiBinTemplate{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) Ticket(opts ...option) *cgiBinTicket {
	self := &cgiBinTicket{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) User(opts ...option) *cgiBinUser {
	self := &cgiBinUser{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBin) WXAApp(opts ...option) *cgiBinWXAApp {
	self := &cgiBinWXAApp{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

type cgiBinToken struct {
	*ErrorReturn
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Token 获取小程序全局唯一后台接口调用凭据（access_token）。调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存。
// GET https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=APPID&secret=APPSECRET
func (t *cgiBin) Token() (*cgiBinToken, error) {
	rsp := new(cgiBinToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/token", getBasePath()),
		request.WithParam("grant_type", "client_credential"),
		request.WithParam("appid", t.AppID),
		request.WithParam("secret", t.AppSecret),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	t.set("AccessToken", rsp.AccessToken)
	return rsp, nil
}

//      --------------------------------------------------------------------
type cgiBinGetCallBackIP struct {
	*ErrorReturn
	IPList []string `json:"ip_list"`
}

// GetCallBackIP 获取微信callback IP地址
// GET https://api.weixin.qq.com/cgi-bin/getcallbackip?access_token=ACCESS_TOKEN
func (t *cgiBin) GetCallBackIP() (*cgiBinGetCallBackIP, error) {
	rsp := new(cgiBinGetCallBackIP)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getcallbackip", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
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
type cgiBinGetApiDomainIP struct {
	*ErrorReturn
	IPList []string `json:"ip_list"`
}

// GetApiDomainIP 获取微信API接口 IP地址
// GET https://api.weixin.qq.com/cgi-bin/get_api_domain_ip?access_token=ACCESS_TOKEN
func (t *cgiBin) GetApiDomainIP() (*cgiBinGetApiDomainIP, error) {
	rsp := new(cgiBinGetApiDomainIP)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/get_api_domain_ip", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
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

type cgiBinClearQuota struct {
	*ErrorReturn
}

// ClearQuota 公众号调用或第三方平台帮公众号调用对公众号的所有api调用（包括第三方帮其调用）次数进行清零
// POST https://api.weixin.qq.com/cgi-bin/clear_quota?access_token=ACCESS_TOKEN
func (t *cgiBin) ClearQuota() (*cgiBinClearQuota, error) {
	rsp := new(cgiBinClearQuota)
	body, err := json.Marshal(map[string]string{
		"appid": t.AppID,
	})
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Post(
		fmt.Sprintf("%s/clear_quota", getBasePath()),
		request.WithBody(body),
		request.WithParam("access_token", t.AccessToken),
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

// GetCurrentAutoReplyInfo 获取公众号的自动回复规则
// GET https://api.weixin.qq.com/cgi-bin/get_current_autoreply_info?access_token=ACCESS_TOKEN
func (t *cgiBin) GetCurrentAutoReplyInfo() {

}

//      --------------------------------------------------------------------

// ShortUrl 长链接转短链接接口
// POST https://api.weixin.qq.com/cgi-bin/shorturl?access_token=ACCESS_TOKEN
func (t *cgiBin) ShortUrl() {

}
