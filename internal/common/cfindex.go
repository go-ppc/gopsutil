// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"errors"
	"math"
)

var errCFIndexOutOfRange = errors.New("value out of range for CFIndex at darwin/ppc32")

// narrowCFIndex32 keeps the int64 wrapper API without truncating C long arguments.
// The wrappers have no error result, so an unrepresentable argument panics before
// resolving or calling the C function.
func narrowCFIndex32(value int64) int32 {
	if value < math.MinInt32 || value > math.MaxInt32 {
		panic(errCFIndexOutOfRange)
	}
	return int32(value)
}
