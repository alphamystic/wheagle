package handlers

import (
  "log"
  "fmt"
  "time"
  "html/template"
  "database/sql"
  "github.com/gorilla/sessions"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"

  "github.com/alphamystic/wheagle/utils"
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



func Authenticate(password,email string)(bool,string){
  var userEmail,hash,userId string
  stmt := "SELECT email,userid,password FROM `prezo`.`users` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  //defer db.Close()
  err := row.Scan(&userEmail,&userId,&hash)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    utils.Logerror(e)
    return false,userId
  }
  err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  if err != nil{
    e := utils.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    utils.Logerror(e)
    return false,userId
  }
  fmt.Println("User id during authentication: ",userId)
  return true,userId
}
