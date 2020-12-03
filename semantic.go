package wxapi

import "reflect"

type semantic struct {
	AccessToken string
}

func (t *semantic) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 子节点 --------------------------------------------------------------------
func (t *semantic) SemProxy(opts ...option) *semanticSemProxy {
	self := &semanticSemProxy{}
	self.AccessToken = t.AccessToken
	for _, opt := range opts {
		opt(self)
	}
	return self
}