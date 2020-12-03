package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinComment struct {
	AccessToken string
}

func (t *cgiBinComment) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinComment) Reply(opts ...option) *cgiBinCommentReply {
	self := &cgiBinCommentReply{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------
type cgiBinCommentClose struct {
	*ErrorReturn
}

// Close 关闭已群发文章评论（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/close?access_token=ACCESS_TOKEN
func (t *cgiBinComment) Close(body struct {
	MsgDataId uint32 `json:"msg_data_id"`     // 群发返回的msg_data_id
	Index     uint32 `json:"index,omitempty"` // 多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
}) (*cgiBinCommentClose, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinCommentClose)
	if err != nil {
		return nil, err
	}
	if err := wxRequest.Post(
		fmt.Sprintf("%s/close", getBasePath()),
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

//      --------------------------------------------------------------------

// Delete 删除评论（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/delete?access_token=ACCESS_TOKEN
func (t *cgiBinComment) Delete() {

}

//      --------------------------------------------------------------------
type cgiBinCommentList struct {
	*ErrorReturn
	Total   uint32 `json:"total"` //  //总数，非comment的size around
	Comment []struct {
		UserCommentId string `json:"user_comment_id"` // 用户评论id
		OpenId        string `json:"openid"`          // openid
		CreateTime    int64  `json:"create_time"`     // 评论时间
		Content       string `json:"content"`         // 评论内容
		CommentType   int    `json:"comment_type"`    // 是否精选评论，0为即非精选，1为true，即精选
		Reply         struct {
			Content    string `json:"content"`     // 作者回复内容
			CreateTime int    `json:"create_time"` // 作者回复时间
		}
	} `json:"comment"`
}

// List 查看指定文章的评论数据（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/list?access_token=ACCESS_TOKEN
func (t *cgiBinComment) List(body struct {
	MsgDataId uint32 `json:"msg_data_id"`     // 群发返回的msg_data_id
	Index     uint32 `json:"index,omitempty"` // 多图文时，用来指定第几篇图文，从0开始，不带默认返回该msg_data_id的第一篇图文
	Begin     uint32 `json:"begin"`           // 起始位置
	Count     uint32 `json:"count"`           // 获取数目（>=50会被拒绝）
	Type      uint32 `json:"type"`            // type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论
}) (*cgiBinCommentList, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinCommentList)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/list", getBasePath()),
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

//      --------------------------------------------------------------------

// MarkElect 将评论标记精选（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token=ACCESS_TOKEN
func (t *cgiBinComment) MarkElect() {

}

//      --------------------------------------------------------------------

// UnMarkElect 将评论取消精选
// POST https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token=ACCESS_TOKEN
func (t *cgiBinComment) UnMarkElect() {

}

//      --------------------------------------------------------------------
type cgiBinCommentOpen struct {
	*ErrorReturn
}

// Open 打开已群发文章评论（新增接口）
// POST https://api.weixin.qq.com/cgi-bin/comment/open?access_token=ACCESS_TOKEN
func (t *cgiBinComment) Open(body struct {
	MsgDataId uint32 `json:"msg_data_id"`     // 群发返回的msg_data_id
	Index     uint32 `json:"index,omitempty"` // 多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
}) (*cgiBinCommentOpen, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinCommentOpen)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/open", getBasePath()),
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
