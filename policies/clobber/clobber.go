package clobber

import (
	"errors"
	"math"

	"github.com/cs161-staff/grades"
)

type ClobberStyle int

const (
	StyleScaled = iota
	StyleZScore
)

// Make returns a policy that clobbers from the source assignment name to the
// target assignment name according to the given clobber type. The
// possibiltiies are always either applying the clobber or not applying the
// clobber (the original student).
//
// See the README for more information about clobber styles.
func Make(source string, target string, style ClobberStyle, students []*grades.Student) grades.Policy {
	switch style {
	case StyleScaled:
		return func(student *grades.Student) []*grades.Student {
			newStudent := student.CloneWithAssignments()
			newAssignment := student.Assignments[target].Clone()
			newAssignment.Grade.Score = student.Assignments[source].Grade.Score / student.Assignments[source].MaxScore * student.Assignments[target].MaxScore
			newStudent.Assignments[target] = newAssignment

			return []*grades.Student{student, newStudent}
		}
	case StyleZScore:
		// Generate grade reports for students.
		reports := make(map[int]*grades.GradeReport)
		for id, student := range students {
			reports[id] = student.GenerateGradeReport()
		}

		// Compute source and target mean.
		sourceMean := 0.0
		targetMean := 0.0
		for _, report := range reports {
			// TODO Using the adjusted score makes things weird when using
			// drops and clobbers at the same time.
			sourceMean += report.Assignments[source].Adjusted
			targetMean += report.Assignments[target].Adjusted
		}
		sourceMean /= float64(len(reports))
		targetMean /= float64(len(reports))

		// Compute source and target standard deviation.
		sourceStdev := 0.0
		targetStdev := 0.0
		for _, report := range reports {
			sourceStdev += math.Pow(report.Assignments[source].Adjusted-sourceMean, 2.0)
			targetStdev += math.Pow(report.Assignments[target].Adjusted-targetMean, 2.0)
		}
		sourceStdev /= float64(len(reports) - 1)
		targetStdev /= float64(len(reports) - 1)

		return func(student *grades.Student) []*grades.Student {
			newStudent := student.CloneWithAssignments()
			newAssignment := student.Assignments[target].Clone()
			newAssignment.Grade.Score = (student.Assignments[source].Grade.Score-sourceMean)/sourceStdev*targetStdev + targetMean
			newStudent.Assignments[target] = newAssignment

			return []*grades.Student{student, newStudent}
		}
	default:
		panic(errors.New("Invalid clobber style"))
	}
}
