package overrides

import (
	"fmt"

	"github.com/cs161-staff/grades"
)

// Make returns takes in a student ID -> assignment name -> override score map
// and returns a policy that overrides any score present for a given student's
// assignment with the new score. A note is also added to indicate the
// override.
func Make(overrides map[int]map[string]float64) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		studentOverrides, ok := overrides[student.SID]
		if !ok {
			return []*grades.Student{student}
		}
		newStudent := student.CloneWithAssignments()
		for assignmentName, newScore := range studentOverrides {
			newAssignment := student.Assignments[assignmentName].Clone()
			newAssignment.Grade.Comments = append(newAssignment.Grade.Comments, fmt.Sprintf("Overridden from %f/%f to %f/%f", newAssignment.Grade.Score, newAssignment.MaxScore, newScore, newAssignment.Grade.Score))
			newAssignment.Grade.Score = newScore
			newStudent.Assignments[assignmentName] = newAssignment
		}
		return []*grades.Student{newStudent}
	}
}
