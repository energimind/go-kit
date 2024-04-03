package shortid

import (
	"encoding/binary"
)

// random returns a random number using the provided random number generator.
func random(r func(b []byte) (n int, err error)) uint64 {
	var b [8]byte

	_, err := r(b[:])
	if err != nil {
		panic(err)
	}

	return binary.LittleEndian.Uint64(b[:])
}
