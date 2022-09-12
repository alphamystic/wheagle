package main

import(
  "fmt"
  "log"
  //"strconv"
  "net/http"
  "github.com/alphamystic/wheagle/server/handlers"
)



func main(){
  fs := http.FileServer(http.Dir("/home/sam/Documents/3l0racle/wheagle/server/assets"))
  http.Handle("/assets/",http.StripPrefix("/assets",fs))
  handlers.StartHttpServerRoutes()
  /*port = 5000
  prt:= strconv.Itoa(port)
  if err != nil{
    log.Println("[-]  Error converting port into string: ",err)
  }*/
  fmt.Println("[+]    Starting server")
  err := http.ListenAndServe("0.0.0.0:5000",nil)
  if err !=  nil {
    log.Println("[-]   Error starting http server on port ")
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
