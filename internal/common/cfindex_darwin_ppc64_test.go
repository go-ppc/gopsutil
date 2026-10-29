// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc64

package common

import "unsafe"

// Widths emitted by Leopard gcc -arch ppc64 (not executed):
// runs/gopsutil-d1-20261029b/layout-ppc64.json in the port workspace.
var (
	_ [8 - int(unsafe.Sizeof(cfIndex(0)))]byte
	_ [int(unsafe.Sizeof(cfIndex(0))) - 8]byte
	_ [8 - int(unsafe.Sizeof(cfTypeID(0)))]byte
	_ [int(unsafe.Sizeof(cfTypeID(0))) - 8]byte
)
