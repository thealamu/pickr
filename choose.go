package pickr

import "math/rand"

const EventChoose Event = "chooseEvent"

// choose makes a choice from an argument list
func choose(r *rand.Rand, args ...string) (string, error) {
	return args[r.Intn(len(args))], nil
}
