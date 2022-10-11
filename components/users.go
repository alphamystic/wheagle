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
  Active bool
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
  var m Minion
  row := db.QueryRow("SELECT * FROM `malbot`.`minion` WHERE minionid	 = ?;",minionId)
  err := row.Scan(&m.MinionId,&m.Name,&m.UName,&m.UserId,&m.GroupId,&m.HomeDir,&m.Os,&m.Description,&m.Installed,&m.MotherShipId,&m.LastSeen,&m.CreatedAt,&m.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("EVM %s %s",minionId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing minon with id of %s",minionId))
  }
  return &m,nil
}

func ListMyUsers(ownerId string,active bool)([]User,error){
  return nil,nil
}


func AdminListUsers(active bool)([]User,error){
  return nil,nil
}

func ListAllUsers()([]User,error){
  stmt := "SELECT * FROM `malbot`.`minion` ORDER BY updated_at ASC;"
  rows,err := db.Query(stmt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ELAMM: %s",err))
    logError(e)
    return nil,errors.New("Server encountered an error while listing all my minions.")
  }
  defer rows.Close()
  var minions []Minion
  for rows.Next(){
    var m Minion
    err = rows.Scan(&m.MinionId,&m.Name,&m.UName,&m.UserId,&m.GroupId,&m.HomeDir,&m.Os,&m.Description,&m.Installed,&m.MotherShipId,&m.LastSeen,&m.CreatedAt,&m.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("ESFAMM: %s",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all my minions.")
    }
    minions = append(minions,m)
  }
  return minions,nil
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
