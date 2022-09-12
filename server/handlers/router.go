package handlers

import(
  "fmt"
  "net/http"
)



func StartHttpServerRoutes(){
  /*var err error
  fs := http.FileServer(http.Dir("/home/sam/Documents/3l0racle/wheagle/server/assets"))
  http.Handle("/assets/",http.StripPrefix("/assets",fs))*/
  http.HandleFunc("/",Home)
  http.HandleFunc("/login",Login)
  http.HandleFunc("/1",func(res http.ResponseWriter,req *http.Request){
    fmt.Fprint(res,"EDR For pentesting: "+ req.Host +" Says Hallo")
  })
  http.HandleFunc("/2",func(res http.ResponseWriter,req *http.Request){
    fmt.Fprint(res,"EDR 2 For pentesting: "+ req.Host +" Says Hallo")
  })
}
