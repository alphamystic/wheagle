package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
)


func main() {
  for {
    reader := bufio.NewReader(os.Stdin)
    fmt.Printf("Name: ")
    name, _ := reader.ReadString('\n')
    name = strings.Trim(name, "\n")
    fmt.Println(name)
  }
}
