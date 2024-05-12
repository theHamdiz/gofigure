// Package gofigure provides a function to format large numbers into abbreviated string representations using suffixes like K, M, B, T, Q, and Qi.
package gofigure

import (
	"fmt"
	"strconv"
)

// formatThresholds and suffixes are aligned in ascending order of magnitude.
var (
	formatThresholds = []float64{1000, 1000000, 1000000000, 1000000000000, 1000000000000000, 1000000000000000000}
	suffixes         = []string{"K", "M", "B", "T", "Q", "Qi"}
)

// FormatFigure takes an interface{} as input and attempts to convert it to a float64 to format the number.
func FormatFigure(num interface{}) (string, error) {
	var value float64

	switch v := num.(type) {
	case int, int32, int64, uint, uint32, uint64:
		value = v.(float64) // This uses Go's ability to convert different integer types to float64
	case float32, float64:
		value = v.(float64) // Same conversion for float32 and float64
	case string:
		var err error
		value, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse string to number: %v", err)
		}
	default:
		return "", fmt.Errorf("unsupported type provided: %T", num)
	}

	if value < 1000 {
		return fmt.Sprintf("%d", int(value)), nil
	}

	for i, threshold := range formatThresholds {
		if value < threshold {
			continue
		}
		unit := suffixes[i]
		return fmt.Sprintf("%.1f%s", value/threshold, unit), nil
	}

	return fmt.Sprintf("%.1f", value), nil // for very large numbers exceeding Qi
}
