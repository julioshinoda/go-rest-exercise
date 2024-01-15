package model

type Departments struct {
	List []Department `json:"departments"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
