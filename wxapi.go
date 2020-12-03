package wxapi

func CGIBin(opts ...option) *cgiBin {
	self := &cgiBin{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func CustomService(opts ...option) *customService {
	self := &customService{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func CV(opts ...option) *cv {
	self := &cv{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func DataCube(opts ...option) *dataCube {
	self := &dataCube{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func Semantic(opts ...option) *semantic {
	self := &semantic{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func SNS(opts ...option) *sns {
	self := &sns{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func WXA(opts ...option) *wxa {
	self := &wxa{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
func WXAApi(opts ...option) *wxaApi {
	self := &wxaApi{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
