package wxapi

import "reflect"

type cgiBinCommentReply struct {
	AccessToken string
}

func (t *cgiBinCommentReply) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// ReplyAdd 回复评论（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token=ACCESS_TOKEN
func (t *cgiBinCommentReply) ReplyAdd() {

}

//      --------------------------------------------------------------------

// ReplyDelete 删除回复（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token=ACCESS_TOKEN
func (t *cgiBinCommentReply) ReplyDelete() {

}