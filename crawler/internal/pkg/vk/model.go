package vk

type User struct {
	UserID uint64   `json:"user_id"`
	Photos []string `json:"photos"`
}
