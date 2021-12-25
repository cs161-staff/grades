package grades

// Category is a group of assignments that collectively contribute a certain
// weight to the final grade.
type Category struct {
	// Name is the name of the category.
	Name string

	// Weight is the proportion of the final grade that the category
	// contributes to.
	Weight float64

	// Drops is the number of lowest assignment scores that are dropped within
	// the category.
	Drops int

	// SlipDays is the number of additional late days that can be applied
	// across all slip groups, treating the assignment lateness as decreased by
	// that number of days.
	SlipDays int

	// HasLateMultiplier is whether this category uses late multipliers.  If
	// true, late multipliers are applied on a sliding scale. If false, any
	// late assignments are automatically treated as 0.
	HasLateMultiplier bool

	// HasOverride is whether Override is present.
	hasOverride bool

	// Override is the overridden weighted score of this category, if
	// HasOverride is true.
	override float64

	// Comments is the human-readable comments added to this category.
	Comments []string
}

// Override returns the overridden weighted score of the category and whether
// it is present.
func (c *Category) Override() (float64, bool) {
	return c.override, c.hasOverride
}

// SetOverride sets the overridden weighted score of the category.
func (c *Category) SetOverride(newOverride float64) {
	c.hasOverride = true
	c.override = newOverride
}

// ClearOverride clears the overridden weighted score.
func (c *Category) ClearOverride() {
	c.hasOverride = false
	c.override = 0.0
}
