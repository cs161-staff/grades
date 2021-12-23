package changeslipdays

import (
	"github.com/cs161-staff/grades"
)

// Make takes in a student ID -> category name -> slip day adjust map and
// returns a policy that adjusts the specified categories for the specified
// students, adding the adjustment to the number of slip days, returning it as
// the only new outcome for the student.
func Make(slipDaysAdjust map[int]map[string]int) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		changes, ok := slipDaysAdjust[student.SID]
		if !ok {
			return []*grades.Student{student}
		}
		newStudent := &*student
		for categoryName, change := range changes {
			newCategory := &*newStudent.Categories[categoryName]
			newCategory.SlipDays += change
			newStudent.Categories[categoryName] = newCategory
		}
		return []*grades.Student{newStudent}
	}
}
