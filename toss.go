package pickr

import "math/rand"

const EventToss Event = "tossEvent"

// toss implements a single coin toss
func toss(r *rand.Rand) (string, error) {
	switch r.Intn(2) {
	case 0:
		return "Heads", nil
	default:
		return "Tails", nil
	}
}
