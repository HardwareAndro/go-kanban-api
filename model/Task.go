package model

type Task struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	DueDate     string   `json:"dueDate"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Priority    string   `json:"priority"`
	Category    Category `json:"category,omitempty"`
	User        User     `json:"user,omitempty"`
}
