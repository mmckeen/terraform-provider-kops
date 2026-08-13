package schemas

import (
	"testing"
	"time"
)

// Duration attributes are exposed to Terraform as strings, so the string form is
// the actual user-facing contract. The generated tests only cover the nil case.
func TestDurationRoundTrip(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want time.Duration
		out  string
	}{
		{name: "hours", in: "48h", want: 48 * time.Hour, out: "48h0m0s"},
		{name: "kops default", in: "18h", want: 18 * time.Hour, out: "18h0m0s"},
		{name: "minutes", in: "15m", want: 15 * time.Minute, out: "15m0s"},
		{name: "seconds", in: "30s", want: 30 * time.Second, out: "30s"},
		{name: "compound", in: "1h30m", want: 90 * time.Minute, out: "1h30m0s"},
		{name: "zero", in: "0s", want: 0, out: "0s"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExpandDuration(tt.in)
			if got.Duration != tt.want {
				t.Errorf("ExpandDuration(%q) = %v, want %v", tt.in, got.Duration, tt.want)
			}
			if back := FlattenDuration(got); back != tt.out {
				t.Errorf("FlattenDuration round trip = %q, want %q", back, tt.out)
			}
		})
	}
}
