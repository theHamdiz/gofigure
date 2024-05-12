// Package gofigure provides a function to format large numbers into abbreviated string representations using suffixes like K, M, B, T, Q, and Qi.
package gofigure

import "fmt"

// formatThresholds and suffixes are aligned in ascending order of magnitude.
var (
	formatThresholds = []int64{1000, 1000000, 1000000000, 1000000000000, 1000000000000000, 1000000000000000000}
	suffixes         = []string{"K", "M", "B", "T", "Q", "Qi"}
)

// FormatFigure converts an int64 to a human-readable string using abbreviations
// such as K (thousand), M (million), B (billion), T (trillion),
// Q (quadrillion), and Qi (quintillion).
func FormatFigure(num int64) string {
	if num < 1000 {
		return fmt.Sprintf("%d", num) // Directly return numbers less than 1000.
	}

	var unit string
	var value float64
	for i, threshold := range formatThresholds {
		if num < threshold {
			continue
		}
		unit = suffixes[i]
		value = float64(num) / float64(threshold)
	}

	return fmt.Sprintf("%.1f%s", value, unit)
}
