package components

/*
  * Implements all kinds of proxy from:
    * TCP Proxy
    * UDP Proxy
    * DNS Proxy
    * DNS Proxy
*/

import (
  "io"
  "fmt"
  "net"
  "log"
)

type Proxify interface{
  TCPProxy() ...interface{}
  UDPProxy() ...interface{}
  HTTPProxy() ...interface{}
  DNSProxy() ...interface{}
}


type ProxyTypes struct{
  HostAddress string
  TargetAddress string
}

//implement this for each proxy type
func (pt ProxyTypes) Proxify() *net.Conn{
  return nil
}

func (pt ProxyTypes) TCPProxy()...interface{}{
  listener,err := net.Listen("tcp",pt.HostAddress)
  if err != nil{panic(err);return errors.New(error.Error)}
  for {
    conn,err := listener.Accept()
    if err != nil{return errors.New("Error accepting connection.");continue}
    go func(){
      conn2,err := net.Dial("tcp",pt.TargetAddress)
      if err != nil{
        return errors.New("Error dialing remote address.")
      }
      go io.Copy(conn,conn2)
      io.Copy(conn2,conn)
      conn2.Close()
      conn.Close()
    }()
  }
}

func RunProxy(name,host,trg string){
  switch name{
    case name == "tcp":
      fmt.Println("[+] Implementing a tcp proxy")
      p := Proxify{HostAddress:host,TargetAddress:trg}
      p.TCPProxy()
    default:
      fmt.Println("Imlementing HTTP Proxy.")
  }
}
