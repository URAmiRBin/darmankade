package model

type Patient struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Token     string `json:"token"`
}
