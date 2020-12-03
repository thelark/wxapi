package wxapi

import "reflect"

type cgiBinTagsMembers struct {
	AccessToken string
}

func (t *cgiBinTagsMembers) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------
// BatchTagging 批量为用户打标签
// POST https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=ACCESS_TOKEN
func (t *cgiBinTagsMembers) BatchTagging() {

}

//      --------------------------------------------------------------------
// BatchUnTagging 批量为用户取消标签
// POST https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=ACCESS_TOKEN
func (t *cgiBinTagsMembers) BatchUnTagging() {

}

//      --------------------------------------------------------------------
// GetBlackList 获取公众号的黑名单列表(公众号可通过该接口来获取帐号的黑名单列表，黑名单列表由一串 OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。 该接口每次调用最多可拉取 10000 个OpenID，当列表数较多时，可以通过多次拉取的方式来满足需求。)
// POST https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN
func (t *cgiBinTagsMembers) GetBlackList() {

}

//      --------------------------------------------------------------------
// BatchBlackList 拉黑用户(公众号可通过该接口来拉黑一批用户，黑名单列表由一串 OpenID （加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。)
// POST https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN
func (t *cgiBinTagsMembers) BatchBlackList() {

}

//      --------------------------------------------------------------------
// BatchUnBlackList 取消拉黑用户
// POST https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN
func (t *cgiBinTagsMembers) BatchUnBlackList() {

}
