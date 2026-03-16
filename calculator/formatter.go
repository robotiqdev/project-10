package calculator

import "strconv"

// FormatResult formats the numeric value of r as a string, stripping
// unnecessary trailing zeros (e.g. 6.0 → "6", 3.5 → "3.5").
func FormatResult(r Result) string {
	return strconv.FormatFloat(r.Value, 'f', -1, 64)
}
