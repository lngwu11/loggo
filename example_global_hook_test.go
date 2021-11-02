package loggo_test

import (
	"os"

	"github.com/lngwu11/loggo"
)

var (
	mystring string
)

type GlobalHook struct {
}

func (h *GlobalHook) Levels() []loggo.Level {
	return loggo.AllLevels
}

func (h *GlobalHook) Fire(e *loggo.Entry) error {
	e.Data["mystring"] = mystring
	return nil
}

func ExampleGlobalHook() {
	l := loggo.New()
	l.Out = os.Stdout
	l.Formatter = &loggo.TextFormatter{DisableTimestamp: true, DisableColors: true}
	l.AddHook(&GlobalHook{})
	mystring = "first value"
	l.Info("first log")
	mystring = "another value"
	l.Info("second log")
	// Output:
	// level=info msg="first log" mystring="first value"
	// level=info msg="second log" mystring="another value"
}
