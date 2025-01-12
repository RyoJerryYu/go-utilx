package iox

import (
	"bytes"
	"context"
	"errors"
	"io"

	"github.com/RyoJerryYu/go-utilx/pkg/coroutine/syncx"
)

type ConsumeReaderFunc func(ctx context.Context, reader io.Reader) error

// ReaderTee read all from in and tee to all teeFuncs and mainFunc
// It will stream the data, low memory usage
// mainFunc should read all data from in
// teeFuncs can read all what mainFunc read
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

// ReaderTeeBuffered read all from in and tee to all teeFuncs
// It will buffer all the read data in memory
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
