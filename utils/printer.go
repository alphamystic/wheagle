package utils

import (
  "fmt"
  "strings"
  "github.com/fatih/color"
)

//log error of errors
func Logerror(e error){
  if e != nil {
    red := color.New(color.FgRed)
    whiteBackground := red.Add(color.BgWhite)
    whiteBackground.Println(e)
  }
}

func Danger(e error){
  if e != nil{
    red := color.New(color.FgRed)
    boldRed := red.Add(color.Bold)
    boldRed.Println(e)
  }
}

func Notice(text string){
  notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
  notice("[+]NOTICE:    " + text + "\n")
}

func Warning(text string){
  redd := color.New(color.FgRed).PrintfFunc()
  redd("[-] WARNING:   ")
  redd(" %s\n", text)
}

func PrintInformation(text string){
  e := color.New(color.FgYellow, color.Bold)
  e.Printf(text)
}

func PrintTextInASpecificColorInBold(colorName,text string){
  switch strings.ToLower(colorName) {
  case "yellow":
    e := color.New(color.FgYellow, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "red":
    e := color.New(color.FgRed, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "green":
    e := color.New(color.FgGreen, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "magenta":
    e := color.New(color.FgMagenta, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "white":
    e := color.New(color.FgWhite, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "blue":
    e := color.New(color.FgBlue, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  default:
    e := color.New(color.FgCyan, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  }
}
