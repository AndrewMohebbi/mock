package structs

//Report struct is the schema for service messages
type Report struct {
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}
