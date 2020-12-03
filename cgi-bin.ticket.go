package wxapi

import (
	"github.com/thelark/request"
	"fmt"
	"reflect"
)

type cgiBinTicket struct {
	AccessToken string
}

func (t *cgiBinTicket) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

type cgiBinTicketGetTicket struct {
	*ErrorReturn
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func (t *cgiBinTicket) GetTicket(_type string) (*cgiBinTicketGetTicket, error) {
	rsp := new(cgiBinTicketGetTicket)
	if err := wxRequest.Get(
		fmt.Sprintf("%s/getticket", getBasePath()),
		request.WithParam("type", _type),
		request.WithParam("access_token", t.AccessToken),
		request.WithResponse(&rsp),
	); err != nil {
		return nil, err
	}
	if err := checkError(rsp); err != nil {
		return nil, err
	}
	return rsp, nil
}
