package structs

//Message struct is the schema for lrs messages
type Message struct {
	Name   string  `json:"name"`
	Events []Event `json:"events"`
}
