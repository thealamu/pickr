/* Package pickr implements functions for making random choices.
 */
package pickr

import "io"

type Event string

type Pickr struct {
	out   io.Writer
	event Event
}
