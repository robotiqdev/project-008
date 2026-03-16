package numparse_test

import (
	"math"
	"testing"

	"github.com/robotiqdev/project-008/internal/numparse"
)

const epsilon = 1e-9

func TestParseNumber(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		// Valid integer strings
		{name: "zero", input: "0", want: 0, wantErr: false},
		{name: "positive integer", input: "42", want: 42, wantErr: false},
		{name: "negative integer", input: "-7", want: -7, wantErr: false},
		{name: "large integer", input: "100", want: 100, wantErr: false},

		// Valid decimal strings
		{name: "pi decimal", input: "3.14", want: 3.14, wantErr: false},
		{name: "negative decimal", input: "-0.5", want: -0.5, wantErr: false},
		{name: "one point zero", input: "1.0", want: 1.0, wantErr: false},

		// Invalid / non-numeric strings
		{name: "alphabetic string", input: "abc", want: 0, wantErr: true},
		{name: "empty string", input: "", want: 0, wantErr: true},
		{name: "word number", input: "one", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := numparse.ParseNumber(tt.input)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseNumber(%q) expected an error, got nil (value=%v)", tt.input, got)
				}
				if got != 0 {
					t.Errorf("ParseNumber(%q) on error expected return value 0, got %v", tt.input, got)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseNumber(%q) unexpected error: %v", tt.input, err)
				return
			}

			if math.Abs(got-tt.want) > epsilon {
				t.Errorf("ParseNumber(%q) = %v, want %v (diff=%v)", tt.input, got, tt.want, math.Abs(got-tt.want))
			}
		})
	}
}
