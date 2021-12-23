package grades

// ReportCategory is the representation of a Category on a Report.
type ReportCategory struct {
	// Raw is the raw score in the cateogry, from 0 to 1.
	Raw float64

	// Adjusted is the true score representing the category, from 0 to 1.
	Adjusted float64

	// Weighted is the score of the category after applying its weight, AKA its
	// contribution to the total score.
	Weighted float64

	// Comments is the human-readable comments on the category.
	Comments []string
}

// ReportAssignment is the representation of an Assignment on a Report.
type ReportAssignment struct {
	// Raw is the raw score on the assignment, from 0 to 1.
	Raw float64

	// Adjusted is the true score representing the assignment, from 0 to 1.
	Adjusted float64

	// Weighted is the score of the category after applying its weight, or its
	// contribution to the category's raw score.
	Weighted float64

	// Comments is the human-readable comments on the assignment.
	Comments []string
}

// GradeReport is a student's final grade report, containing all scores and
// information used in calculations.
type GradeReport struct {
	// Student is the student that the grade report is generated for.
	Student *Student

	// TotalScore is the student's total score in the course, from 0 to 1.
	TotalScore float64

	// Categories is the ReportCategories in the report.
	Categories map[string]*ReportCategory

	// Assignments is the ReqportAssignments in the report.
	Assignments map[string]*ReportAssignment
}
