package handlers

/*
  * Contains helper functions for server running without unnescescaryu imports
*/
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
var DarkMode bool

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
  DarkMode = true
  log.Println("[+] Parsed the templates")
}



func Authenticate(password,email string)(bool,string,string,string){
  var userEmail,userName,hash,userId,role string
  stmt := "SELECT userid,username,email,password,role FROM `siapp`.`users` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&userId,&userName,&userEmail,&hash,&role)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    utils.Logerror(e)
    return false,userId,userName,role
  }
  err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  if err != nil{
    e := utils.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    utils.Logerror(e)
    return false,userId,userName,role
  }
  fmt.Sprintf("User id is %s with role of %s",userId,role)
  return true,userId,userName,role
}
