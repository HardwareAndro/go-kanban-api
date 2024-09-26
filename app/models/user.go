package model

type User struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Surname  string `json:"surname,omitempty" bson:"surname,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
}
