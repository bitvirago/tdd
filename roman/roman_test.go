package roman_test

import (
	"github.com/bitvirago/tdd/roman"
	"testing"
)

func TestRoman(t *testing.T) {
	c := roman.Converter{}
	got := c.Convert(7)
	want := "VII"
	if got != want {
		t.Errorf("Required: %s Actual: %s", got, want)
	}

}
