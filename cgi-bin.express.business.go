package wxapi

import "reflect"

type cgiBinExpressBusiness struct {
	AccessToken string
}

func (t *cgiBinExpressBusiness) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinExpressBusiness) Account(opts ...option) *cgiBinExpressBusinessAccount {
	self := &cgiBinExpressBusinessAccount{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressBusiness) Delivery(opts ...option) *cgiBinExpressBusinessDelivery {
	self := &cgiBinExpressBusinessDelivery{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressBusiness) Order(opts ...option) *cgiBinExpressBusinessOrder {
	self := &cgiBinExpressBusinessOrder{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressBusiness) Path(opts ...option) *cgiBinExpressBusinessPath {
	self := &cgiBinExpressBusinessPath{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressBusiness) Printer(opts ...option) *cgiBinExpressBusinessPrinter {
	self := &cgiBinExpressBusinessPrinter{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressBusiness) Quota(opts ...option) *cgiBinExpressBusinessQuota {
	self := &cgiBinExpressBusinessQuota{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}

// 方法 --------------------------------------------------------------------
// TestUpdateOrder 模拟快递公司更新订单状态, 该接口只能用户测试
// POST https://api.weixin.qq.com/cgi-bin/express/business/test_update_order?access_token=ACCESS_TOKEN
func (t *cgiBinExpressBusiness) TestUpdateOrder() {

}
