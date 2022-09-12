package cli

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
  "github.com/alphamystic/wheagle/server/routes"
)
type CommandLine struct{}


func (cli *CommandLine) startServer(port int){
  go func (p int){
    routes.StartHttpServer(p)
  }(port)
}


func (cli *CommandLine) Run(){
  fmt.Println("[+]    Running Wheagle")
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("[+]  Enter Server port number: ")
    port, _ := reader.ReadString('\n')
    port = strings.Trim(port, "\n")
    prt,_ := strconv.Atoi(port)
    cli.startServer(5000)
    fmt.Println(prt)
  }
}
