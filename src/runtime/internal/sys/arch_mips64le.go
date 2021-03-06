// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

const (
	TheChar       = '0'
	BigEndian     = 0
	CacheLineSize = 32
	PhysPageSize  = 16384
	PCQuantum     = 4
	Int64Align    = 8
	HugePageSize  = 0
	MinFrameSize  = 8
)

type Uintreg uint64
type Intptr int64 // TODO(rsc): remove
