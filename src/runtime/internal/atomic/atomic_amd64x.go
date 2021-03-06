// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64 amd64p32

package atomic

import "unsafe"

// The calls to nop are to keep these functions from being inlined.
// If they are inlined we have no guarantee that later rewrites of the
// code by optimizers will preserve the relative order of memory accesses.

//go:nosplit
func Load(ptr *uint32) uint32 {
	nop()
	return *ptr
}

//go:nosplit
func Loadp(ptr unsafe.Pointer) unsafe.Pointer {
	nop()
	return *(*unsafe.Pointer)(ptr)
}

//go:nosplit
func Load64(ptr *uint64) uint64 {
	nop()
	return *ptr
}

//go:noescape
func Xadd(ptr *uint32, delta int32) uint32

//go:noescape
func Xadd64(ptr *uint64, delta int64) uint64

//go:noescape
func Xadduintptr(ptr *uintptr, delta uintptr) uintptr

//go:noescape
func Xchg(ptr *uint32, new uint32) uint32

//go:noescape
func Xchg64(ptr *uint64, new uint64) uint64

//go:noescape
func Xchguintptr(ptr *uintptr, new uintptr) uintptr

//go:noescape
func And8(ptr *uint8, val uint8)

//go:noescape
func Or8(ptr *uint8, val uint8)

// NOTE: Do not add atomicxor8 (XOR is not idempotent).

//go:noescape
func Cas64(ptr *uint64, old, new uint64) bool

//go:noescape
func Store(ptr *uint32, val uint32)

//go:noescape
func Store64(ptr *uint64, val uint64)

// NO go:noescape annotation; see atomic_pointer.go.
func Storep1(ptr unsafe.Pointer, val unsafe.Pointer)
