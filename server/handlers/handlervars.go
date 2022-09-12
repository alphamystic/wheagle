package handlers

import (
  "log"
  "time"
  "html/template"
  "database/sql"
  "github.com/gorilla/sessions"
)


var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")
var tpl  *template.Template
var store = sessions.NewCookieStore([]byte("WHEAGLE Pentesting Suite System"))

func init(){
  var err error
  db,err = sql.Open("mysql", "root:@tcp(localhost:3306)/wheagle")
  if err != nil{
    log.Println("[-]  Error connecting to DB: ",err)
  }
  tpl,err = template.ParseGlob("/home/sam/Documents/3l0racle/wheagle/server/templates/*.html")
  if err != nil{
    log.Println("[-]  Error getting html templates: ",err)
  }
  log.Println("[+] Parsed the templates")
}
