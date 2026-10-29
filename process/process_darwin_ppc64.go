// SPDX-License-Identifier: BSD-3-Clause
//go:build darwin && ppc64

package process

// Leopard (xnu-1228.15.4) sys/sysctl.h, sys/proc.h and sys/proc_info.h.
// Layout evidence: runs/gopsutil-d1-20261029b in the port workspace.
// Kernel addresses are integers, not Go pointers: these are copied-out records.

type KinfoProc struct {
	Proc  ExternProc
	Eproc Eproc
}

type ExternProc struct {
	P_un        [16]byte
	P_vmspace   uint64
	P_sigacts   uint64
	P_flag      int32
	P_stat      int8
	Pad_cgo_0   [3]byte
	P_pid       int32
	P_oppid     int32
	P_dupfd     int32
	Pad_cgo_1   [4]byte
	User_stack  uint64
	Exit_thread uint64
	P_debugger  int32
	Sigwait     uint32
	P_estcpu    uint32
	P_cpticks   int32
	P_pctcpu    uint32
	Pad_cgo_2   [4]byte
	P_wchan     uint64
	P_wmesg     uint64
	P_swtime    uint32
	P_slptime   uint32
	P_realtimer Itimerval
	P_rtime     Timeval
	P_uticks    uint64
	P_sticks    uint64
	P_iticks    uint64
	P_traceflag int32
	Pad_cgo_3   [4]byte
	P_tracep    uint64
	P_siglist   int32
	Pad_cgo_4   [4]byte
	P_textvp    uint64
	P_holdcnt   int32
	P_sigmask   uint32
	P_sigignore uint32
	P_sigcatch  uint32
	P_priority  uint8
	P_usrpri    uint8
	P_nice      int8
	P_comm      [17]int8
	Pad_cgo_5   [4]byte
	P_pgrp      uint64
	P_addr      uint64
	P_xstat     uint16
	P_acflag    uint16
	Pad_cgo_6   [4]byte
	P_ru        uint64
}

type Eproc struct {
	Paddr     uint64
	Sess      uint64
	Pcred     Upcred
	Ucred     Uucred
	Pad_cgo_0 [4]byte
	Vm        Vmspace
	Ppid      int32
	Pgid      int32
	Jobc      int16
	Pad_cgo_1 [2]byte
	Tdev      int32
	Tpgid     int32
	Pad_cgo_2 [4]byte
	Tsess     uint64
	Wmesg     [8]int8
	Xsize     int32
	Xrssize   int16
	Xccount   int16
	Xswrss    int16
	Pad_cgo_3 [2]byte
	Flag      int32
	Login     [12]int8
	Spare     [4]int32
	Pad_cgo_4 [4]byte
}

type Upcred struct {
	Pc_lock   [72]int8
	Pc_ucred  uint64
	P_ruid    uint32
	P_svuid   uint32
	P_rgid    uint32
	P_svgid   uint32
	P_refcnt  int32
	Pad_cgo_0 [4]byte
}

type Uucred struct {
	Ref       int32
	UID       uint32
	Ngroups   int16
	Pad_cgo_0 [2]byte
	Groups    [16]uint32
}

type Vmspace struct {
	Dummy     int32
	Pad_cgo_0 [4]byte
	Dummy2    uint64
	Dummy3    [5]int32
	Pad_cgo_1 [4]byte
	Dummy4    [3]uint64
}

type Timeval struct {
	Sec       int64
	Usec      int32
	Pad_cgo_0 [4]byte
}

type Itimerval struct {
	Interval Timeval
	Value    Timeval
}

type ProcTaskInfo struct {
	Virtual_size      uint64
	Resident_size     uint64
	Total_user        uint64
	Total_system      uint64
	Threads_user      uint64
	Threads_system    uint64
	Policy            int32
	Faults            int32
	Pageins           int32
	Cow_faults        int32
	Messages_sent     int32
	Messages_received int32
	Syscalls_mach     int32
	Syscalls_unix     int32
	Csw               int32
	Threadnum         int32
	Numrunning        int32
	Priority          int32
}

type vinfoStat struct {
	Dev           uint32
	Mode          uint16
	Nlink         uint16
	Ino           uint64
	Uid           uint32
	Gid           uint32
	Atime         int64
	Atimensec     int64
	Mtime         int64
	Mtimensec     int64
	Ctime         int64
	Ctimensec     int64
	Birthtime     int64
	Birthtimensec int64
	Size          int64
	Blocks        int64
	Blksize       int32
	Flags         uint32
	Gen           uint32
	Rdev          uint32
	Qspare        [2]int64
}

type fsid struct {
	Val [2]int32
}

type vnodeInfo struct {
	Stat vinfoStat
	Type int32
	// Leopard puts Fsid before Pad, unlike the current amd64 definition.
	Fsid fsid
	Pad  int32
}

type vnodeInfoPath struct {
	Vi   vnodeInfo
	Path [1024]int8
}

type vnodePathInfo struct {
	Cdir vnodeInfoPath
	Rdir vnodeInfoPath
}
