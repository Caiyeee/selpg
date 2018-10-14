package main

import (
	"fmt"
	"io"
)

// Log helper type to wrap fmt.PrintXX
type Log struct {
	Progname string
	Stdout   io.Writer
	Stderr   io.Writer
}

// Print wraps fmt.Fprint(stdout, ...)
func (Log *Log) Print(args ...interface{}) (n int, err error) {
	return fmt.Fprint(Log.Stdout, args...)
}

// Printf wraps fmt.Fprintf(stdout, ...)
func (Log *Log) Printf(format string, args ...interface{}) (n int, err error) {
	return fmt.Fprintf(Log.Stdout, format, args...)
}

// Println wraps fmt.Fprintln(stdout, ...)
func (Log *Log) Println(args ...interface{}) (n int, err error) {
	return fmt.Fprintln(Log.Stdout, args...)
}

func (Log *Log) errPrintHeader() (n int, err error) {
	return fmt.Fprintf(Log.Stderr, "%s: ", Log.Progname)
}

// ErrPrint wraps fmt.FPrint(os.Stderr, ...)
func (Log *Log) ErrPrint(args ...interface{}) (n int, err error) {
	n, err = Log.errPrintHeader()
	if err != nil {
		return
	}
	n2, err := fmt.Fprint(Log.Stderr, args...)
	n += n2
	return
}

// ErrPrintf wraps fmt.FPrintf(os.Stderr, ...)
func (Log *Log) ErrPrintf(format string, args ...interface{}) (n int, err error) {
	n, err = Log.errPrintHeader()
	if err != nil {
		return
	}
	n2, err := fmt.Fprintf(Log.Stderr, format, args...)
	n += n2
	return
}

// ErrPrintln wraps fmt.FPrintln(os.Stderr, ...)
func (Log *Log) ErrPrintln(args ...interface{}) (n int, err error) {
	n, err = Log.errPrintHeader()
	if err != nil {
		return
	}
	n2, err := fmt.Fprintln(Log.Stderr, args...)
	n += n2
	return
}
