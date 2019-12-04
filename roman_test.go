package tdd_test

import (
	"testing"
)

func TestRoman(t *testing.T) {
	c := roman.Converter{}
	got := c.Convert(12)
	want := "XII"
	if got != want {
		t.Errorf("Required: %s Actual: %s", got, want)
	}
}
