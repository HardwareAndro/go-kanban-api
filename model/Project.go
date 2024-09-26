package model

type Project struct {
	ID         string     `json:"id,omitempty"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}
