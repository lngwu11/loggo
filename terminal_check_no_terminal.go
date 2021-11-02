// +build js nacl plan9

package loggo

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return false
}
