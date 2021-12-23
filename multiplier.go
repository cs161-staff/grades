package grades

// Multiplier describes a multiplicative factor applied to an assignment score
// with an associated description for the multiplier.
type Multiplier struct {
	// Factor is the multiplicative factor to be applied on the assignment
	// score.
	Factor float64

	// Description is the human-readable description associated with the
	// multiplier.
	Description string
}
