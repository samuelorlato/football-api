package dtos

type BroadcastRequest struct {
	Type    string  `json:"tipo" validate:"required"`
	Team    string  `json:"time" validate:"required"`
	Score   *string `json:"placar,omitempty"`
	Message string  `json:"mensagem" validate:"required"`
}
