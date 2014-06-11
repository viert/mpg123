package mpg123

import (
	"io"
	"runtime"
)

var (
	handleCreated = false
)

const (
	OUTPUT_MAX_LENGTH = 32768
)

type PcmWriter struct {
	output           io.Writer
	handle           Handle
	DecodedChunkSize int
}

func (pw *PcmWriter) Write(p []byte) (int, error) {
	pw.DecodedChunkSize = 0

	chunk, retcode, err := Decode(pw.handle, &p, OUTPUT_MAX_LENGTH)

	n, writeErr := pw.output.Write(chunk)
	if writeErr != nil {
		return 0, writeErr
	}

	for retcode != MPG123_NEED_MORE && err == nil {
		chunk, retcode, err = Decode(pw.handle, nil, OUTPUT_MAX_LENGTH)
		m, writeErr = pw.output.Write(chunk)
		if writeErr != nil {
			return 0, writeErr
		}
		n += m
	}

	if err != nil {
		return 0, err
	}

	pw.DecodedChunkSize = n

	return len(p), nil
}

func Cleanup(pw *PcmWriter) {
	if handleCreated {
		Delete(pw.handle)
	}
	if initialized {
		Exit()
	}
}

func NewWriter(out io.Writer) (*PcmWriter, error) {
	pw := new(PcmWriter)
	pw.output = out

	runtime.SetFinalizer(pw, Cleanup)

	err := Init()
	if err != nil {
		return nil, err
	}

	pw.handle, err = New("")
	if err != nil {
		return nil, err
	}
	handleCreated = true

	err = OpenFeed(pw.handle)
	if err != nil {
		return nil, err
	}

	return pw, nil
}
