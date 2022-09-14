package main

import (
  "os"
  "log"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/alphamystic/wheagle/beacon/protos/stream"
)


func main(){
  var (
    err error
    opts []grpc.DialOption
    conn *grpc.ClientConn
    client stream.AdminCLient
  )
  opts = append(opts,grpc.WithInsecure())
  if conn,err = grpc.Dial(fmt.Sprintf("localhost:%d",5001),opts...); err != nil{
    log.Fatalf("[-] Error connecting to admin server: %v",err)
  }
  defer conn.Close()
  client = stream.NewAdminClient(conn)
  var cmd = new(stream.Command)
  cmd.In = os.Args[1]
  ctx := context.Background()
  cmd,err = client.RunCommand(ctx,cmd)
  if err != nil{
    log.Fatalf("[-] Error running commads: %v",err)
  }
  fmt.Println(cmd.Out)


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
}


//https://www.youtube.com/watch?v=wqArtyADnXc
