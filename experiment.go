package ablib

import (
	"fmt"
	"sort"
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
func (e Experiment) Treatment(input string) (string, error) {
	// Determine total scale of components
	scale := 0
	for _, comp := range e.Comp {
		if comp.Dist < 0 {
			return "", fmt.Errorf("dist < 0 for %s->%s", e.Name, comp.Name)
		}
		scale += comp.Dist
	}
	// Determine buckets to hash into
	offset := 0
	buckets := make([]int, 0, len(e.Comp))
	for _, comp := range e.Comp {
		offset += comp.Dist
		buckets = append(buckets, offset)
	}
	// Hash into a bucket
	target := hash(input) % scale
	idx := sort.SearchInts(buckets, target)
	if idx < len(buckets) && buckets[idx] == target {
		idx++
	}
	return e.Comp[idx].Name, nil
}

// Returns the experiment name and description (if available).
func (e Experiment) String() string {
	if e.Desc != "" {
		return fmt.Sprintf("%s (%s)", e.Name, e.Desc)
	}
	return e.Name
}
