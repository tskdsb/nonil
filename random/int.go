package random

import (
	"math/rand"
)

var (
	r *rand.Rand
)

func randomInt() int {
	return r.Intn(100)
}
