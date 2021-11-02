// +build appengine

package loggo

import (
	"io"
)

func checkIfTerminal(w io.Writer) bool {
	return true
}
