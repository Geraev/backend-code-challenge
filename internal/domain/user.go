package domain

type User struct {
	ID      uint64   `json:"user_id"`
	Friends []uint64 `json:"friends"`
}
