package latemultipliers

import (
	"time"

	"github.com/cs161-staff/grades"
)

const MultiplierDesc = "Late multipier"

// Make constructs a late multiplier policy based on a sliding scale based on
// the number of days late. If an assignment is n days late, the late
// multiplier is scale[n - 1]. If an assignment is past len(scale) + 1 days
// late, a x0 multiplier is applied. If an assignment is not late, no
// multiplier is applied.
func Make(scale []float64, grace time.Duration) grades.Policy {
	return func(student *grades.Student) []*grades.Student {
		// Get a map of the lateness of all slip groups. The lateness of a
		// group is the maximum lateness of any assignment in the group.
		groupLatenesses := make(map[int]time.Duration)
		for _, assignment := range student.Assignments {
			if curLateness, ok := groupLatenesses[assignment.SlipGroup]; !ok || curLateness < assignment.Grade.Lateness {
				groupLatenesses[assignment.SlipGroup] = assignment.Grade.Lateness
			}
		}

		// Apply lateness multipliers based on the lateness of the groups.
		newStudent := student.CloneWithAssignments()
		for _, assignment := range student.Assignments {
			category := student.Categories[assignment.CategoryName]

			// Lateness is based on individual assignment if no slip group,
			// else use the slip groups value.
			var lateness time.Duration
			if assignment.SlipGroup == 0 {
				lateness = assignment.Grade.Lateness
			} else {
				lateness = groupLatenesses[assignment.SlipGroup]
			}

			// Subtract grace.
			lateness -= grace

			// Skip if not late.
			if lateness <= 0 {
				continue
			}

			// If the category has a late multiplier, use the scale in the
			// parameter. Else, use an empty slice, which for which all
			// assignments will immediately receive x0 multiplier.
			var curScale []float64
			if category.HasLateMultiplier {
				curScale = scale
			} else {
				curScale = []float64{}
			}

			newAssignment := assignment.Clone()
			daysLate := durationToDays(lateness)

			// Apply late multiplier based on scale.
			var multiplier grades.Multiplier
			if daysLate-1 > len(curScale) {
				// Too late; x0 multipiler.
				multiplier.Factor = 0.0
			} else {
				// Get multiplier from scale.
				multiplier.Factor = curScale[daysLate-1]
			}
			multiplier.Description = MultiplierDesc
			newAssignment.Grade.MultipliersApplied = append(newAssignment.Grade.MultipliersApplied, multiplier)

			newStudent.Assignments[assignment.Name] = newAssignment
		}

		return []*grades.Student{newStudent}
	}
}

// durationToDays rounds the given duration up to the nearest integer number of
// days.
func durationToDays(duration time.Duration) int {
	rounded := duration.Truncate(time.Hour * 24)
	if rounded < duration {
		rounded += time.Hour * 24
	}
	return int(rounded.Hours()) / 24
}
