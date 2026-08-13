package schemas

import (
	"strings"
	"testing"
)

func TestOptionalDurationRejectsUnparsableValues(t *testing.T) {
	validate := OptionalDuration().ValidateFunc
	if validate == nil {
		t.Fatal("OptionalDuration has no ValidateFunc")
	}

	valid := []string{"48h", "18h", "15m", "30s", "1h30m", "0s"}
	for _, v := range valid {
		t.Run("valid/"+v, func(t *testing.T) {
			if _, errs := validate(v, "admin_lifetime"); len(errs) != 0 {
				t.Errorf("ExpandDuration accepts %q, so validation should too: %v", v, errs)
			}
		})
	}

	// Each of these silently expanded to a zero duration before validation existed.
	invalid := []string{"48 hours", "48", "abc", "48H", "1 h"}
	for _, v := range invalid {
		t.Run("invalid/"+v, func(t *testing.T) {
			_, errs := validate(v, "admin_lifetime")
			if len(errs) == 0 {
				t.Fatalf("expected %q to be rejected", v)
			}
			if !strings.Contains(errs[0].Error(), "admin_lifetime") {
				t.Errorf("error should name the attribute, got %v", errs[0])
			}
		})
	}
}

// An empty string means the attribute is unset: the expanders map it to nil and
// the caller falls back to its default. Rejecting it would break configurations
// that pass an optional variable straight through.
func TestOptionalDurationAcceptsEmptyAsUnset(t *testing.T) {
	if _, errs := OptionalDuration().ValidateFunc("", "admin_lifetime"); len(errs) != 0 {
		t.Errorf("empty string should be accepted as unset, got %v", errs)
	}
}

func TestOptionalQuantityRejectsUnparsableValues(t *testing.T) {
	validate := OptionalQuantity().ValidateFunc
	if validate == nil {
		t.Fatal("OptionalQuantity has no ValidateFunc")
	}

	for _, v := range []string{"1Gi", "100m", "2", ""} {
		t.Run("valid/"+v, func(t *testing.T) {
			if _, errs := validate(v, "memory_request"); len(errs) != 0 {
				t.Errorf("expected %q to be accepted: %v", v, errs)
			}
		})
	}
	for _, v := range []string{"1 Gi", "abc", "1GGi"} {
		t.Run("invalid/"+v, func(t *testing.T) {
			if _, errs := validate(v, "memory_request"); len(errs) == 0 {
				t.Errorf("expected %q to be rejected", v)
			}
		})
	}
}

// Computed-only attributes must not carry a ValidateFunc; the SDK rejects that
// combination, and only the Optional variants are ever set by a user.
func TestComputedHelpersHaveNoValidateFunc(t *testing.T) {
	if ComputedDuration().ValidateFunc != nil {
		t.Error("ComputedDuration must not set a ValidateFunc")
	}
	if ComputedQuantity().ValidateFunc != nil {
		t.Error("ComputedQuantity must not set a ValidateFunc")
	}
}
