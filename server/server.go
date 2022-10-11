package main

import(
  "fmt"
  "log"
  //"strconv"
  "net/http"
  "github.com/alphamystic/wheagle/server/handlers"
)



func main(){
  fs := http.FileServer(http.Dir("/home/sam/Documents/3l0racle/wheagle/server/static"))
  http.Handle("/static/",http.StripPrefix("/static",fs))
  handlers.StartHttpServerRoutes()
  /*port = 5000
  prt:= strconv.Itoa(port)
  if err != nil{
    log.Println("[-]  Error converting port into string: ",err)
  }*/
  go func(){
    for {
      //keep checking the temp table and update it.
    }
  }()
  go func(){
    fmt.Println("[+]    Starting HTTP server at 5000")
    err := http.ListenAndServe("0.0.0.0:5000",nil)
    if err !=  nil {
      log.Println("[-]   Error starting http server on port ")
    }
  }()
  fmt.Println("[+]    Starting HTTPS Server AT 5151")
  err := http.ListenAndServeTLS(
    "0.0.0.0:5151",
    "certs/wheagle_server_cert.pem",
    "certs/wheagle_server_key.pem",
    nil,
  )
  if err != nil {
    log.Fatal("Error creating server. ", err)
  }
}

//connectplus


/*
func main() {

    // Show on console the application stated
    log.Println("Server started on: http://localhost:9000")
    main_server := http.NewServeMux()
    main_server.HandleFunc("/",func(res http.ResponseWriter,req *http.Request){
      fmt.Fprint(res,"EDR For pentesting: "+ req.Host +" Says Hallo")
    })

    //Creating sub-domain
    server1 := http.NewServeMux()
    server1.HandleFunc("/", server1func)

    server2 := http.NewServeMux()
    server2.HandleFunc("/", server2func)

    //Running First Server
    go func() {
        log.Println("Server started on: http://localhost:9001")
        http.ListenAndServe("localhost:9001", server1)
    }()

    //Running Second Server
    go func() {
        log.Println("Server started on: http://localhost:9002")
        http.ListenAndServe("localhost:9002", server2)
    }()

    //Running Main Server
    http.ListenAndServe("localhost:9000", main_server)
}

func server1func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running First Server")
}

func server2func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Running Second Server")
}
*/
