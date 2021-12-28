package slipdays

import (
	"time"

	"github.com/cs161-staff/grades"
)

// Apply applies a slip days policy. Slip days reduce the lateness of an
// assignment by one day. Since slip days can be applied in any particular
// manner and may interact with late policies in arbitrary manners, a
// brute-force search of the slip day application space is returned as a series
// of possibilities.
//
// However, we use the following heuristics:
// - It's better to use more slip days than fewer slip days (TODO not yet
//   implemented).
// - If applying an addition slip day does not help lateness at all (because
//   the assignment isn't late enough), don't apply any slip days.
var Apply grades.Policy = apply

func apply(student *grades.Student) []*grades.Student {
	// Get all slip possibilities for mutually exclusive subsets of slip
	// groups, which should all belong to a distinct category. The slip groups
	// in slipGroupSets[i] have the possibilities in
	// slipGroupSetPossibilities[i].
	slipGroupSets := make([]map[int]time.Duration, 0)
	slipGroupSetPossibilities := make([][]map[int]int, 0)
	for _, category := range student.Categories {
		// Find all late slip groups in the category.
		groupLatenesses := make(map[int]time.Duration)
		for _, assignment := range student.Assignments {
			if assignment.CategoryName != category.Name {
				continue
			}
			if assignment.Grade.Lateness > 0 {
				// The lateness for a slip group is judged by the latest
				// assignment in the group, so use the max lateness value.
				if curLateness, ok := groupLatenesses[assignment.SlipGroup]; !ok || curLateness < assignment.Grade.Lateness {
					groupLatenesses[assignment.SlipGroup] = assignment.Grade.Lateness
				}
			}
		}

		// Get possibilities for these slip groups and append them to the slice.
		possibilities := getSlipPossibilities(groupLatenesses, category.SlipDays)
		slipGroupSets = append(slipGroupSets, groupLatenesses)
		slipGroupSetPossibilities = append(slipGroupSetPossibilities, possibilities)
	}

	// All slip group possibililtes is the cross product of each set of
	// possibilities for each slip group set.
	possibilities := crossProduct(slipGroupSetPossibilities...)
	newStudents := make([]*grades.Student, 0, len(possibilities))
	for _, possibility := range possibilities {
		newStudent := student.CloneWithCategories().CloneWithAssignments()
		for slipGroupSetIndex := range possibility {
			slipGroups := slipGroupSets[slipGroupSetIndex]
			slipGroupSlips := possibility[slipGroupSetIndex]
			for slipGroup := range slipGroups {
				slipDays := slipGroupSlips[slipGroup]
				for _, assignment := range newStudent.Assignments {
					if assignment.SlipGroup == slipGroup {
						newAssignment := assignment.Clone()
						newAssignment.Grade.Lateness -= time.Duration(slipDays * 24)
						newStudent.Assignments[assignment.Name] = newAssignment
					}
				}
			}
		}
		newStudents = append(newStudents, newStudent)
	}

	return newStudents
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

// getSlipPossibilities returns all possible assignments of slip days to slip
// groups. latenesses contains the lateness of each slip group so that not too
// many slip days are applied, and slipDays is the total number of slip days
// that can be assigned.
func getSlipPossibilities(latenesses map[int]time.Duration, slipDays int) []map[int]int {
	// Get a list of groups in an ordered slice.
	groups := make([]int, len(latenesses))
	for group := range latenesses {
		groups = append(groups, group)
	}

	// The helper function finds all possibilities of assigning slips days to
	// all groups from index to the end of the groups parameter.
	var helper func(groups []int, index int, daysLeft int) []map[int]int
	helper = func(groups []int, index int, daysLeft int) []map[int]int {
		if index == len(groups) {
			return []map[int]int{}
		}

		// Apply 0 to the max number of slip days to the cururent group and
		// recurse on the rest of the groups.
		curGroup := groups[index]
		daysLate := durationToDays(latenesses[curGroup])
		var maxSlip int
		if daysLate < daysLeft {
			maxSlip = daysLate
		} else {
			maxSlip = daysLeft
		}
		allPossibilities := make([]map[int]int, 0)
		for slips := 0; slips <= maxSlip; slips++ {
			slipPossibilities := helper(groups, index + 1, daysLeft - slips)
			for i := range slipPossibilities {
				slipPossibilities[i][curGroup] = slips
			}
			allPossibilities = append(allPossibilities, slipPossibilities...)
		}

		return allPossibilities
	}

	return helper(groups, 0, slipDays)
}

// Returns the cross product of the given slices.
func crossProduct(slices... []map[int]int) [][]map[int]int {
	if len(slices) == 0 {
		return [][]map[int]int{}
	}

	// Get length of the cross product so that allocation can be done at once.
	crossLen := 1
	for _, slice := range(slices) {
		crossLen *= len(slice)
	}

	// Get cross product.
	indices := make([]int, len(slices))
	cross := make([][]map[int]int, 0, crossLen)
	for indices[0] < len(slices[0]) {
		next := make([]map[int]int, len(slices))
		for i := range(slices) {
			next[i] = slices[i][indices[i]]
		}
		cross = append(cross, next)

		// Increment indices from right to left.
		indices[len(indices) - 1]++
		for i := len(indices) - 1; i >= 1; i-- {
			if indices[i] >= len(slices[i]) {
				indices[i] = 0
				indices[i - 1]++
			}
		}
	}

	return cross
}
