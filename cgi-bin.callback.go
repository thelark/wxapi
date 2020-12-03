package wxapi

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/thelark/request"
)

type cgiBinCallBack struct {
	AccessToken string
}

func (t *cgiBinCallBack) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinCallBackCheck struct {
	*ErrorReturn
	DNS []struct {
		IP           string `json:"ip"`
		RealOperator string `json:"real_operator"`
	} `json:"dns"`
	Ping []struct {
		IP           string `json:"ip"`
		FromOperator string `json:"from_operator"`
		PackageLoss  string `json:"package_loss"`
		Time         string `json:"time"`
	} `json:"ping"`
}

// Check 为了帮助开发者排查回调连接失败的问题，提供这个网络检测的API。它可以对开发者URL做域名解析，然后对所有IP进行一次ping操作，得到丢包率和耗时。
// POST https://api.weixin.qq.com/cgi-bin/callback/check?access_token=ACCESS_TOKEN
func (t *cgiBinCallBack) Check(body struct {
	Action        string `json:"action"`
	CheckOperator string `json:"check_operator"`
}) (*cgiBinCallBackCheck, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinCallBackCheck)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/check", getBasePath()),
		request.WithBody(buf),
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
