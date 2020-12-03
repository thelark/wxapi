package wxapi

import "reflect"

type cgiBinExpress struct {
	AccessToken string
}

func (t *cgiBinExpress) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinExpress) Business(opts ...option) *cgiBinExpressBusiness {
	self := &cgiBinExpressBusiness{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpress) Delivery(opts ...option) *cgiBinExpressDelivery {
	self := &cgiBinExpressDelivery{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}