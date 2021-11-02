package loggo_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/lngwu11/loggo"
	"github.com/stretchr/testify/require"
)

func TestLevelJsonEncoding(t *testing.T) {
	type X struct {
		Level loggo.Level
	}

	var x X
	x.Level = loggo.WarnLevel
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	require.NoError(t, enc.Encode(x))
	dec := json.NewDecoder(&buf)
	var y X
	require.NoError(t, dec.Decode(&y))
}

func TestLevelUnmarshalText(t *testing.T) {
	var u loggo.Level
	for _, level := range loggo.AllLevels {
		t.Run(level.String(), func(t *testing.T) {
			require.NoError(t, u.UnmarshalText([]byte(level.String())))
			require.Equal(t, level, u)
		})
	}
	t.Run("invalid", func(t *testing.T) {
		require.Error(t, u.UnmarshalText([]byte("invalid")))
	})
}

func TestLevelMarshalText(t *testing.T) {
	levelStrings := []string{
		"panic",
		"fatal",
		"error",
		"warning",
		"info",
		"debug",
		"trace",
	}
	for idx, val := range loggo.AllLevels {
		level := val
		t.Run(level.String(), func(t *testing.T) {
			var cmp loggo.Level
			b, err := level.MarshalText()
			require.NoError(t, err)
			require.Equal(t, levelStrings[idx], string(b))
			err = cmp.UnmarshalText(b)
			require.NoError(t, err)
			require.Equal(t, level, cmp)
		})
	}
}