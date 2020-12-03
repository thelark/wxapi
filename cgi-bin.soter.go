package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinSoTer struct {
	AccessToken string
}

func (t *cgiBinSoTer) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinSoTerVerifySignature struct {
	*ErrorReturn
	IsOk bool `json:"is_ok"` // 验证结果
}

// VerifySignature SOTER 生物认证秘钥签名验证
// POST https://api.weixin.qq.com/cgi-bin/soter/verify_signature?access_token=ACCESS_TOKEN
func (t *cgiBinSoTer) VerifySignature(body struct {
	OpenID        string `json:"openid"`         // 用户 openid
	JsonString    string `json:"json_string"`    // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSON 字段
	JsonSignature string `json:"json_signature"` // 通过 wx.startSoterAuthentication 成功回调获得的 resultJSONSignature 字段
}) (*cgiBinSoTerVerifySignature, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinSoTerVerifySignature)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/verify_signature", getBasePath()),
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
