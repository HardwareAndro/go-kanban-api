package model

type Project struct {
	ID         string     `json:"id" bson:"_id,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Categories []Category `json:"categories" bson:"categories,omitempty"`
}
