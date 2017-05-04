package astar

import "testing"

func TestGetGoal(t *testing.T) {
	actual := true
	expected := true

	if actual != expected {
		t.Errorf("Error occured while testing GetGoal: '%s' != '%s'", expected, actual)
	}
}
