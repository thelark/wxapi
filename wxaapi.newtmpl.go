package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type wxaApiNewTmpl struct {
	AccessToken string
}

func (t *wxaApiNewTmpl) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type wxaApiNewTmplAddTemplate struct {
	*ErrorReturn
	PriTmplId string `json:"priTmplId"` // 添加至帐号下的模板id，发送小程序订阅消息时所需
}

// AddTemplate 组合模板并添加至帐号下的个人模板库
// POST https://api.weixin.qq.com/wxaapi/newtmpl/addtemplate?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) AddTemplate(body struct {
	Tid       string `json:"tid"`                 // 模板标题 id，可通过接口获取，也可登录小程序后台查看获取
	KidList   []int  `json:"kidList"`             // 开发者自行组合好的模板关键词列表，关键词顺序可以自由搭配（例如 [3,5,4] 或 [4,5,3]），最多支持5个，最少2个关键词组合
	SceneDesc string `json:"sceneDesc,omitempty"` // 服务场景描述，15个字以内
}) (*wxaApiNewTmplAddTemplate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaApiNewTmplAddTemplate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/addtemplate", getBasePath()),
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

type wxaApiNewTmplDeleteTemplate struct {
	*ErrorReturn
}

// DeleteTemplate 删除帐号下的个人模板
// POST https://api.weixin.qq.com/wxaapi/newtmpl/deltemplate?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) DeleteTemplate(body struct {
	PriTmplId string `json:"priTmplId"` // 要删除的模板id
}) (*wxaApiNewTmplDeleteTemplate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaApiNewTmplDeleteTemplate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/deltemplate", getBasePath()),
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

type wxaApiNewTmplGetCategory struct {
	*ErrorReturn
	Data []struct {
		Id   int    `json:"id"`   // 类目id，查询公共库模版时需要
		Name string `json:"name"` // 类目的中文名
	} `json:"data"` // 类目列表
}

// GetCategory 获取小程序账号的类目
// GET https://api.weixin.qq.com/wxaapi/newtmpl/getcategory?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) GetCategory() (*wxaApiNewTmplGetCategory, error) {
	rsp := new(wxaApiNewTmplGetCategory)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getcategory", getBasePath()),
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

type wxaApiNewTmplGetPubTemplateKeyWordsById struct {
	*ErrorReturn
	Count int `json:"count"` // 模版标题列表总数
	Data  []struct {
		Kid     int    `json:"kid"`     // 关键词 id，选用模板时需要
		Name    string `json:"name"`    // 关键词内容
		Example string `json:"example"` // 关键词内容对应的示例
		Rule    string `json:"rule"`    // 参数类型
	} `json:"data"` // 关键词列表
}

// GetPubTemplateKeyWordsById 获取模板标题下的关键词列表
// GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatekeywords?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) GetPubTemplateKeyWordsById(body struct {
	Tid string `json:"tid"` // 模板标题 id，可通过接口获取
}) (*wxaApiNewTmplGetPubTemplateKeyWordsById, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaApiNewTmplGetPubTemplateKeyWordsById)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getpubtemplatekeywords", getBasePath()),
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

type wxaApiNewTmplGetPubTemplateTitleList struct {
	*ErrorReturn
	Count int `json:"count"` // 模版标题列表总数
	Data  []struct {
		Tid        int    `json:"tid"`        // 模版标题 id
		Title      string `json:"title"`      // 模版标题
		Type       int    `json:"type"`       // 模版类型，2 为一次性订阅，3 为长期订阅
		CategoryId int    `json:"categoryId"` // 模版所属类目 id
	} `json:"data"` // 模板标题列表
}

// GetPubTemplateTitleList 获取帐号所属类目下的公共模板标题
// GET https://api.weixin.qq.com/wxaapi/newtmpl/getpubtemplatetitles?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) GetPubTemplateTitleList(body struct {
	Ids   string `json:"ids"`   // 类目 id，多个用逗号隔开
	Start int    `json:"start"` // 用于分页，表示从 start 开始。从 0 开始计数。
	Limit int    `json:"limit"` // 用于分页，表示拉取 limit 条记录。最大为 30。
}) (*wxaApiNewTmplGetPubTemplateTitleList, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(wxaApiNewTmplGetPubTemplateTitleList)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getpubtemplatetitles", getBasePath()),
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

type wxaApiNewTmplGetTemplateList struct {
	*ErrorReturn
	Data []struct {
		PriTmplId string `json:"priTmplId"` // 添加至帐号下的模板 id，发送小程序订阅消息时所需
	} `json:"data"` // 个人模板列表
	Title   string `json:"title"`   // 模版标题
	Content string `json:"content"` // 模版内容
	Example string `json:"example"` // 模板内容示例
	Type    int    `json:"type"`    //  模版类型，2 为一次性订阅，3 为长期订阅
}

// GetTemplateList 获取当前帐号下的个人模板列表
// GET https://api.weixin.qq.com/wxaapi/newtmpl/gettemplate?access_token=ACCESS_TOKEN
func (t *wxaApiNewTmpl) GetTemplateList() (*wxaApiNewTmplGetTemplateList, error) {
	rsp := new(wxaApiNewTmplGetTemplateList)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/gettemplate", getBasePath()),
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
