package components


/*
  * Defines the different kind of ussers that can be able to use wheagle
*/

import(
  "errors"
  "database/sql"
  "github.com/wheagle/utils"
  "golang.org/x/crypto/bcrypt"
)

type User struct{
  UserId string
  OwnerId string
  UserName string
  Email string
  Password string
  Active string
  Anonymous bool
  Verify bool
  Admin bool
  CreatedAt string
  UpdatedAt string
}

func CreateUser(u User)error{
  return  nil
}

func ViewUser(userId string)(*User,error){
  return nil.nil
}

func ListMyUsers(ownerId string,active bool)([]User,error){
  return nil,nil
}


func AdminListUsers(active bool)([]User,error){
  return nil,nil
}

func MarkUserAsInActive(ownerId,userId string)error{
  return nil
}


func CheckIfOwner(userId,ownerId string)bool{
  return false
}

func CheckIfAdmin(adminId string)bool{
  return false
}


func Authenticate(password,email string)(bool,string,string){
  var userEmail,hash,userId,role string
  stmt := "SELECT userid,email,password,role FROM `siapp`.`users` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&userId,&userEmail,&hash,&role)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    utils.LogError(e)
    return false,userId,role
  }
  err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  if err != nil{
    e := utils.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    utils.LogError(e)
    return false,userId,role
  }
  fmt.Sprintf("User id is %s with role of %s",userId,role)
  return true,userId,role
}
