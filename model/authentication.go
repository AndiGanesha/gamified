package model

import (
	"encoding/json"
)

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type ResponseSign struct {
	Result Result `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Result struct {
	Token  string `json:"token,omitempty"`
	Others string `json:"others,omitempty"`
}

func (s User) MarshalBinary() ([]byte, error) {
	return json.Marshal(s)
}
