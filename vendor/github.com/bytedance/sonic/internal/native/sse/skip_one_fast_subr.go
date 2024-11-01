//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package sse

import (
	"github.com/bytedance/sonic/loader"
)

const (
	_entry__skip_one_fast = 128
)

const (
	_stack__skip_one_fast = 136
)

const (
	_size__skip_one_fast = 3348
)

var (
	_pcsp__skip_one_fast = [][2]uint32{
		{0x1, 0},
		{0x6, 8},
		{0x8, 16},
		{0xa, 24},
		{0xc, 32},
		{0xd, 40},
		{0x11, 48},
		{0x25c, 136},
		{0x25d, 48},
		{0x25f, 40},
		{0x261, 32},
		{0x263, 24},
		{0x265, 16},
		{0x266, 8},
		{0x267, 0},
		{0xd14, 136},
	}
)

var _cfunc_skip_one_fast = []loader.CFunc{
	{"_skip_one_fast_entry", 0, _entry__skip_one_fast, 0, nil},
	{"_skip_one_fast", _entry__skip_one_fast, _size__skip_one_fast, _stack__skip_one_fast, _pcsp__skip_one_fast},
}
