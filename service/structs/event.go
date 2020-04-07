package structs

//Event struct logs a single event
type Event struct {
	Course    string `json:"course"`
	Section   int    `json:"section"`
	Completed bool   `json:"completed"`
}
