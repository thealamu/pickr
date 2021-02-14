/* Package pickr implements functions for making random choices.
 */
package pickr

import (
	"fmt"
	"io"
	"math/rand"
)

type Event string

type Pickr struct {
	out        io.Writer
	event      Event
	randSource *rand.Rand
}

func (p *Pickr) Do(seed int64, args ...string) error {
	if p.event == "" {
		panic("no event specified, there is nothing to do")
	}

	if p.randSource == nil {
		p.randSource = rand.New(rand.NewSource(0))
	}
	p.randSource.Seed(seed)

	switch p.event {
	case EventToss:
		return p.doToss(&r)

	case EventRoll:
		var n string
		if len(args) == 0 {
			n = "6"
		}
		n = args[0]
		return p.doRoll(n)

	case EventChoose:
		return p.doChoose(args...)

	default:
		return fmt.Errorf("unknown event '%s'", p.event)
	}
}

func (p *Pickr) doToss() error {
	v, err := toss(r)
	fmt.Fprintln(p.out, v)
	return err
}

func (p *Pickr) doRoll() {

}

func (p *Pickr) doChoose() {

}
