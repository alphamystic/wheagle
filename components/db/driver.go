package db

import (
  "log"
  "fmt"
  "encoding/json"
  "net/http"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)
var dbCon = "root:@tcp(localhost:3306)/wheagle?charset=utf8mb4"

var db,_ = gorm.Open(mysql.Open(dbCon),&gorm.Config{})

type User struct{
  Email string
  Password string
}

func lmain(){
  http.HandleFunc("/createuser",DBCreate)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func DBCreate(res http.ResponseWriter, req *htpp.Request){
  user := User{
    Email: "sam@mail.com",
    Password: "pass",
  }
  db.Create(&user)
  if err := db.Create(&user).Error; err != nil {
    log.Fatalln((err))
  }
  json.NewEncoder(res).Encode(User)
  fmt.Println("It worked: ",user)
}
