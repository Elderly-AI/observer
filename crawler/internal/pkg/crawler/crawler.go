package crawler

import "github.com/go-vk-api/vk"

type VkClientFacade interface {
	CallMethod(method string, params vk.RequestParams, response interface{}) error
}

type Facade struct {
	VkClient VkClientFacade
}

func New(vkClient VkClientFacade) *Facade {
	return &Facade{
		VkClient: vkClient,
	}
}
