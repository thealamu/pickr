package pickr

import (
	"fmt"
	"math/rand"
)

const EventRoll Event = "rollEvent"

// roll implements a single n-sided die roll
func roll(r rand.Rand, n int) (string, error) {
	return fmt.Sprintf("%d", r.Intn(n)+1), nil
}
