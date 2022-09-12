package handlers

import (
  "fmt"
  "net/http"
  "golang.org/x/crypto/bcrypt"
  "github.com/alphamystic/wheagle/utils"
  //"github.com/alphamystic/wheagle/components"
)




func Home(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"index.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"index.html",nil)
  return
}

func Login(res http.ResponseWriter,req *http.Request){
  if req.Method == "POST"{
    req.ParseForm()
    pass := req.FormValue("password")
    email := req.FormValue("email")
    isAuth,paId := Authenticate(pass,email)
    if !isAuth {
      tpl.ExecuteTemplate(res,"login.html","Wrong username or password")
      return
    }
    //set session
    session,_ := store.Get(req,"session")
    session.Values["PaId"] = paId
    session.Save(req,res)
    //redirect to dashboard or get the dash data and execute dash
    http.Redirect(res,req,"/",http.StatusSeeOther)
    return
  }
  tpl.ExecuteTemplate(res,"login.html",nil)
  return
}

func Logout(res http.ResponseWriter, req *http.Request){
  session,_ := store.Get(req,"session")
  delete(session.Values,"PaId")
  session.Save(req,res)
  tpl.ExecuteTemplate(res,"login.html","Logged Out ADIOS!!!")
  return
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
