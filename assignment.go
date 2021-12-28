package grades

// Assignment is an assignment that students submit work for.
type Assignment struct {
	// Name is the name of the assignment.
	Name string

	// CategoryName is the name of the category belonging to the assignment.
	CategoryName string

	// MaxScore is the maxinum raw score that can be received on this
	// assignment.
	MaxScore float64

	// Weight is the weight that this assignment contributes to the category's
	// score. For example, for 3 assignments of weight {2, 1, 1}, the
	// assignment with weight 2 would contribute 50% to the category.
	Weight float64

	// Slip group is the group of assignemnts that this assignment is a part
	// of. Slip days are applied to a whole group. If -1, no slip days can be
	// applied to this assignment.
	SlipGroup int

	// Grade is the submission present on the assignment.
	Grade AssignmentSubmission
}

// Clone returns a copy of the assignment.
func (a *Assignment) Clone() *Assignment {
	newAsssignment := *a
	newAsssignment.Grade.Comments = make([]string, len(a.Grade.Comments))
	copy(newAsssignment.Grade.Comments, a.Grade.Comments)
	return &newAsssignment
}
