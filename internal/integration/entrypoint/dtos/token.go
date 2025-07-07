package dtos

type Token struct {
	Token string `json:"token"`
}

func NewToken(token string) *Token {
	return &Token{
		Token: token,
	}
}
