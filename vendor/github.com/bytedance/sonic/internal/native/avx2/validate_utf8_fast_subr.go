//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__validate_utf8_fast = 272
)

const (
	_stack__validate_utf8_fast = 176
)

const (
	_size__validate_utf8_fast = 2656
)

var (
	_pcsp__validate_utf8_fast = [][2]uint32{
		{0x1, 0},
		{0x5, 8},
		{0xc, 16},
		{0x6aa, 176},
		{0x6ab, 16},
		{0x6ac, 8},
		{0x6b0, 0},
		{0x7d3, 176},
		{0x7d4, 16},
		{0x7d5, 8},
		{0x7d9, 0},
		{0xa60, 176},
	}
)

var _cfunc_validate_utf8_fast = []loader.CFunc{
	{"_validate_utf8_fast_entry", 0, _entry__validate_utf8_fast, 0, nil},
	{"_validate_utf8_fast", _entry__validate_utf8_fast, _size__validate_utf8_fast, _stack__validate_utf8_fast, _pcsp__validate_utf8_fast},
}
