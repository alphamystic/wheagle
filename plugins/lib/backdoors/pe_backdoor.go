package main

import (
  "os"
  "fmt"
  "plugin"
  "io/ioutil"
  "github.com/alphamystic/wheagle/components/orca"
)

const PluginsDir = "../../plugins/lib/"

// This should be implemented into core operator
func main(){
  var(
    err error
    p *plugin.Plugin
    n plugin.Symbol
    bd orca.BACKDOOR
    bdRes *orca.Backdoored
    files []os.FileInfo
  )
  if files,err = ioutil.ReadDir(PluginsDir); err != nil{
    log.Fatalln("Error reading plugins dir: ",err)
  }
  for idx := range files {
    fmt.Println("[+] Available plugins are:")
    fmt.Println("     >",files[idx].Name())
    if p,err = plugin.Open(PluginsDir + "/" + files[idx].Name()); err !=  nil{
      log.Fatal("")
    }
  }
}
