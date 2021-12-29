package addcomments

import (
	"github.com/cs161-staff/grades"
)

// Make takes in a student ID -> assignment name -> comments map and returns a
// policy that adds those comments to the specified students.
func Make(comments map[int]map[string][]string) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		studentComments, ok := comments[student.SID]
		if !ok {
			return []*grades.Student{student}
		}
		newStudent := student.CloneWithAssignments()
		for assignmentName, assignmentComments := range studentComments {
			newAssignment := student.Assignments[assignmentName].Clone()
			newAssignment.Grade.Comments = append(newAssignment.Grade.Comments, assignmentComments...)
			newStudent.Assignments[assignmentName] = newAssignment
		}
		return []*grades.Student{newStudent}
	}
}
