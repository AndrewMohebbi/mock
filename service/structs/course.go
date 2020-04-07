package structs

//Course struct logs a single event
type Course struct {
	Name      string `json:"name"`
	Completed int    `json:"completed"`
}
