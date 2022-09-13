package main

import (
  "log"
  "net"

  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


func main(){
  lis,err := net.Listen("tcp",":5001")
  if err != nil{
    log.Fatalf("[-] Failed to start listening on port 5001 for beacons: %v",err)
  }

  s := stream.Server{}
  grpcServer := grpc.NewServer()
  stream.RegisterChatServiceServer(grpcServer,&s)
  if err := grpcServer.Serve(lis); err != nil{
    log.Fatalf("[-] Failed to start grpc Server at 95001. %v ",err)
  }
}
