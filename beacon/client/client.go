package main

import (
//  "os"
  "log"
  "fmt"
  "flag"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


func main(){
  var (
    err error
    opts []grpc.DialOption
    conn *grpc.ClientConn
    client stream.AdminClient
  )
  _ = flag.String("cmnd","foo","Command to execute")
  opts = append(opts,grpc.WithInsecure())
  if conn,err = grpc.Dial(fmt.Sprintf("localhost:%d",5001),opts...); err != nil{
    log.Fatalf("[-] Error connecting to admin server: %v",err)
  }
  defer conn.Close()
  client = stream.NewAdminClient(conn)
  var cmd *stream.Command = new(stream.Command)
  //var cmd = new(stream.Command)
  //fmt.Sprintf("Coomand give: %s",*cmnd)
  cmd.In = "id"//*cmnd
  ctx := context.Background()
  cmd,err = client.RunCommand(ctx,cmd)
  //fmt.Println("This was the command given: ",cmd.In)
  if err != nil{
  //  log.Fatal(err)
    log.Fatalf("[-] Error running commads: %v",err)
  }
  fmt.Println(cmd.Out)
}

func RegisterListener()error{
  return nil
}


//CHAT SERVER CLIENT
/*conn, err := grpc.Dial(":5003",grpc.WithInsecure())
if err != nil{
  log.Fatalf("[-] Couldn't connect: %v",err)
}
defer conn.Close()
c := stream.NewChatServiceClient(conn)
message := stream.Message{
  Body:"Saying hello from the client!",
}
resp,err := c.SayHello(context.Background(),&message )
if err != nil {
  log.Fatalf("[-] Error sayiong hallo: %v",err)
}
log.Printf("[+] Response from server: %s",resp.Body)*/

//https://www.youtube.com/watch?v=wqArtyADnXc
