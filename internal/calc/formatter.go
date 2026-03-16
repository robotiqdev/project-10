package calc

import (
	"math"
	"strconv"
)

// FormatResult formats a float64 result as a string.
// Whole numbers are formatted without a decimal point.
// Decimals use the minimum number of digits necessary to represent the value uniquely.
func FormatResult(result float64) string {
	if result == math.Trunc(result) && !math.IsInf(result, 0) {
		return strconv.FormatFloat(result, 'f', 0, 64)
	}
	return strconv.FormatFloat(result, 'f', -1, 64)
}
