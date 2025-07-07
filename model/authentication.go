package model

import (
	"encoding/json"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	Experience int    `json:"exp"`
	Badge      string `json:"badge"`
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

func (u *User) DetermineBadge() {
	// According to the requirements, the "Veteran Sales" badge is awarded at 300 XP.
	if u.Experience >= 300 {
		u.Badge = "Veteran Sales"
	} else if u.Experience < 100 {
		u.Badge = "Newbie Sales"
	} else {
		u.Badge = "Normal Sales"
	}
}
