package vk

import (
	vkLib "github.com/go-vk-api/vk"
)

type Facade interface {
	CallMethod(method string, params vkLib.RequestParams, response interface{}) error
}

type Vk struct {
	Client Facade
}

func New(Client Facade) Vk {
	return Vk{
		Client: Client,
	}
}
