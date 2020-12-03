package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinUser struct {
	AccessToken string
}

func (t *cgiBinUser) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *cgiBinUser) Tag(opts ...option) *cgiBinUserTag {
	self := &cgiBinUserTag{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------
type cgiBinUserGetData struct {
	OpenID []string `json:"openid"`
}
type cgiBinUserGet struct {
	*ErrorReturn
	Total      int64             `json:"total"`       // 关注该公众账号的总用户数
	Count      int64             `json:"count"`       // 拉取的OPENID个数，最大值为10000
	Data       cgiBinUserGetData `json:"data"`        // 列表数据，OPENID的列表
	NextOpenID string            `json:"next_openid"` // 拉取列表的最后一个用户的OPENID
}

// Get 公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求。
// GET https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func (t *cgiBinUser) Get(nextOpenID string) (*cgiBinUserGet, error) {
	rsp := new(cgiBinUserGet)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/get", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithParam("next_openid", nextOpenID),
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
type cgiBinUserInfo struct {
	*ErrorReturn
	Subscribe      int    `json:"subscribe"`                 // 用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenID         string `json:"openid"`                    // 用户的标识，对当前公众号唯一
	Nickname       string `json:"nickname,omitempty"`        // 用户的昵称
	Sex            int    `json:"sex,omitempty"`             // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string `json:"language,omitempty"`        // 用户的语言，简体中文为zh_CN
	City           string `json:"city,omitempty"`            // 用户所在城市
	Province       string `json:"province,omitempty"`        // 用户所在省份
	Country        string `json:"country,omitempty"`         // 用户所在国家
	HeadImgUrl     string `json:"headimgurl,omitempty"`      // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int64  `json:"subscribe_time,omitempty"`  // 用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionID        string `json:"unionid,omitempty"`         // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark,omitempty"`          // 公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID        int    `json:"groupid,omitempty"`         // 用户所在的分组ID（兼容旧的用户分组接口）
	TagIDList      []int  `json:"tagid_list,omitempty"`      // 用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene,omitempty"` // 返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_OTHERS 其他
	QRScene        int    `json:"qr_scene,omitempty"`        // 二维码扫码场景（开发者自定义）
	QRSceneStr     string `json:"qr_scene_str,omitempty"`    // 二维码扫码场景描述（开发者自定义）
}

// Info 获取用户基本信息（包括UnionID机制）
// GET https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
func (t *cgiBinUser) Info(openid string) (*cgiBinUserInfo, error) {
	rsp := new(cgiBinUserInfo)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/info", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
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

type cgiBinUserInfoBatchGet struct {
	*ErrorReturn
	UserInfoList []cgiBinUserInfo `json:"user_info_list"`
}

// InfoBatchGet 批量获取用户基本信息
// POST https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
func (t *cgiBinUser) InfoBatchGet(body struct {
	UserList []struct {
		OpenID string `json:"openid"`         // 用户的标识，对当前公众号唯一
		Lang   string `json:"lang,omitempty"` // 国家地区语言版本，zh_CN 简体，zh_TW 繁体，en 英语，默认为zh-CN
	} `json:"user_list"`
}) (*cgiBinUserInfoBatchGet, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinUserInfoBatchGet)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/info/batchget", getBasePath()),
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

// InfoUpdateRemark 设置用户备注名
// POST https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
func (t *cgiBinUser) InfoUpdateRemark() {

}
