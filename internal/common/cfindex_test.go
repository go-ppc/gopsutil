// SPDX-License-Identifier: BSD-3-Clause

package common

import (
	"math"
	"strconv"
	"testing"
)

func TestNarrowCFIndex32(t *testing.T) {
	for _, value := range []int64{math.MinInt32, -1, 0, 1, math.MaxInt32} {
		if got := narrowCFIndex32(value); int64(got) != value {
			t.Errorf("narrowCFIndex32(%d) = %d", value, got)
		}
	}
}

func TestNarrowCFIndex32OutOfRange(t *testing.T) {
	for _, value := range []int64{math.MinInt32 - 1, math.MaxInt32 + 1, math.MinInt64, math.MaxInt64} {
		t.Run(strconv.FormatInt(value, 10), func(t *testing.T) {
			defer func() {
				if got := recover(); got != errCFIndexOutOfRange {
					t.Errorf("narrowCFIndex32(%d) panic = %v, want %v", value, got, errCFIndexOutOfRange)
				}
			}()
			narrowCFIndex32(value)
		})
	}
}
