package ablib

import (
	"fmt"
)

type Experiments map[string]Experiment
type Experiment []Component
type Component struct {
	Name string
	Dist int
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
	for _, comp := range e {
		if comp.Dist < 0 {
			return "", fmt.Errorf("dist < 0 for %s", comp.Name)
		}
		n += comp.Dist
	}
	if n == 0 {
		return "", fmt.Errorf("total treatment allocation is 0")
	}
	// Hash into a bucket
	key := hash(in) % n
	off := 0
	for _, comp := range e {
		off += comp.Dist
		if key < off {
			return comp.Name, nil
		}
	}
	return "", fmt.Errorf("failed to bucket key %d", key)
}
