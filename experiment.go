package ablib

import (
	"fmt"
)

type Experiment struct {
	Name string
	Desc string
	Comp []Component
}

// Treatment returns the treatment name deterministically
// given an input string.
//
// The input string depends on the business logic of your application.
// For example, if you want to bucket users based on user ID, you can
// pass the user ID as the input string.
func (e Experiment) Treatment(in string) (string, error) {
	// Get sum of all distributions
	n := 0
	for _, cmp := range e.Comp {
		if cmp.Dist < 0 {
			return "", fmt.Errorf("dist < 0 for %s->%s", e.Name, cmp.Name)
		}
		n += cmp.Dist
	}
	// Hash into a bucket
	key := hash(in) % n
	off := 0
	for i, cmp := range e.Comp {
		off += cmp.Dist
		if key < off {
			return e.Comp[i].Name, nil
		}
	}
	panic("Reaching here indicates a bug. Please create an issue at https://github.com/swkeever/ablib/issues.")
}

// Returns the experiment name and description (if available).
func (e Experiment) String() string {
	if e.Desc != "" {
		return fmt.Sprintf("%s (%s)", e.Name, e.Desc)
	}
	return e.Name
}
