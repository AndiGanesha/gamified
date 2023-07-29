package model

type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type ResponseSign struct {
	Result string `json:"result"`
	Error  error  `json:"error"`
}
