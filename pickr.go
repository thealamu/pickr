/* Package pickr implements functions for making random choices.
 */
package pickr

import (
	"fmt"
	"io"
	"math/rand"
	"strconv"
)

type Event string

type Pickr struct {
	Out        io.Writer
	Event      Event
	RandSource *rand.Rand
}

// Do performs the event specified in the pickr struct
func (p *Pickr) Do(seed int64, args ...string) error {
	if p.Event == "" {
		panic("no event specified, there is nothing to do")
	}

	if p.RandSource == nil {
		p.RandSource = rand.New(rand.NewSource(0))
	}
	p.RandSource.Seed(seed)

	switch p.Event {
	case EventToss:
		return p.doToss(p.RandSource)

	case EventRoll:
		var n = "6"
		if len(args) > 0 {
			n = args[0]
		}
		return p.doRoll(p.RandSource, n)

	case EventChoose:
		return p.doChoose(p.RandSource, args...)

	default:
		return fmt.Errorf("unknown event '%s'", p.Event)
	}
}

func (p *Pickr) doToss(r *rand.Rand) error {
	v, err := toss(r)
	fmt.Fprintln(p.Out, v)
	return err
}

func (p *Pickr) doRoll(r *rand.Rand, n string) error {
	nAsInt, err := strconv.Atoi(n)
	if err != nil {
		return err
	}
	v, err := roll(r, nAsInt)
	fmt.Fprintln(p.Out, v)
	return err
}

func (p *Pickr) doChoose(r *rand.Rand, args ...string) error {
	v, err := choose(r, args...)
	fmt.Fprintln(p.Out, v)
	return err
}
