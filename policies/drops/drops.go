package drops

import (
	"github.com/cs161-staff/grades"
)

// Apply applies a drop policy by returning all possible combinations of
// dropping assignments as possibilities, based on the number of drops in each
// category.
var Apply grades.Policy = apply

func apply(student *grades.Student) []*grades.Student {
	// Get combinations of assignments in each category.
	categoryCombos := make([][][]*grades.Assignment, len(student.Categories))
	for _, category := range student.Categories {
		assignmentsInCategory := make([]*grades.Assignment, 0)
		for _, assignment := range student.Assignments {
			if assignment.CategoryName == category.Name {
				assignmentsInCategory = append(assignmentsInCategory, assignment)
			}
		}
		categoryCombos = append(categoryCombos, combinations(assignmentsInCategory, category.Drops))
	}

	// Get cross product of all category combos.
	combos := crossProduct(categoryCombos...)
	newStudents := make([]*grades.Student, len(combos))
	for i, combo := range combos {
		newStudent := student.CloneWithAssignments()
		for _, categoryInCombo := range combo {
			for _, assignmentInCombo := range categoryInCombo {
				newAssignment := assignmentInCombo.Clone()
				newAssignment.Grade.Dropped = true
				newStudent.Assignments[assignmentInCombo.Name] = newAssignment
			}
		}
		newStudents[i] = newStudent
	}

	return newStudents
}

// combinations returns all ways of choosing n elements from elems.
func combinations(elems []*grades.Assignment, n int) [][]*grades.Assignment {
	if len(elems) < n {
		return [][]*grades.Assignment{}
	}
	if n == 0 {
		return [][]*grades.Assignment{{}}
	}

	withoutLastHead := combinations(elems[:len(elems)-1], n)
	withLastHead := combinations(elems[:len(elems)-1], n-1)
	ret := make([][]*grades.Assignment, 0, len(withoutLastHead)+len(withLastHead))
	for _, comboWithoutLast := range withoutLastHead {
		ret = append(ret, comboWithoutLast)
	}
	for _, comboWithLast := range withLastHead {
		ret = append(ret, append(comboWithLast, elems[len(elems)-1]))
	}

	return ret
}

// Returns the cross product of the given slices.
func crossProduct(slices ...[][]*grades.Assignment) [][][]*grades.Assignment {
	if len(slices) == 0 {
		return [][][]*grades.Assignment{}
	}

	// Get length of the cross product so that allocation can be done at once.
	crossLen := 1
	for _, slice := range slices {
		crossLen *= len(slice)
	}

	// Get cross product.
	indices := make([]int, len(slices))
	cross := make([][][]*grades.Assignment, 0, crossLen)
	for indices[0] < len(slices[0]) {
		next := make([][]*grades.Assignment, len(slices))
		for i := range slices {
			next[i] = slices[i][indices[i]]
		}
		cross = append(cross, next)

		// Increment indices from right to left.
		indices[len(indices)-1]++
		for i := len(indices) - 1; i >= 1; i-- {
			if indices[i] >= len(slices[i]) {
				indices[i] = 0
				indices[i-1]++
			}
		}
	}

	return cross
}
