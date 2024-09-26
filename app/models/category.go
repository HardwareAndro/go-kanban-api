package model

type Category struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name,omitempty" bson:"name"`
	Tasks []Task `json:"tasks,omitempty" bson:"tasks,omitempty"`
}
