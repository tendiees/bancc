package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/tendiees/bancc/bancc"
	"github.com/tendiees/bancc/ep"
	"github.com/tendiees/bancc/svc"
	"github.com/tendiees/bancc/tp"
)

func main() {
	base := grpc.NewServer()
	server := tp.NewGRPC(ep.Set{
		Ping: ep.MakePing(svc.New()),
	})
	bancc.RegisterBanccServer(base, server)
	if l, err := net.Listen("tcp", ":8080"); err == nil {
		log.Fatal(base.Serve(l))
	}
}
