package test

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/lngwu11/loggo"
	"github.com/stretchr/testify/assert"
)

func TestAllHooks(t *testing.T) {
	assert := assert.New(t)

	logger, hook := NewNullLogger()
	assert.Nil(hook.LastEntry())
	assert.Equal(0, len(hook.Entries))

	logger.Error("Hello error")
	assert.Equal(loggo.ErrorLevel, hook.LastEntry().Level)
	assert.Equal("Hello error", hook.LastEntry().Message)
	assert.Equal(1, len(hook.Entries))

	logger.Warn("Hello warning")
	assert.Equal(loggo.WarnLevel, hook.LastEntry().Level)
	assert.Equal("Hello warning", hook.LastEntry().Message)
	assert.Equal(2, len(hook.Entries))

	hook.Reset()
	assert.Nil(hook.LastEntry())
	assert.Equal(0, len(hook.Entries))

	hook = NewGlobal()

	loggo.Error("Hello error")
	assert.Equal(loggo.ErrorLevel, hook.LastEntry().Level)
	assert.Equal("Hello error", hook.LastEntry().Message)
	assert.Equal(1, len(hook.Entries))
}

func TestLoggingWithHooksRace(t *testing.T) {

	rand.Seed(time.Now().Unix())
	unlocker := rand.Int() % 100

	assert := assert.New(t)
	logger, hook := NewNullLogger()

	var wgOne, wgAll sync.WaitGroup
	wgOne.Add(1)
	wgAll.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			logger.Info("info")
			wgAll.Done()
			if i == unlocker {
				wgOne.Done()
			}
		}(i)
	}

	wgOne.Wait()

	assert.Equal(loggo.InfoLevel, hook.LastEntry().Level)
	assert.Equal("info", hook.LastEntry().Message)

	wgAll.Wait()

	entries := hook.AllEntries()
	assert.Equal(100, len(entries))
}

func TestFatalWithAlternateExit(t *testing.T) {
	assert := assert.New(t)

	logger, hook := NewNullLogger()
	logger.ExitFunc = func(code int) {}

	logger.Fatal("something went very wrong")
	assert.Equal(loggo.FatalLevel, hook.LastEntry().Level)
	assert.Equal("something went very wrong", hook.LastEntry().Message)
	assert.Equal(1, len(hook.Entries))
}
