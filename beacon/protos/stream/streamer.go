package stream
import (
  "log"
  "golang.org/x/net/context"
)


type Server struct{
}


func (s *Server) SayHello(ctx context.Context, message *Message)(*Message,error){
  //Don't log messages but prrint them out instead
  log.Printf("[+] Received message body from client: %s",message.Body)
  return &Message{Body:"Hello too"},nil
}
