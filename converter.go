// Package gofigure provides a function to format large numbers into abbreviated string representations using suffixes like K, M, B, T, Q, and Qi.
package gofigure

import (
	"fmt"
	"strconv"
)

// formatThresholds and suffixes are aligned in ascending order of magnitude.
var (
	formatThresholds = []float64{
		1000,                                                                                  // Thousand
		1_000_000,                                                                             // Million
		1_000_000_000,                                                                         // Billion
		1_000_000_000_000,                                                                     // Trillion
		1_000_000_000_000_000,                                                                 // Quadrillion
		1_000_000_000_000_000_000,                                                             // Quintillion
		1_000_000_000_000_000_000_000,                                                         // Sextillion
		1_000_000_000_000_000_000_000_000,                                                     // Septillion
		1_000_000_000_000_000_000_000_000_000,                                                 // Octillion
		1_000_000_000_000_000_000_000_000_000_000,                                             // Nonillion
		1_000_000_000_000_000_000_000_000_000_000_000,                                         // Decillion
		1_000_000_000_000_000_000_000_000_000_000_000_000,                                     // Undecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000,                                 // Duodecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000,                             // Tredecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,                         // Quattuordecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,                     // Quindecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,                 // Sexdecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,             // Septendecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,         // Octodecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000,     // Novemdecillion
		1_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000_000, // Vigintillion
	}
	suffixes = []string{
		"K", "M", "B", "T", "Q", "Qi",
		"Sx", "Sp", "Oc", "N", "Dc",
		"Ud", "Dd", "Td", "Qd", "Qn",
		"Sd", "St", "Od", "Nd", "Vg",
	}
)

// FormatFigure takes an interface{} as input and attempts to convert it to a float64 to format the number.
func FormatFigure(num interface{}) (string, error) {
	var value float64

	switch v := num.(type) {
	case int:
		value = float64(v) // Convert int to float64
	case int32:
		value = float64(v) // Convert int32 to float64
	case int64:
		value = float64(v) // Convert int64 to float64
	case uint:
		value = float64(v) // Convert uint to float64
	case uint32:
		value = float64(v) // Convert uint32 to float64
	case uint64:
		value = float64(v) // Convert uint64 to float64
	case float32:
		value = float64(v) // Convert float32 to float64
	case float64:
		value = v // No conversion needed for float64
	case string:
		var err error
		value, err = strconv.ParseFloat(v, 64)
		if err != nil {
			return "", fmt.Errorf("could not parse string to number: %v", err)
		}
	default:
		return "", fmt.Errorf("unsupported type provided: %T", num)
	}

	// Assuming formatThresholds and suffixes are defined elsewhere in your code
	if value <= 1000 {
		return fmt.Sprintf("%.1f", value), nil
	}

	for i, threshold := range formatThresholds {
		// 1M
		// 1.2M > 1K && (1.2M <= 1M)
		if value >= threshold && value <= formatThresholds[i+1] {
			unit := suffixes[i]
			return fmt.Sprintf("%.1f%s", value/threshold, unit), nil
		}
	}

	// Fallback for very large numbers
	return fmt.Sprintf("%.1f", value), nil
}
