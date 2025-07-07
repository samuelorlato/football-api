package dtos

type LoginRequest struct {
	User     string `json:"usuario"`
	Password string `json:"senha"`
}
