package iox

import (
	"context"
	"io"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type counter struct {
	sync.Mutex
	count int
}

func (c *counter) inc() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func (c *counter) get() int {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func TestReaderTee(t *testing.T) {
	word := "hello world"
	input := strings.NewReader(word)
	cnt := &counter{}
	readFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw, err := io.ReadAll(reader)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word, string(raw), "readFunc %d", i)
			cnt.inc()
			return err
		}
	}

	err := ReaderTee(context.Background(), input, readFunc(0), readFunc(1), readFunc(2))
	require.NoError(t, err)
	assert.Equal(t, 3, cnt.get())
}

// It tests that ReaderTee can handle the case when the mainFunc reads only part of the input
func TestReaderTee_ReadHalf(t *testing.T) {
	word := "hello world"
	input := strings.NewReader(word)
	cnt := &counter{}
	readHalfFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw := make([]byte, len(word)/2)
			_, err := reader.Read(raw)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word[:len(word)/2], string(raw), "readFunc %d", i)

			cnt.inc()
			return err
		}
	}
	readFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw, err := io.ReadAll(reader)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word[:len(word)/2], string(raw), "readFunc %d", i)
			cnt.inc()
			return err
		}
	}

	err := ReaderTee(context.Background(), input, readHalfFunc(0), readFunc(1), readFunc(2))
	require.NoError(t, err)
	assert.Equal(t, 3, cnt.get())
}

// It test if ReaderTee can handle stream right
// when mainFunc does not read all data at once
func TestReaderTee_ReadBreak(t *testing.T) {
	word := "hello world"
	input := strings.NewReader(word)
	cnt := &counter{}
	readHalfFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw := make([]byte, len(word)/2)
			_, err := reader.Read(raw)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word[:len(word)/2], string(raw), "readFunc %d", i)
			rest := make([]byte, len(word)-len(word)/2)
			_, err = io.ReadFull(reader, rest)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word[len(word)/2:], string(rest), "readFunc %d", i)

			cnt.inc()
			return err
		}
	}
	readFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw, err := io.ReadAll(reader)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word, string(raw), "readFunc %d", i)
			cnt.inc()
			return err
		}
	}

	err := ReaderTee(context.Background(), input, readHalfFunc(0), readFunc(1), readFunc(2))
	require.NoError(t, err)
	assert.Equal(t, 3, cnt.get())
}

func TestReaderTeeBuffered(t *testing.T) {
	word := "hello world"
	input := strings.NewReader(word)
	cnt := &counter{}
	readFunc := func(i int) func(ctx context.Context, reader io.Reader) error {
		return func(ctx context.Context, reader io.Reader) error {
			raw, err := io.ReadAll(reader)
			require.NoError(t, err, "readFunc %d", i)
			assert.Equal(t, word, string(raw), "readFunc %d", i)
			cnt.inc()
			return err
		}
	}

	err := ReaderTeeBuffered(context.Background(), input, readFunc(0), readFunc(1), readFunc(2))
	require.NoError(t, err)
	assert.Equal(t, 3, cnt.get())
}
