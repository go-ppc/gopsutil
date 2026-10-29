// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc64

package host

// /var/run/utmpx uses the fixed 32-bit disk record, even for ppc64.
// The public utmpx.h struct is NOT the disk format: its timeval grows to
// 16 bytes at ppc64, whereas the disk keeps an 8-byte timeval32.
// Source: Apple Libc-498.1.7/gen/utmpx-darwin.h and target sys/_structs.h.
// Evidence: runs/gopsutil-d1-20261029b (ppc32 disk executed; ppc64 compiled).
type utmpx32 struct {
	User [256]int8
	Id   [4]int8
	Line [32]int8
	Pid  int32
	Type int16
	_    [2]byte
	Tv   timeval32
	Host [256]int8
	Pad  [16]uint32
}

type timeval32 struct {
	Sec  int32
	Usec int32
}
