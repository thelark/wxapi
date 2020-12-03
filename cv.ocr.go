package wxapi

import (
	"github.com/thelark/request"
	"fmt"
	"reflect"
)

type cvOCR struct {
	AccessToken string
}

func (t *cvOCR) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cvOCRBankcard struct {
	*ErrorReturn
	Number string `json:"number"`
}

// Bankcard 本接口提供基于小程序的银行卡 OCR 识别
// POST https://api.weixin.qq.com/cv/ocr/bankcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN

func (t *cvOCR) Bankcard(fileName string, fileData []byte) (*cvOCRBankcard, error) {
	rsp := new(cvOCRBankcard)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/bankcard", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
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

type cvOCRBizLicense struct {
	*ErrorReturn
	RegNum              string `json:"reg_num"`              // 注册号
	Serial              string `json:"serial"`               // 编号
	LegalRepresentative string `json:"legal_representative"` // 法定代表人姓名
	EnterpriseName      string `json:"enterprise_name"`      // 企业名称
	TypeOfOrganization  string `json:"type_of_organization"` // 组成形式
	Address             string `json:"address"`              // 经营场所/企业住所
	TypeOfEnterprise    string `json:"type_of_enterprise"`   // 公司类型
	BusinessScope       string `json:"business_scope"`       // 经营范围
	RegisteredCapital   string `json:"registered_capital"`   // 注册资本
	PaidInCapital       string `json:"paid_in_capital"`      // 实收资本
	ValidPeriod         string `json:"valid_period"`         // 营业期限
	RegisteredDate      string `json:"registered_date"`      // 注册日期/成立日期
	CertPosition        string `json:"cert_position"`        // 营业执照位置
	ImgSize             string `json:"img_size"`             // 图片大小
}

// BizLicense 本接口提供基于小程序的营业执照 OCR 识别
// POST https://api.weixin.qq.com/cv/ocr/bizlicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvOCR) BizLicense(fileName string, fileData []byte) (*cvOCRBizLicense, error) {
	rsp := new(cvOCRBizLicense)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/bizlicense", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
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

type cvOCRComm struct {
	*ErrorReturn
	Items   string `json:"items"`    // 识别结果
	ImgSize string `json:"img_size"` // 图片大小
}

// Comm 本接口提供基于小程序的通用印刷体 OCR 识别
// POST https://api.weixin.qq.com/cv/ocr/comm?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvOCR) Comm(fileName string, fileData []byte) (*cvOCRComm, error) {
	rsp := new(cvOCRComm)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/comm", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
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

type cvOCRDriving struct {
	*ErrorReturn
	VehicleType                string `json:"vehicle_type"`                 // 车辆类型
	Owner                      string `json:"owner"`                        // 所有人
	Addr                       string `json:"addr"`                         // 住址
	UseCharacter               string `json:"use_character"`                // 使用性质
	Model                      string `json:"model"`                        // 品牌型号
	Vin                        string `json:"vin"`                          // 车辆识别代
	EngineNum                  string `json:"engine_num"`                   // 发动机号码
	RegisterDate               string `json:"register_date"`                // 注册日期
	IssueDate                  string `json:"issue_date"`                   // 发证日期
	PlateNumB                  string `json:"plate_num_b"`                  // 车牌号码
	Record                     string `json:"record"`                       // 号牌
	PassengersNum              string `json:"passengers_num"`               // 核定载人数
	TotalQuality               string `json:"total_quality"`                // 总质量
	TotalPrepareQualityQuality string `json:"totalprepare_quality_quality"` // 整备质量
}

// Driving 本接口提供基于小程序的行驶证 OCR 识别(type:图片识别模式，photo（拍照模式）或 scan（扫描模式）)
// POST https://api.weixin.qq.com/cv/ocr/driving?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvOCR) Driving(fileName string, fileData []byte) (*cvOCRDriving, error) {
	rsp := new(cvOCRDriving)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/driving", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
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

type cvOCRDrivingLicense struct {
	*ErrorReturn
	IdNum        string `json:"id_num"`        // 证号
	Name         string `json:"name"`          // 姓名
	Sex          string `json:"sex"`           // 性别
	Address      string `json:"address"`       // 地址
	BirthDate    string `json:"birth_date"`    // 出生日期
	IssueDate    string `json:"issue_date"`    // 初次领证日期
	CarClass     string `json:"car_class"`     // 准驾车型
	ValidFrom    string `json:"valid_from"`    // 有效期限起始日
	ValidTo      string `json:"valid_to"`      // 有效期限终止日
	OfficialSeal string `json:"official_seal"` // 印章文构
}

// DrivingLicense 本接口提供基于小程序的驾驶证 OCR 识别
// POST https://api.weixin.qq.com/cv/ocr/drivinglicense?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvOCR) DrivingLicense(fileName string, fileData []byte) (*cvOCRDrivingLicense, error) {
	rsp := new(cvOCRDrivingLicense)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/drivinglicense", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
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

type cvOCRIdCard struct {
	*ErrorReturn
	Type        string `json:"type"` // 正面或背面，Front / Back
	Name        string `json:"name"` //
	ID          string `json:"id"`   //
	Addr        string `json:"addr"` //
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"` //
	ValidDate   string `json:"valid_date"`  // 有效期
}

// IdCard 本接口提供基于小程序的身份证 OCR 识别
// POST https://api.weixin.qq.com/cv/ocr/idcard?type=MODE&img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvOCR) IdCard(fileName string, fileData []byte) (*cvOCRIdCard, error) {
	rsp := new(cvOCRIdCard)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/idcard", getBasePath()),
		request.WithParam("access_token", t.AccessToken),
		request.WithFile(fileName, fileData),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
