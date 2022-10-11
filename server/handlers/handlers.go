package handlers

import (
  "fmt"
  "net/http"
  //"github.com/alphamystic/wheagle/utils"
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

func Edr(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"edr.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"edr.html",nil)
  return
}

func AptServe(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"apt.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"apt.html",nil)
  return
}

func Reports(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"reports.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"reports.html",nil)
  return
}

func Livediscovery(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"livediscovery.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"livediscovery.html",nil)
  return
}

func AssetsHandler(res http.ResponseWriter,req *http.Request){
  if req.Method != "GET"{
    tpl.ExecuteTemplate(res,"assets.html",nil)
    return
  }
  tpl.ExecuteTemplate(res,"assets.html",nil)
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

/* API HANDLES */

func Registerimplant(res http.ResponseWriter,req *http.Request){
  res.Write([]byte(fmt.Sprintf("[+] Registered")))
}

func Registerbot(res http.ResponseWriter,req *http.Request){
  res.Write([]byte(fmt.Sprintf("[+] Registered")))
}

func RegisterAgent(res http.ResponseWriter,req *http.Request){
  res.Write([]byte(fmt.Sprintf("[+] Registered")))
}
