//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package avx2

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__skip_array = 704
)

const (
	_stack__skip_array = 208
)

const (
	_size__skip_array = 15888
)

var (
	_pcsp__skip_array = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0x14, 48},
		{0x3b41, 208},
		{0x3b42, 48},
		{0x3b44, 40},
		{0x3b46, 32},
		{0x3b48, 24},
		{0x3b4a, 16},
		{0x3b4b, 8},
		{0x3b4f, 0},
		{0x3e10, 208},
	}
)

var _cfunc_skip_array = []loader.CFunc{
	{"_skip_array_entry", 0, _entry__skip_array, 0, nil},
	{"_skip_array", _entry__skip_array, _size__skip_array, _stack__skip_array, _pcsp__skip_array},
}