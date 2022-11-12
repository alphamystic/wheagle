package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type User struct {
  Name string
  Id string
  Age int
  Gender []byte
}


func main(){
  fmt.Println("[+] Encoding ")
  user := User{
    Name:"Cynthia",
    Id: "GF Status Updated",
    Age: 22,
    Gender: []byte("female"),
  }
  var buf bytes.Buffer
  enc := gob.NewEncoder(&buf)
  if err := enc.Encode(user); err != nil{
    log.Fatal(err)
  }
  input := buf.Bytes()
  fmt.Println(input)
  fmt.Println("")
  buf2 := bytes.NewBuffer(input)
  dec := gob.NewDecoder(buf2)
  var u User
  if err := dec.Decode(&u); err != nil{
    log.Fatal (err)
  }
  fmt.Println("[+] Decoding ")
  fmt.Println(u.Name)
  fmt.Println(u.Id)
  fmt.Println(u.Age)
  fmt.Println(string(u.Gender))
}
