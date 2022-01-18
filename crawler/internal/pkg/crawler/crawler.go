package crawler

import "github.com/Elderly-AI/observer/crawler/internal/pkg/vk"

type Facade struct {
	Vk vk.Vk
}

func New(Vk vk.Vk) *Facade {
	return &Facade{
		Vk: Vk,
	}
}
