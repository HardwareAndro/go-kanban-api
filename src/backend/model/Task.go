package model

type Task struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	DueDate     string   `json:"dueDate" bson:"due_date"`
	StartDate   string   `json:"startDate" bson:"start_date"`
	EndDate     string   `json:"endDate" bson:"end_date"`
	Priority    string   `json:"priority" bson:"priority"`
	Category    Category `json:"category,omitempty" bson:"category,omitempty"`
	User        User     `json:"user,omitempty" bson:"user,omitempty"`
}
