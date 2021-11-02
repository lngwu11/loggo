package loggo_test

import (
	"os"

	"github.com/lngwu11/loggo"
)

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []loggo.Level {
	return loggo.AllLevels
}

func (h *DefaultFieldHook) Fire(e *loggo.Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func ExampleDefaultFieldHook() {
	l := loggo.New()
	l.Out = os.Stdout
	l.Formatter = &loggo.TextFormatter{DisableTimestamp: true, DisableColors: true}

	l.AddHook(&DefaultFieldHook{GetValue: func() string { return "with its default value" }})
	l.Info("first log")
	// Output:
	// level=info msg="first log" aDefaultField="with its default value"
}
