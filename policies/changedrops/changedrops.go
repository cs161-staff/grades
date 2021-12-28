package changedrops

import (
	"github.com/cs161-staff/grades"
)

// Make takes in a student ID -> category name -> drop adjust map and returns a
// policy that adjusts the specified categories for the specified students,
// adding the adjustment to the number of drops, returning it as the only new
// outcome for the student.
func Make(dropsAdjust map[int]map[string]int) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		changes, ok := dropsAdjust[student.SID]
		if !ok {
			return []*grades.Student{student}
		}
		newStudent := student.CloneWithCategories()
		for categoryName, change := range changes {
			newCategory := newStudent.Categories[categoryName].Clone()
			newCategory.Drops += change
			newStudent.Categories[categoryName] = newCategory
		}
		return []*grades.Student{newStudent}
	}
}
