package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	// trace method could also have multiple params, eg. fmt.printf, logf, errorf
	Trace(...interface{})
}

// this is unexported(lowercase) as the user will never be directly interacting with it.
// the user will be interacting with New, Trace method and the Tracer interface.
type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}

// accepting io.Writer means that the user can decide where the output will be written
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
