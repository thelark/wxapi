package wxapi

import "reflect"

type cgiBinMaterial struct {
	AccessToken string
}

func (t *cgiBinMaterial) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// AddNews 新增永久图文素材
// POST https://api.weixin.qq.com/cgi-bin/material/add_news?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) AddNews() {

}

//      --------------------------------------------------------------------
// UpdateNews 修改永久图文素材
// POST https://api.weixin.qq.com/cgi-bin/material/update_news?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) UpdateNews() {

}

//      --------------------------------------------------------------------
// AddMaterial 新增其他类型永久素材
// POST https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=ACCESS_TOKEN&type=TYPE
func (t *cgiBinMaterial) AddMaterial() {

}

//      --------------------------------------------------------------------
// GetMaterial 获取永久素材
// POST https://api.weixin.qq.com/cgi-bin/material/get_material?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) GetMaterial() {

}

//      --------------------------------------------------------------------
// DelMaterial 删除永久素材
// POST https://api.weixin.qq.com/cgi-bin/material/del_material?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) DelMaterial() {

}

//      --------------------------------------------------------------------
// GetMaterialCount 获取素材总数(永久素材的总数，也会计算公众平台官网素材管理中的素材 2.图片和图文消息素材（包括单图文和多图文）的总数上限为5000，其他素材的总数上限为1000 3.调用该接口需https协议)
// GET https://api.weixin.qq.com/cgi-bin/material/get_materialcount?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) GetMaterialCount() {

}

//      --------------------------------------------------------------------
// BatchGetMaterial 获取素材列表(获取永久素材的列表，也包含公众号在公众平台官网素材管理模块中新建的图文消息、语音、视频等素材 2、临时素材无法通过本接口获取 3、调用该接口需https协议)
// POST https://api.weixin.qq.com/cgi-bin/material/batchget_material?access_token=ACCESS_TOKEN
func (t *cgiBinMaterial) BatchGetMaterial() {

}
