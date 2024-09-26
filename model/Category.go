package model

type Category struct {
	Name  string `json:"name"`
	ID    string `json:"id,omitempty"`
	Tasks []Task `json:"tasks"`
}
