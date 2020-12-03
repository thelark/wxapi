package wxapi

import (
	"github.com/thelark/request"
	"fmt"
	"reflect"
)

type cvImg struct {
	AccessToken string
}

func (t *cvImg) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 通用变量 ----------------------------------------------------------------
type cvImgSize struct {
	W int64 `json:"w"`
	H int64 `json:"h"`
}
type cvImgPoint struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

// 方法 --------------------------------------------------------------------

type cvImgAiCrop struct {
	*ErrorReturn
	Results []struct {
		CropLeft   int `json:"crop_left"` //
		CropTop    int `json:"crop_top"`
		CropRight  int `json:"crop_right"`
		CropBottom int `json:"crop_bottom"`
	} `json:"results"` // 智能裁剪结果
	ImgSize cvImgSize `json:"img_size"` // 图片大小
}

// AiCrop 本接口提供基于小程序的图片智能裁剪能力。
// POST https://api.weixin.qq.com/cv/img/aicrop?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvImg) AiCrop(fileName string, fileData []byte) (*cvImgAiCrop, error) {
	rsp := new(cvImgAiCrop)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/aicrop", getBasePath()),
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

type cvImgQRCode struct {
	*ErrorReturn
	CodeResults []struct {
		TypeName string `json:"type_name"`
		Data     string `json:"data"`
		Pos      *struct {
			LeftTop     cvImgPoint `json:"left_top"`
			RightTop    cvImgPoint `json:"right_top"`
			RightBottom cvImgPoint `json:"right_bottom"`
			LeftBottom  cvImgPoint `json:"left_bottom"`
		} `json:"pos,omitempty"`
	} `json:"code_results"`
	ImgSize cvImgSize `json:"img_size"` // 图片大小
}

// QRCode 本接口提供基于小程序的条码/二维码识别的API。
// POST https://api.weixin.qq.com/cv/img/qrcode?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvImg) QRCode(fileName string, fileData []byte) (*cvImgQRCode, error) {
	rsp := new(cvImgQRCode)
	if err := wxRequest.Post(
		fmt.Sprintf("%s/qrcode", getBasePath()),
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

type cvImgSuperResolution struct {
	*ErrorReturn
}

// SuperResolution 本接口提供基于小程序的图片高清化能力。
// POST https://api.weixin.qq.com/cv/img/superresolution?img_url=ENCODE_URL&access_token=ACCESS_TOCKEN
func (t *cvImg) SuperResolution() {

}
