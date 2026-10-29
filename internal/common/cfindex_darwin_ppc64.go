// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc64

package common

// CFBase.h: CFIndex is signed long; CFTypeID is unsigned long.
// CFNumber.h: CFNumberType is CFIndex, not a fixed-width C enum.
type cfIndex = int64
type cfTypeID = uint64

func toCFIndex(value int64) cfIndex {
	return value
}
