package wxapi

import "reflect"

type cgiBinExpressDelivery struct {
	AccessToken string
}

func (t *cgiBinExpressDelivery) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinExpressDelivery) Contact(opts ...option) *cgiBinExpressDeliveryContact {
	self := &cgiBinExpressDeliveryContact{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressDelivery) Path(opts ...option) *cgiBinExpressDeliveryPath {
	self := &cgiBinExpressDeliveryPath{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinExpressDelivery) Template(opts ...option) *cgiBinExpressDeliveryTemplate {
	self := &cgiBinExpressDeliveryTemplate{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}