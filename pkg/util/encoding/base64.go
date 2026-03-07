package encoding

import (
	"encoding/base64"
)

// B64Encode base64-encodes a string and returns the encoded string.
// data: the string to be encoded.
func B64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// B64Decode base64-decodes a string and returns the decoded string or an error.
// data: the base64-encoded string to be decoded.
func B64Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	return string(decoded), err
}
