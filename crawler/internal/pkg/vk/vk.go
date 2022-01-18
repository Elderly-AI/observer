package vk

import (
	vkLib "github.com/go-vk-api/vk"
)

type UserPhotos struct {
	UserID uint64
	Photos []string
}

type User struct {
	UserID uint64   `json:"user_id"`
	Photos []Photos `json:"photos"`
}

type Photos struct {
	Sizes []Size `json:"sizes"`
}

type Size struct {
	Type string `json:"type"`
	Url  string `json:"url"`
}

type Facade interface {
	CallMethod(method string, params vkLib.RequestParams, response interface{}) error
}

type Vk struct {
	Client Facade
}

func New(token string) (Vk, error) {
	client, err := vkLib.NewClientWithOptions(
		vkLib.WithToken(token),
	)
	return Vk{
		Client: client,
	}, err
}
