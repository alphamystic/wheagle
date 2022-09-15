package main

import (
  "fmt"
  "log"
  "net"

  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


func main(){
  var (
    implantListener,adminListener net.Listener
    err error
    opts []grpc.ServerOption
    work,output chan *stream.Command
  )
  //initialize a work channels for commands
  work,output = make(chan *stream.Command),make(chan *stream.Command)
  implant := stream.NewImplantServer(work,output)
  admin := stream.NewAdminServer(work, output)
  //start implant listener
  if implantListener,err = net.Listen("tcp",fmt.Sprintf(":%d",5002)); err != nil{
    log.Fatalf("[-]   Failed to start implant listener: %v",err)
  }
  //start admin listener
  if adminListener,err = net.Listen("tcp",fmt.Sprintf(":%d",5001)); err != nil{
    log.Fatalf("[-]   Failed to start admin listener: %v",err)
  }
  chatListener,err := net.Listen("tcp",fmt.Sprintf(":%d",5003))
  if err != nil{
    log.Fatalf("[-] Failed to start chat server: %v",err)
  }
  s := stream.ChatServer{}
  grpcChatServer := grpc.NewServer()
  grpcAdminServer,grpcImplantServer := grpc.NewServer(opts...),grpc.NewServer(opts...)
  stream.RegisterImplantServer(grpcImplantServer,implant)
  stream.RegisterAdminServer(grpcAdminServer,admin)
  stream.RegisterChatServiceServer(grpcChatServer,&s) //probaly use & at s
  go func(){
    grpcImplantServer.Serve(implantListener)
  }()
  go func(){
    grpcImplantServer.Serve(adminListener)
  }()
  grpcChatServer.Serve(chatListener)

/*
  lis,err := net.Listen("tcp",":5001")
  if err != nil{
    log.Fatalf("[-] Failed to start listening on port 5001 for beacons: %v",err)
  }
  fmt.Println("[+]  Initializing beacon server")
  s := stream.Server{}
  grpcServer := grpc.NewServer()
  stream.RegisterChatServiceServer(grpcServer,&s)
  if err := grpcServer.Serve(lis); err != nil{
    log.Fatalf("[-] Failed to start grpc Server at 95001. %v ",err)
  }*/
}
