//go:build !noasm || !appengine
// +build !noasm !appengine

// Code generated by asm2asm, DO NOT EDIT.

package neon

//go:nosplit
//go:noescape
//goland:noinspection ALL
func __html_escape_entry__() uintptr

var (
	_subr__html_escape uintptr = __html_escape_entry__() + 32
)

const (
	_stack__html_escape = 32
)

var (
	_ = _subr__html_escape
)

const (
	_ = _stack__html_escape
)
