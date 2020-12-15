package wxapi

import (
	"fmt"
	"github.com/thelark/request"
	"reflect"
)

type snsOAuth2Component struct {
	AppID                string // 公众号的 appid
	ComponentAppID       string // 服务开发方的 appid
	ComponentAccessToken string // 服务开发方的 access_token
}

func (t *snsOAuth2Component) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type snsOAuth2ComponentAccessToken struct {
	*ErrorReturn
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// GET https://api.weixin.qq.com/sns/oauth2/component/access_token?appid=APPID&code=CODE&grant_type=authorization_code&component_appid=COMPONENT_APPID&component_access_token=COMPONENT_ACCESS_TOKEN
func (t *snsOAuth2Component) AccessToken(code string) (*snsOAuth2ComponentAccessToken, error) {
	rsp := new(snsOAuth2ComponentAccessToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/access_token", getBasePath()),
		request.WithParam("appid", t.AppID),
		request.WithParam("code", code),
		request.WithParam("grant_type", "authorization_code"),
		request.WithParam("component_appid", t.ComponentAppID),
		request.WithParam("component_access_token", t.ComponentAccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}

type snsOAuth2ComponentRefreshToken struct {
	*ErrorReturn
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// RefreshToken 由于 access_token 拥有较短的有效期，当 access_token 超时后，可以使用 refresh_token 进行刷新，refresh_token 拥有较长的有效期（30 天），当 refresh_token 失效的后，需要用户重新授权。
// GET https://api.weixin.qq.com/sns/oauth2/component/refresh_token?appid=APPID&grant_type=refresh_token&component_appid=COMPONENT_APPID&component_access_token=COMPONENT_ACCESS_TOKEN&refresh_token=REFRESH_TOKEN
func (t *snsOAuth2Component) RefreshToken(refreshToken string) (*snsOAuth2ComponentRefreshToken, error) {
	rsp := new(snsOAuth2ComponentRefreshToken)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/refresh_token", getBasePath()),
		request.WithParam("appid", t.AppID),
		request.WithParam("grant_type", "refresh_token"),
		request.WithParam("component_appid", t.ComponentAppID),
		request.WithParam("component_access_token", t.ComponentAccessToken),
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
