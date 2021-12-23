package grades

// Roster is a set of lists of possible students (outcomes).
type Roster map[int][]*Student

// Policy is an operation on a student (outcome) that returns one or more
// students as a result.
type Policy func(student *Student) []*Student

// ApplyPolicy takes each outcome in the roster, applies the policy, and
// concatenates the results into a new Roster, performing this action for each
// key in the roster.
func (roster Roster) ApplyPolicy(policy Policy) *Roster {
	newRoster := make(Roster)
	// TODO Parallelize this, but I also want to see if this runs faster than
	// Python to begin with first.
	for key, outcomes := range roster {
		newRoster[key] = []*Student{}
		for _, outcome := range outcomes {
			newRoster[key] = append(newRoster[key], policy(outcome)...)
		}
	}
	return &newRoster
}
