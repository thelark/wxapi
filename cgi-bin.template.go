package wxapi

import (
	"github.com/thelark/request"
	"encoding/json"
	"fmt"
	"reflect"
)

type cgiBinTemplate struct {
	AccessToken string
}

func (t *cgiBinTemplate) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinTemplateApiSetIndustry struct {
	*ErrorReturn
}

// ApiSetIndustry 设置所属行业(设置行业可在微信公众平台后台完成，每月可修改行业1次，帐号仅可使用所属行业中相关的模板，为方便第三方开发者，提供通过接口调用的方式来修改账号所属行业)
// POST https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token=ACCESS_TOKEN
func (t *cgiBinTemplate) ApiSetIndustry(body struct {
	IndustryID1 string `json:"industry_id1"`
	IndustryID2 string `json:"industry_id2"`
}) (*cgiBinTemplateApiSetIndustry, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinTemplateApiSetIndustry)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/api_set_industry", getBasePath()),
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

type cgiBinTemplateApiAddTemplate struct {
	*ErrorReturn
	TemplateId string `json:"template_id"`
}

// ApiAddTemplate 获得模板ID(从行业模板库选择模板到帐号后台，获得模板ID的过程可在微信公众平台后台完成。为方便第三方开发者，提供通过接口调用的方式来获取模板ID)
// POST https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token=ACCESS_TOKEN
func (t *cgiBinTemplate) ApiAddTemplate(body struct {
	TemplateIDShort string `json:"template_id_short"`
}) (*cgiBinTemplateApiAddTemplate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinTemplateApiAddTemplate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/api_add_template", getBasePath()),
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

type cgiBinTemplateDelPrivateTemplate struct {
	*ErrorReturn
}

// DelPrivateTemplate 删除模板(删除模板可在微信公众平台后台完成，为方便第三方开发者，提供通过接口调用的方式来删除某帐号下的模板)
// POST https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token=ACCESS_TOKEN
func (t *cgiBinTemplate) DelPrivateTemplate(body struct {
	TemplateID string `json:"template_id"`
}) (*cgiBinTemplateDelPrivateTemplate, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	rsp := new(cgiBinTemplateDelPrivateTemplate)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/del_private_template", getBasePath()),
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

type cgiBinTemplateGetAllPrivateTemplate struct {
	*ErrorReturn
	TemplateList []struct {
		TemplateId      string `json:"template_id"`      // 模板ID
		Title           string `json:"title"`            // 模板标题
		PrimaryIndustry string `json:"primary_industry"` // 模板所属行业的一级行业
		DeputyIndustry  string `json:"deputy_industry"`  // 模板所属行业的二级行业
		Content         string `json:"content"`          // 模板内容
		Example         string `json:"example"`          // 模板示例
	} `json:"template_list"`
}

// GetAllPrivateTemplate 获取模板列表(获取已添加至帐号下所有模板列表，可在微信公众平台后台中查看模板列表信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号下所有模板信息)
// GET https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token=ACCESS_TOKEN
func (t *cgiBinTemplate) GetAllPrivateTemplate() (*cgiBinTemplateGetAllPrivateTemplate, error) {
	rsp := new(cgiBinTemplateGetAllPrivateTemplate)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/get_all_private_template", getBasePath()),
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

type cgiBinTemplateGetIndustry struct {
	*ErrorReturn
	PrimaryIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"primary_industry"`
	SecondaryIndustry struct {
		FirstClass  string `json:"first_class"`
		SecondClass string `json:"second_class"`
	} `json:"secondary_industry"`
}

// GetIndustry 获取设置的行业信息(获取帐号设置的行业信息。可登录微信公众平台，在公众号后台中查看行业信息。为方便第三方开发者，提供通过接口调用的方式来获取帐号所设置的行业信息)
// GET https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token=ACCESS_TOKEN
func (t *cgiBinTemplate) GetIndustry() (*cgiBinTemplateGetIndustry, error) {
	rsp := new(cgiBinTemplateGetIndustry)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/get_industry", getBasePath()),
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
