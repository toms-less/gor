package server

import (
	context "context"

	rt "gor/build/protos/runtime"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// RuntimeServer interface.
type RuntimeService interface {
	Call(context.Context, *rt.RuntimeRequest) (*rt.RuntimeResponse, error)
}

// RuntimeServer
type RuntimeServer struct {
}

func (*RuntimeServer) Call(ctx context.Context, request *rt.RuntimeRequest) (*rt.RuntimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method not implemented")
}
