package workers

/*
  * Contains variables for components functioning
*/

import(
  "time"
  "database/sql"
)

var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
var isDDoS bool


func SetDDoSMode(mode bool){
  if mode == true {
    isDDoS == true
    return
  }
  isDDoS == false
  return
}

func SignPlugin(name,oid string) (string,error){
  return "",nil
}
