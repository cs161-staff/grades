package extensions

import (
	"time"

	"github.com/cs161-staff/grades"
)

// Make takes in a student ID -> assignment name -> extension days map and
// returns a policy that subtracts the specified number of days from the
// assignments from the specified students, returning it as the only new
// outcome for the student.
func Make(extensions map[int]map[string]int) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		studentExtensions, ok := extensions[student.SID]
		if !ok {
			return []*grades.Student{student}
		}
		newStudent := &*student
		for assignmentName, extensionDays := range studentExtensions {
			newAssignment := &*newStudent.Assignments[assignmentName]
			newAssignment.Grade.Lateness -= time.Hour * 24 * time.Duration(extensionDays)
			newStudent.Assignments[assignmentName] = newAssignment
		}
		return []*grades.Student{newStudent}
	}
}
