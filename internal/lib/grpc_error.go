package lib

import "google.golang.org/grpc/codes"

type GRPCError struct {
	Code codes.Code
	Msg  string
}

func NewGRPCError(code codes.Code, msg string) *GRPCError {
	return &GRPCError{
		Code: code,
		Msg:  msg,
	}
}
