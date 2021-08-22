package models

type Student struct {
	Name      string `json:"name"`
	Class     string `json:"class"`
	Nim       string `json:"nim"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
