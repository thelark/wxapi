package wxapi

import "reflect"

type cgiBinMessageWxOpen struct {
	AccessToken string
}

func (t *cgiBinMessageWxOpen) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------

func (t *cgiBinMessageWxOpen) ActivityID(opts ...option) *cgiBinMessageWxOpenActivityID {
	self := &cgiBinMessageWxOpenActivityID{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessageWxOpen) Template(opts ...option) *cgiBinMessageWxOpenTemplate {
	self := &cgiBinMessageWxOpenTemplate{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *cgiBinMessageWxOpen) UpDataBleMsg(opts ...option) *cgiBinMessageWxOpenUpDataBleMsg {
	self := &cgiBinMessageWxOpenUpDataBleMsg{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}