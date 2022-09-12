package main

import(
  "os"
  "fmt"
  "github.com/alphamystic/wheagle/cli"
)


func main(){
  fmt.Println("[+] Starting Wheagle")
  defer os.Exit(0)
  cmd := cli.CommandLine{}
  cmd.Run()
}
