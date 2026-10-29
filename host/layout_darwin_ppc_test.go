// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc

// Code generated from the Leopard C oracle by generate-host-layouts.py; DO NOT EDIT.
// Source/output: runs/gopsutil-d1-20261029b (ppc32 executed; ppc64 compiled only).

package host

import "unsafe"

var (
	_ [628 - int(unsafe.Sizeof(utmpx32{}))]byte
	_ [int(unsafe.Sizeof(utmpx32{})) - 628]byte
	_ [0 - int(unsafe.Offsetof(utmpx32{}.User))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.User)) - 0]byte
	_ [256 - int(unsafe.Sizeof(utmpx32{}.User))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.User)) - 256]byte
	_ [256 - int(unsafe.Offsetof(utmpx32{}.Id))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Id)) - 256]byte
	_ [4 - int(unsafe.Sizeof(utmpx32{}.Id))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Id)) - 4]byte
	_ [260 - int(unsafe.Offsetof(utmpx32{}.Line))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Line)) - 260]byte
	_ [32 - int(unsafe.Sizeof(utmpx32{}.Line))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Line)) - 32]byte
	_ [292 - int(unsafe.Offsetof(utmpx32{}.Pid))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Pid)) - 292]byte
	_ [4 - int(unsafe.Sizeof(utmpx32{}.Pid))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Pid)) - 4]byte
	_ [296 - int(unsafe.Offsetof(utmpx32{}.Type))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Type)) - 296]byte
	_ [2 - int(unsafe.Sizeof(utmpx32{}.Type))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Type)) - 2]byte
	_ [300 - int(unsafe.Offsetof(utmpx32{}.Tv))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Tv)) - 300]byte
	_ [8 - int(unsafe.Sizeof(utmpx32{}.Tv))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Tv)) - 8]byte
	_ [308 - int(unsafe.Offsetof(utmpx32{}.Host))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Host)) - 308]byte
	_ [256 - int(unsafe.Sizeof(utmpx32{}.Host))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Host)) - 256]byte
	_ [564 - int(unsafe.Offsetof(utmpx32{}.Pad))]byte
	_ [int(unsafe.Offsetof(utmpx32{}.Pad)) - 564]byte
	_ [64 - int(unsafe.Sizeof(utmpx32{}.Pad))]byte
	_ [int(unsafe.Sizeof(utmpx32{}.Pad)) - 64]byte
	_ [8 - int(unsafe.Sizeof(timeval32{}))]byte
	_ [int(unsafe.Sizeof(timeval32{})) - 8]byte
	_ [0 - int(unsafe.Offsetof(timeval32{}.Sec))]byte
	_ [int(unsafe.Offsetof(timeval32{}.Sec)) - 0]byte
	_ [4 - int(unsafe.Sizeof(timeval32{}.Sec))]byte
	_ [int(unsafe.Sizeof(timeval32{}.Sec)) - 4]byte
	_ [4 - int(unsafe.Offsetof(timeval32{}.Usec))]byte
	_ [int(unsafe.Offsetof(timeval32{}.Usec)) - 4]byte
	_ [4 - int(unsafe.Sizeof(timeval32{}.Usec))]byte
	_ [int(unsafe.Sizeof(timeval32{}.Usec)) - 4]byte
)
