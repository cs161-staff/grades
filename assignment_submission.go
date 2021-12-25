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

	// HasOverride is whether Override is present.
	hasOverride bool

	// Override is the overridden raw score for the assignment, if HasOverride
	// is true.
	override float64

	// Comments is the human-readable comments on this submission.
	Comments []string
}

// Override returns the overridden raw score of the assignment and whether it
// is present.
func (s *AssignmentSubmission) Override() (float64, bool) {
	return s.override, s.hasOverride
}

// SetOverride sets the overridden raw score of the assignment.
func (s *AssignmentSubmission) SetOverride(newOverride float64) {
	s.hasOverride = true
	s.override = newOverride
}

// ClearOverride clears the overridden raw score.
func (s *AssignmentSubmission) ClearOverride() {
	s.hasOverride = false
	s.override = 0.0
}
