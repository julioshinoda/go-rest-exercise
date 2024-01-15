package model

type Request struct {
	Department string   `json:"department"`
	JobItem    string   `json:"job_item"`
	Locations  []string `json:"locations"`
}

type CreateJobRequest struct {
	Item        Item       `json:"item"`
	Priority    string     `json:"priority"`
	Department  Department `json:"department"`
	Role        Role       `json:"role"`
	Location    []Location `json:"location"`
	Action      string     `json:"action"`
	Notes       []Notes    `json:"notes"`
	Attachments []string   `json:"attachments"`
	Assignee    Assignee   `json:"assignee"`
	DueBy       string     `json:"dueBy"`
}
type Item struct {
	Name string `json:"name"`
}

type Role struct {
	ID string `json:"id"`
}
type Location struct {
	ID string `json:"id"`
}
type Notes struct {
	Note string `json:"note"`
}
type Assignee struct {
	EmployeeID string `json:"employeeId"`
	Username   string `json:"username"`
	AutoAssign string `json:"autoAssign"`
}
