package wxapi

import (
	"github.com/thelark/request"
	"fmt"
	"reflect"
)

type sns struct {
	AppID     string
	AppSecret string
}

func (t *sns) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *sns) OAuth2(opts ...option) *snsOAuth2 {
	self := &snsOAuth2{}
	self.AppID = t.AppID
	self.AppSecret = t.AppSecret
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

type snsUserInfo struct {
	*ErrorReturn
	OpenID     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"unionid"`
}

func (t *sns) UserInfo(accessToken string, openid string) (*snsUserInfo, error) {
	rsp := new(snsUserInfo)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/userinfo", getBasePath()),
		request.WithParam("access_token", accessToken),
		request.WithParam("openid", openid),
		request.WithParam("lang", "zh_CN"),
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

type snsAuth struct {
	*ErrorReturn
}

func (t *sns) Auth(accessToken string, openid string) (*snsAuth, error) {
	rsp := new(snsAuth)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/auth", getBasePath()),
		request.WithParam("access_token", accessToken),
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

type snsJsCode2Session struct {
	*ErrorReturn
	OpenID     string `json:"openid"`      // 用户唯一标识
	SessionKey string `json:"session_key"` // 会话密钥
	UnionID    string `json:"unionid"`     // 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
}

// JsCode2Session 登录凭证校验。通过 wx.login 接口获得临时登录凭证 code 后传到开发者服务器调用此接口完成登录流程。更多使用方法详见 小程序登录。
// GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func (t *sns) JsCode2Session(code string) (*snsJsCode2Session, error) {
	rsp := new(snsJsCode2Session)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/jscode2session", getBasePath()),
		request.WithParam("js_code", code),
		request.WithParam("grant_type", "authorization_code"),
		request.WithParam("appid", t.AppID),
		request.WithParam("secret", t.AppSecret),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
