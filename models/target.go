package models

type Target struct {
	TargetId  string `json:"target_id,omitempty" db:"target_id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
	Notes     string `json:"notes,omitempty"`
	Country   string `json:"country"`
}
