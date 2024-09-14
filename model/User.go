package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	ID       string `json:"id,omitempty"`
}
