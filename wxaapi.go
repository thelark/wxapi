package wxapi

import "reflect"

type wxaApi struct {
	AccessToken string
}

func (t *wxaApi) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *wxaApi) Log(opts ...option) *wxaApiLog {
	self := &wxaApiLog{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *wxaApi) NewTmpl(opts ...option) *wxaApiNewTmpl {
	self := &wxaApiNewTmpl{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *wxaApi) UserLog(opts ...option) *wxaApiUserLog {
	self := &wxaApiUserLog{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}