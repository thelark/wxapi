package wxapi

import (
	"fmt"
	"github.com/thelark/request"
	"reflect"
)

type snsOAuth2 struct {
	AppID     string
	AppSecret string
}

func (t *snsOAuth2) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *snsOAuth2) Component(opts ...option) *snsOAuth2Component {
	self := &snsOAuth2Component{}
	self.AppID = t.AppID
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------

type snsOAuth2AccessToken struct {
	*ErrorReturn
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (t *snsOAuth2) AccessToken(code string) (*snsOAuth2AccessToken, error) {
	rsp := new(snsOAuth2AccessToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/access_token", getBasePath()),
		request.WithParam("code", code),
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

//      --------------------------------------------------------------------

type snsOAuth2RefreshToken struct {
	*ErrorReturn
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

func (t *snsOAuth2) RefreshToken(refreshToken string) (*snsOAuth2RefreshToken, error) {
	rsp := new(snsOAuth2RefreshToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/refresh_token", getBasePath()),
		request.WithParam("grant_type", "client_credential"),
		request.WithParam("appid", t.AppID),
		request.WithParam("refresh_token", refreshToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
