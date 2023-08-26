package unpacker

import (
	"errors"
	"io"
)

// readAtMost reads at most n bytes from the stream. A nil, empty, or short
// stream is not an error. The returned slice of bytes may have length < n
// without an error.
func readAtMost(stream io.Reader, n int) ([]byte, error) {
	if stream == nil || n <= 0 {
		return []byte{}, nil
	}

	buf := make([]byte, n)
	nr, err := io.ReadFull(stream, buf)

	// Return the bytes read if there was no error OR if the
	// error was EOF (stream was empty) or UnexpectedEOF (stream
	// had less than n). We ignore those errors because we aren't
	// required to read the full n bytes; so an empty or short
	// stream is not actually an error.
	if err == nil ||
		errors.Is(err, io.EOF) ||
		errors.Is(err, io.ErrUnexpectedEOF) {
		return buf[:nr], nil
	}

	return nil, err
}
