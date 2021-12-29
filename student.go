package grades

import (
	"fmt"
)

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

// GenerateGradeReport generates a GradeReport based on the student's current
// information.
func (student *Student) GenerateGradeReport() *GradeReport {
	gradeReport := &GradeReport{}

	// Build assignment reports.
	for _, assignment := range student.Assignments {
		rawScore := assignment.Grade.Score / assignment.MaxScore
		comments := make([]string, len(assignment.Grade.Comments))
		for i, comment := range assignment.Grade.Comments {
			comments[i] = comment
		}
		var adjustedScore float64
		if override, present := assignment.Grade.Override(); present {
			adjustedScore = override
		} else {
			adjustedScore = rawScore
			for _, multiplier := range assignment.Grade.MultipliersApplied {
				adjustedScore *= multiplier.Factor
				comments = append(comments, fmt.Sprintf("x%f (%s)", multiplier.Factor, multiplier.Description))
			}
		}
		weightedScore := adjustedScore * assignment.Weight
		gradeReport.Assignments[assignment.Name] = &ReportAssignment{
			Raw:      rawScore,
			Adjusted: adjustedScore,
			Weighted: weightedScore,
			Comments: comments,
		}
	}

	// Build category reports and total score.
	for _, category := range student.Categories {
		// Track total numerator as sum of assignments' adjusted score * weight
		// and denominator as sum of weights.
		categoryNumerator := 0.0
		categoryDenominator := 0.0

		for _, assignment := range student.Assignments {
			if assignment.CategoryName != category.Name {
				continue
			}
			if assignment.Grade.Dropped {
				continue
			}

			assignmentReport := gradeReport.Assignments[assignment.Name]
			categoryNumerator += assignmentReport.Weighted
			categoryDenominator += assignment.Weight
		}

		var rawScore float64
		if categoryDenominator > 0.0 {
			rawScore = categoryNumerator / categoryDenominator
		} else {
			rawScore = 0.0
		}
		comments := make([]string, len(category.Comments))
		for i, comment := range category.Comments {
			comments[i] = comment
		}
		var adjustedScore float64
		if override, present := category.Override(); present {
			adjustedScore = override
		} else {
			adjustedScore = rawScore
		}
		weightedScore := adjustedScore / category.Weight
		gradeReport.Categories[category.Name] = &ReportCategory{
			Raw:      rawScore,
			Adjusted: adjustedScore,
			Weighted: weightedScore,
			Comments: comments,
		}

		// Sum total score.
		gradeReport.TotalScore += weightedScore
	}

	return gradeReport
}
