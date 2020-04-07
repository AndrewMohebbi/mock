package calculate

import (
	"service/structs"
)

//Takes the body from lrs, which is a lrs/structs/message struct
//returns a service/structs/message struct
//Figures what courses there are and adds to the sturtc
//Increments the completed seciton of the course struct by one for every section completed

//Calculate takes an lrs message and returns a progress report
func Calculate(body []byte) structs.Report {
	report := structs.Report{
		Name: "Sheila",
		Courses: []structs.Course{
			structs.Course{Name: "Matlab", Completed: 3},
			structs.Course{Name: "Parallel Toolbox", Completed: 2},
		},
	}

	return report

}
