// file: converter_test.go
package gofigure

import "testing"

func TestFormatFigure(t *testing.T) {
	cases := []struct {
		number   float64
		expected string
	}{
		{123, "123.0"},
		{1234, "1.2K"},
		{1234567, "1.2M"},
		{1234567890, "1.2B"},
		{1234567890123, "1.2T"},
		{1234567890123456, "1.2Q"},
		{1234567890123456789, "1.2Qi"},
	}

	for _, c := range cases {
		got, _ := FormatFigure(c.number)
		if got != c.expected {
			t.Errorf("FormatFigure(%.1f) = %s; want %s", c.number, got, c.expected)
		}
	}
}
