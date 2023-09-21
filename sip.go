package compare

import (
	"go.k6.io/k6/js/modules"
)

// init is called by the Go runtime at application startup.

func init() {

	modules.Register("k6/x/compare", new(Sip))

}

// Sip is the type for our custom API.

type Sip struct{}

func (c *Sip) IsGreater() bool {
	return false
}
