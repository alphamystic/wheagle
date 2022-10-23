package wokers

import (
  "os"
  "log"
  "time"
)

func ChangeFileOwnership(filename string,daysToExtend int)error{
  err := os.Chmod(filename,0777)//update input to take permisions too
  if err != nil{
    return err
  }
  //change ownership
  err = os.Chown(filename,os.Getuid(),os.Getgid())
  if err != nil{
    return err
  }
  //change the timestamp
  extraDays := time.Now().Add(daysToExtend * time.Hour)
  err = os.Chtimes(filename,extraDays,extraDays)
  if err !=  nil{
    return err
  }
  return nil
}
