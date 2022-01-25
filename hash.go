package ablib

import "hash/fnv"

func hash(s string) int {
	h := fnv.New64()
	_, _ = h.Write([]byte(s))
	k := int(h.Sum64())
	if k < 0 {
		return -k
	}
	return k
}
