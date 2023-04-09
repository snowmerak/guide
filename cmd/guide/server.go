package main

import (
	"context"
	"github.com/snowmerak/guide/src/service"
)

type ServiceServer struct {
	service.UnimplementedServiceServer
}

func (s ServiceServer) Send(ctx context.Context, datagram *service.Datagram) (*service.Blank, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceServer) Request(ctx context.Context, datagram *service.Datagram) (*service.Datagram, error) {
	//TODO implement me
	panic("implement me")
}
