package wxapi

import "reflect"

type customService struct {
	AccessToken string
}

func (t *customService) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *customService) KFAccount(opts ...option) *customServiceKFAccount {
	self := &customServiceKFAccount{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *customService) KFSession(opts ...option) *customServiceKFSession {
	self := &customServiceKFSession{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func (t *customService) MsgRecord(opts ...option) *customServiceMsgRecord {
	self := &customServiceMsgRecord{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}
