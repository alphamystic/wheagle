package main

import (
  "fmt"
  "errors"
  "github.com/fatih/color"
)

func PrintTextInASpecificColorInBold(colorName,text string){
  switch colorName {
  case "Yellow":
    e := color.New(color.FgYellow, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "Red":
    e := color.New(color.FgRed, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "Green":
    e := color.New(color.FgGreen, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "Magenta":
    e := color.New(color.FgMagenta, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "White":
    e := color.New(color.FgWhite, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  case "Blue":
    e := color.New(color.FgBlue, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  default:
    e := color.New(color.FgCyan, color.Bold)
    fmt.Printf("[+]   ")
    e.Printf(text + "\n")
  }
}
//fmt.Printf("[+]   ")
func main(){
  c := color.New(color.FgCyan).Add(color.Underline)
c.Println("Initializing payloads.")

// Or just add them to New()
d := color.New(color.FgCyan, color.Bold)
d.Printf("Generating Random keys %s\n", "!!!.")
 e := color.New(color.FgYellow, color.Bold)
 e.Printf("[+]     Disabling Ransom capabilitiies.....")
// Mix up foreground and background colors, create new mixes!
red := color.New(color.FgRed)

boldRed := red.Add(color.Bold)
boldRed.Println("      WHEAGLE!!!!    The NEXTGEN C2")

whiteBackground := red.Add(color.BgWhite)
whiteBackground.Println("[+]    Run as sudo.")
err := errors.New("[+]    Failed tomload opensssl, start as server")
// Create a custom print function for convenience
redd := color.New(color.FgRed).PrintfFunc()
redd("Warning")
redd("Error: %s\n", err)

// Mix up multiple attributes
notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
    notice("[+]     INITILIZING")
notice("[+] NOTICE    " + "Don't Use this for malevolent purposes...")
  PrintTextInASpecificColorInBold("Yellow","Starting Listeners")
  PrintTextInASpecificColorInBold("Red", "Could not start dns loot server")
  PrintTextInASpecificColorInBold("Green", "Encypting TLS Coms")
  PrintTextInASpecificColorInBold("Magenta", "Setting Up UPNP Exfill")
  PrintTextInASpecificColorInBold("Blue", "Everything is set.....")
  PrintTextInASpecificColorInBold("White", "Starting wheagle on pty")
  PrintTextInASpecificColorInBold("Cyan", "Payloads at ~/bin/payloads")
  //PrintTextInASpecificColorInBold("ioionjn", "Should be cyan")
}
