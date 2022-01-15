package model

type VkUser struct {
	UserID uint64   `json:"user_id"`
	Photos []string `json:"photos"`
}
