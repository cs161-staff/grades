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
