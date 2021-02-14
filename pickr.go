/* Package pickr implements functions for making random choices.
 */
package pickr

import (
	"fmt"
	"io"
)

type Event string

type Pickr struct {
	out   io.Writer
	event Event
}

func (p *Pickr) Do(args ...string) error {
	if p.event == "" {
		panic("no event specified, there is nothing to do")
	}

	switch p.event {
	case EventToss:
		return p.doToss()

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
