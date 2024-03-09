package model

type ResponseParseToken struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

type SessionPayload struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}
