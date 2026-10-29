// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc

package common

import (
	"math"
	"strconv"
	"testing"
	"unsafe"
)

// Widths from CFBase.h/CFNumber.h compiled and executed on PowerBook5,6:
// runs/gopsutil-d1-20261029b/c-probe.out in the port workspace.
var (
	_ [4 - int(unsafe.Sizeof(cfIndex(0)))]byte
	_ [int(unsafe.Sizeof(cfIndex(0))) - 4]byte
	_ [4 - int(unsafe.Sizeof(cfTypeID(0)))]byte
	_ [int(unsafe.Sizeof(cfTypeID(0))) - 4]byte
)

func TestCoreFoundationCFIndexOutOfRange(t *testing.T) {
	// A nil library makes it impossible to call C: the named range panic must
	// happen before function resolution, not after a truncating conversion.
	var c CoreFoundationLib
	wrappers := map[string]func(int64){
		"CFNumberCreate":         func(v int64) { c.CFNumberCreate(0, v, nil) },
		"CFNumberGetValue":       func(v int64) { c.CFNumberGetValue(0, v, nil) },
		"CFDictionaryCreate":     func(v int64) { c.CFDictionaryCreate(0, nil, nil, v, 0, 0) },
		"CFArrayGetValueAtIndex": func(v int64) { c.CFArrayGetValueAtIndex(0, v) },
		"CFStringCreateMutable":  func(v int64) { c.CFStringCreateMutable(0, v) },
		"CFStringGetCString":     func(v int64) { c.CFStringGetCString(0, nil, v, 0) },
	}
	for name, call := range wrappers {
		for _, value := range []int64{math.MinInt32 - 1, math.MaxInt32 + 1} {
			t.Run(name+"/"+strconv.FormatInt(value, 10), func(t *testing.T) {
				defer func() {
					if got := recover(); got != errCFIndexOutOfRange {
						t.Errorf("panic = %v, want %v", got, errCFIndexOutOfRange)
					}
				}()
				call(value)
			})
		}
	}
}
