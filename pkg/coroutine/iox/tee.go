package iox

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/RyoJerryYu/go-utilx/pkg/coroutine/syncx"
)

// ConsumeReaderFunc is a function type that consumes data from an io.Reader.
// The function should handle the reading and processing of the data.
type ConsumeReaderFunc func(ctx context.Context, reader io.Reader) error

// ReaderTee reads data from the input reader and distributes it to multiple consumers.
// It uses io.TeeReader to create a chain of readers, allowing each consumer to read
// the same data stream.
//
// The mainFunc is responsible for driving the reading process - it must read all data
// from the input reader for the tee operation to work properly. All teeFuncs will
// receive the same data that mainFunc reads.
//
// Example usage:
//
//	err := ReaderTee(ctx, inputReader,
//	    func(ctx context.Context, r io.Reader) error {
//	        // Main consumer that must read all data
//	        return processData(r)
//	    },
//	    func(ctx context.Context, r io.Reader) error {
//	        // Additional consumer 1
//	        return saveToFile(r)
//	    },
//	    func(ctx context.Context, r io.Reader) error {
//	        // Additional consumer 2
//	        return uploadToCloud(r)
//	    },
//	)
func ReaderTee(ctx context.Context, in io.Reader, mainFunc ConsumeReaderFunc, teeFuncs ...ConsumeReaderFunc) error {
	var (
		mainErr    error            = nil
		teeErrs    []error          = make([]error, len(teeFuncs))          // all nil
		teeReaders []io.Reader      = make([]io.Reader, len(teeFuncs))      // all nil
		teeWriters []io.WriteCloser = make([]io.WriteCloser, len(teeFuncs)) // all nil
	)

	for i := range teeReaders {
		reader, writer := io.Pipe()
		in = io.TeeReader(in, writer)
		teeReaders[i] = reader
		teeWriters[i] = writer
	}

	wg := syncx.WG(ctx)
	wg.Go(func(ctx context.Context) {
		mainErr = mainFunc(ctx, in)
		for i := range teeWriters {
			teeWriters[i].Close()
		}
	})

	for i := range teeFuncs {
		wg.Go(func(i int) func(ctx context.Context) {
			return func(ctx context.Context) {
				teeErrs[i] = teeFuncs[i](ctx, teeReaders[i])
			}
		}(i))
	}

	wg.Wait()

	errs := make([]error, 0, len(teeErrs)+1) // []error{}
	errs = append(errs, mainErr)
	errs = append(errs, teeErrs...)

	return errors.Join(errs...) // errors.Join automatically filters out nil errors
}

// ReaderTeeBuffered reads all data from the input reader into memory first,
// then provides the buffered data to all consumers. Unlike ReaderTee, this function
// stores all data in memory before processing, which may not be suitable for large
// data streams.
//
// The function creates separate buffer copies for each consumer (including mainFunc),
// ensuring that each consumer gets the complete data regardless of how others
// process it.
//
// Example usage:
//
//	err := ReaderTeeBuffered(ctx, inputReader,
//	    func(ctx context.Context, r io.Reader) error {
//	        // Main consumer
//	        return processBufferedData(r)
//	    },
//	    func(ctx context.Context, r io.Reader) error {
//	        // Additional consumer
//	        return saveBufferedData(r)
//	    },
//	)
func ReaderTeeBuffered(ctx context.Context, in io.Reader, mainFunc ConsumeReaderFunc, teeFuncs ...ConsumeReaderFunc) error {
	var (
		errs []error        = make([]error, len(teeFuncs)+1)        // all nil, main + tee
		bufs []bytes.Buffer = make([]bytes.Buffer, len(teeFuncs)+1) // main buffer + tee buffers
	)

	writers := make([]io.Writer, len(teeFuncs)+1)
	for i := range writers {
		writers[i] = &bufs[i]
	}

	writer := io.MultiWriter(writers...)
	_, err := io.Copy(writer, in)
	if err != nil {
		return err
	}

	funcs := append([]ConsumeReaderFunc{mainFunc}, teeFuncs...)
	wg := syncx.WG(ctx)
	for i := range funcs {
		wg.Go(func(i int) func(ctx context.Context) {
			return func(ctx context.Context) {
				errs[i] = funcs[i](ctx, &bufs[i])
			}
		}(i))
	}

	wg.Wait()

	return errors.Join(errs...) // errors.Join automatically filters out nil errors
}
