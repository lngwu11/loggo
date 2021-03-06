// +build !windows,!nacl,!plan9

package syslog

import (
	"fmt"
	"log/syslog"
	"os"

	"github.com/lngwu11/loggo"
)

// SyslogHook to send logs via syslog.
type SyslogHook struct {
	Writer        *syslog.Writer
	SyslogNetwork string
	SyslogRaddr   string
}

// Creates a hook to be added to an instance of logger. This is called with
// `hook, err := NewSyslogHook("udp", "localhost:514", syslog.LOG_DEBUG, "")`
// `if err == nil { log.Hooks.Add(hook) }`
func NewSyslogHook(network, raddr string, priority syslog.Priority, tag string) (*SyslogHook, error) {
	w, err := syslog.Dial(network, raddr, priority, tag)
	return &SyslogHook{w, network, raddr}, err
}

func (hook *SyslogHook) Fire(entry *loggo.Entry) error {
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}

	switch entry.Level {
	case loggo.PanicLevel:
		return hook.Writer.Crit(line)
	case loggo.FatalLevel:
		return hook.Writer.Crit(line)
	case loggo.ErrorLevel:
		return hook.Writer.Err(line)
	case loggo.WarnLevel:
		return hook.Writer.Warning(line)
	case loggo.InfoLevel:
		return hook.Writer.Info(line)
	case loggo.DebugLevel, loggo.TraceLevel:
		return hook.Writer.Debug(line)
	default:
		return nil
	}
}

func (hook *SyslogHook) Levels() []loggo.Level {
	return loggo.AllLevels
}
