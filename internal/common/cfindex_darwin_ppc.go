// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc

package common

// CFBase.h: CFIndex is signed long; CFTypeID is unsigned long.
// CFNumber.h: CFNumberType is CFIndex, not a fixed-width C enum.
type cfIndex = int32
type cfTypeID = uint32

func toCFIndex(value int64) cfIndex {
	return narrowCFIndex32(value)
}
