package calculate

import (
	"encoding/json"
	"service/structs"
)

//Calculate takes an lrs message and returns a progress report
func Calculate(body []byte) structs.Report {
	//Populate lrs message struct
	message := structs.Message{}
	json.Unmarshal(body, &message)

	//Populate progress map
	courses := make(map[string]int)
	for _, v := range message.Events {
		name := v.Course
		if _, ok := courses[name]; ok {
			courses[name]++
		} else {
			courses[name] = 1
		}
	}

	//Prepare report
	report := structs.Report{
		Name:    message.Name,
		Courses: courses,
	}
	return report
}
