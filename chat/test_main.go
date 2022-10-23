package main


func main(){
  s := NewChatServer()
  go s.Run()
  listener,err := net.Listen("tcp",":8000")
  if err != nil{
    log.Fatalf("[-] Unable to start server: %s",err.Error())
  }
  defer listener.Close()
  log.Printf("[+] Started chat server at 8000")
  for{
    conn,err := listener.Accept()
    if err != nil{
      log.Printf("[-] Unable to accept connection: %s",err.Error())
      continue
    }
    go s.NewClient(conn)
  }
}
