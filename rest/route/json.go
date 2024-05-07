package route

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

// json is the JSON codec.
//
//nolint:gochecknoglobals // this is by the book
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// decodeJSON decodes the JSON from the reader into the value.
func decodeJSON(r io.Reader, v any) error {
	if err := json.NewDecoder(r).Decode(v); err != nil {
		return NewBadJSONError("invalid JSON: %s", err)
	}

	return nil
}

// encodeJSON encodes the value into the writer.
func encodeJSON(w io.Writer, v any) error {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return NewBadJSONError("failed to encode JSON: %s", err)
	}

	return nil
}
