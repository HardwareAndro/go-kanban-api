package model

type Category struct {
	Name  string `json:"name,omitempty"`
	ID    string `json:"id,omitempty"`
	Tasks []Task `json:"tasks,omitempty"`
}
