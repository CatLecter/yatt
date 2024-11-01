//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__vstring = 48
)

const (
	_stack__vstring = 104
)

const (
	_size__vstring = 2392
)

var (
	_pcsp__vstring = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0x11, 48},
		{0x893, 104},
		{0x894, 48},
		{0x896, 40},
		{0x898, 32},
		{0x89a, 24},
		{0x89c, 16},
		{0x89d, 8},
		{0x89e, 0},
		{0x958, 104},
	}
)

var _cfunc_vstring = []loader.CFunc{
	{"_vstring_entry", 0, _entry__vstring, 0, nil},
	{"_vstring", _entry__vstring, _size__vstring, _stack__vstring, _pcsp__vstring},
}
