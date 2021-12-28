package grades

// Student is a student whose grade is being calculated.
type Student struct {
	// SID is the student's student ID.
	SID int

	// Name is the student's name.
	Name string

	// Categories is the categories relevant to the student.
	Categories map[string]*Category

	// Assignments is the assignments relevant tot he student.
	Assignments map[string]*Assignment

	// SlipDaysUsed tracks how many slip days the student has used so far.
	SlipDaysUsed int
}

// Clone returns a shallow copy of the student.
func (student *Student) Clone() *Student {
	return &*student
}

// CloneWithCategories returns a shallow copy of the student with a new
// categories map.
func (student *Student) CloneWithCategories() *Student {
	newStudent := student.Clone()
	newStudent.Categories = make(map[string]*Category, len(student.Categories))
	for name, category := range student.Categories {
		newStudent.Categories[name] = category
	}
	return newStudent
}

// CloneWithAssignments returns a shallow copy of the student with a new
// assignments map.
func (student *Student) CloneWithAssignments() *Student {
	newStudent := student.Clone()
	newStudent.Assignments = make(map[string]*Assignment, len(student.Assignments))
	for name, assignment := range student.Assignments {
		newStudent.Assignments[name] = assignment
	}
	return newStudent
}
