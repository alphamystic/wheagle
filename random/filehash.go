package main

import (
    "io/ioutil"
    "crypto/sha256"
    //"os"
    "log"
    "encoding/hex"
)

func main(){
  hash,err := HashFile("somfile")
  if err != nil{
    log.Fatal(err)
  }
  log.Println(hash)
}

func HashFile(filename string)(string,error){
  var hash string
  hasher := sha256.New()
  file,err := ioutil.ReadFile(filename)
  if err != nil{
    return hash,err
  }
  hasher.Write(file)
  hash = string(hex.EncodeToString(hasher.Sum(nil)))
  return hash,nil
}
