//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__lookup_small_key = 96
)

const (
	_stack__lookup_small_key = 56
)

const (
	_size__lookup_small_key = 794
)

var (
	_pcsp__lookup_small_key = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0xe, 48},
		{0x30c, 56},
		{0x30d, 48},
		{0x30f, 40},
		{0x311, 32},
		{0x313, 24},
		{0x315, 16},
		{0x316, 8},
		{0x31a, 0},
	}
)

var _cfunc_lookup_small_key = []loader.CFunc{
	{"_lookup_small_key_entry", 0, _entry__lookup_small_key, 0, nil},
	{"_lookup_small_key", _entry__lookup_small_key, _size__lookup_small_key, _stack__lookup_small_key, _pcsp__lookup_small_key},
}