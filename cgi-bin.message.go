package wxapi

import "reflect"

type cgiBinMessage struct {
	AccessToken string
}

func (t *cgiBinMessage) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinMessage) Custom(opts ...option) *cgiBinMessageCustom {
	self := &cgiBinMessageCustom{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessage) Mass(opts ...option) *cgiBinMessageMass {
	self := &cgiBinMessageMass{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessage) Subscribe(opts ...option) *cgiBinMessageSubscribe {
	self := &cgiBinMessageSubscribe{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessage) Template(opts ...option) *cgiBinMessageTemplate {
	self := &cgiBinMessageTemplate{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessage) WxOpen(opts ...option) *cgiBinMessageWxOpen {
	self := &cgiBinMessageWxOpen{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}