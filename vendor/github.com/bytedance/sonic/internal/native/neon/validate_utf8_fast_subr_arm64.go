//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package neon

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __validate_utf8_fast_entry__() uintptr

var (
	_subr__validate_utf8_fast uintptr = __validate_utf8_fast_entry__() + 0
)

const (
	_stack__validate_utf8_fast = 48
)

var (
	_ = _subr__validate_utf8_fast
)

const (
	_ = _stack__validate_utf8_fast
)