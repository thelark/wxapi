package wxapi

import "reflect"

type cgiBinTags struct {
	AccessToken string
}

func (t *cgiBinTags) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *cgiBinTags) Members(opts ...option) *cgiBinTagsMembers {
	self := &cgiBinTagsMembers{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------
// Create 创建标签(一个公众号，最多可以创建100个标签。)
// POST https://api.weixin.qq.com/cgi-bin/tags/create?access_token=ACCESS_TOKEN
func (t *cgiBinTags) Create() {

}

//      --------------------------------------------------------------------
// Get 获取公众号已创建的标签
// GET https://api.weixin.qq.com/cgi-bin/tags/get?access_token=ACCESS_TOKEN
func (t *cgiBinTags) Get() {

}

//      --------------------------------------------------------------------
// Update 编辑标签
// POST https://api.weixin.qq.com/cgi-bin/tags/update?access_token=ACCESS_TOKEN
func (t *cgiBinTags) Update() {

}

//      --------------------------------------------------------------------
// Delete 删除标签(请注意，当某个标签下的粉丝超过10w时，后台不可直接删除标签。此时，开发者可以对该标签下的openid列表，先进行取消标签的操作，直到粉丝数不超过10w后，才可直接删除该标签。)
// POST https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=ACCESS_TOKEN
func (t *cgiBinTags) Delete() {

}

//      --------------------------------------------------------------------
// GetIDList 获取用户身上的标签列表
// POST https://api.weixin.qq.com/cgi-bin/tags/getidlist?access_token=ACCESS_TOKEN
func (t *cgiBinTags) GetIDList() {

}
