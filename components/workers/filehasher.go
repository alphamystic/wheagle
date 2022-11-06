package workers

import (
    "io/ioutil"
    "crypto/sha256"
    "encoding/hex"
)

func GetFileHash256(filename string)(string,error){
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


//  @TODO Add more hsh techniques
