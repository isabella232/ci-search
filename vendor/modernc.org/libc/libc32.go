// Copyright 2020 The Libc Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build 386 arm

package libc // import "modernc.org/libc"

type (
	RawMem   [1<<31 - 1]byte
	RawMem64 [1<<27 - 1]uint64

	// 32-5*4 = 12 bytes left to pad
	stackHeaderPadding struct {
		a uintptr
		b uintptr
		c uintptr
	}
)

type bits []int

func newBits(n int) (r bits)  { return make(bits, (n+31)>>5) }
func (b bits) has(n int) bool { return b != nil && b[n>>5]&(1<<uint(n&31)) != 0 }
func (b bits) set(n int)      { b[n>>5] |= 1 << uint(n&31) }
