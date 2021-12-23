package grades

import (
	"time"
)

// AssignmentSubmission describes a student's graded submission to an
// Assignment.
type AssignmentSubmission struct {
	// Score is the raw score on the submission.
	Score float64

	// Lateness is the amount of time after the deadline that the submission
	// was posted.
	Lateness time.Duration

	// SlipDaysApplied is the number of slip days applied to this submission.
	SlipDaysApplied int

	// MultipliersApplied is the Multipliers applied to this submission.
	MultipliersApplied []Multiplier

	// Dropped is whether the assignment was dropped.
	Dropped bool

	// Override is the overridden raw score for the assignment, if non-nil.
	Override *float64

	// Comments is the human-readable comments on this submission.
	Comments []string
}
