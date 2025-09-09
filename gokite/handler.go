package main

import (
	"context"

	"github.com/kelein/cookbook/gokite/kitgen/service"
)

// EchoFacade implements the last service interface defined in the IDL.
type EchoFacade struct{}

// Echo implements the EchoFacade interface.
func (s *EchoFacade) Echo(ctx context.Context, req *service.Request) (resp *service.Response, err error) {
	// TODO: Your code here...
	return
}

// EchoServiceFacade implements the last service interface defined in the IDL.
type EchoServiceFacade struct{}

// Echo implements the last service interface defined in the IDL.
func (s *EchoServiceFacade) Echo(ctx context.Context, req *service.Request) (resp *service.Response, err error) {
	// TODO: Your code here...
	return
}
