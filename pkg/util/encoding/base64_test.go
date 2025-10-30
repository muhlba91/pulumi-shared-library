package encoding_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/muhlba91/pulumi-shared-library/pkg/util/encoding"
)

func TestB64_RoundtripAndKnownEncodings(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantEnc   string // optional expected encoding; empty means don't check
		wantError bool
	}{
		{"empty", "", "", false},
		{"hello_known", "hello", "aGVsbG8=", false},
		{"utf8", "café", "", false},
		{"special_chars", "with spaces /+ special?", "", false},
		{"multiline", "multiline\nline2\nline3", "", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			enc := encoding.B64Encode(tc.input)
			if tc.wantEnc != "" {
				assert.Equal(t, tc.wantEnc, enc)
			}
			dec, err := encoding.B64Decode(enc)
			if tc.wantError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.input, dec)
			}
		})
	}
}

func TestB64_Decode_InvalidInputs(t *testing.T) {
	invalids := []struct {
		name string
		in   string
	}{
		{"invalid_chars", "not-base64!!"},
		{"missing_padding", "aGVsbG8"}, // missing '='
		{"illegal_char", "aGV*sbG8="},  // '*' illegal
		{"random_short", "abc"},        // too short / invalid
	}

	for _, tc := range invalids {
		t.Run(tc.name, func(t *testing.T) {
			_, err := encoding.B64Decode(tc.in)
			require.Error(t, err)
		})
	}
}
