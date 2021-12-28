package drops

import (
	"reflect"
	"testing"

	"github.com/cs161-staff/grades"
)

func TestCombinations(t *testing.T) {
	elems := []*grades.Assignment{
		{SlipGroup: 1},
		{SlipGroup: 2},
		{SlipGroup: 3},
		{SlipGroup: 4},
		{SlipGroup: 5},
	}
	expected := [][]*grades.Assignment{
		{{SlipGroup: 1}, {SlipGroup: 2}, {SlipGroup: 3}},
		{{SlipGroup: 1}, {SlipGroup: 2}, {SlipGroup: 4}},
		{{SlipGroup: 1}, {SlipGroup: 3}, {SlipGroup: 4}},
		{{SlipGroup: 2}, {SlipGroup: 3}, {SlipGroup: 4}},
		{{SlipGroup: 1}, {SlipGroup: 2}, {SlipGroup: 5}},
		{{SlipGroup: 1}, {SlipGroup: 3}, {SlipGroup: 5}},
		{{SlipGroup: 2}, {SlipGroup: 3}, {SlipGroup: 5}},
		{{SlipGroup: 1}, {SlipGroup: 4}, {SlipGroup: 5}},
		{{SlipGroup: 2}, {SlipGroup: 4}, {SlipGroup: 5}},
		{{SlipGroup: 3}, {SlipGroup: 4}, {SlipGroup: 5}},
	}
	combos := combinations(elems, 3)
	if !reflect.DeepEqual(combos, expected) {
		t.Fail()
	}
}
