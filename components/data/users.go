package data


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

// 	userid 	ownerid 	username 	email 	password 	active 	anonymous 	verified 	admin 	created_at 	updated_at
func CreateUser(u User)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`user` (userid,ownerid,username,email,password,active,anonymous,verified,admin,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTCU: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create user. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&u.UserId,&u.OwnerId,&u.UserName,&u.Email,&u.Password,&u.Active,&u.Anonymous,&u.Verify,&u.Admin,&u.CreatedAt,&u.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECU: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating user.")
  }
  return nil
}


func ViewUser(userId string)(*User,error){
  var u User
  row := db.QueryRow("SELECT * FROM `wheagle`.`user` WHERE userid	 = ?;",userId)
  err := row.Scan(&u.UserId,&u.OwnerId,&u.UserName,&u.Email,&u.Password,&u.Active,&u.Anonymous,&u.Verify,&u.Admin,&u.CreatedAt,&u.UpdatedAt)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("EVU %s %s",userId,err))
    logError(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing user with id of %s",userId))
  }
  return &u,nil
}

func ListMyUsers(ownerId string,active bool)([]User,error){
  stmt := "SELECT * FROM `wheagle`.`user` WHERE (`active` = ? AND `ownerid` = ?) ORDER BY updated_at ASC;"
  rows,err := db.Query(stmt,active,ownerId)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ELAU: %s",err))
    logError(e)
    return nil,errors.New("Server encountered an error while listing all specified users.")
  }
  defer rows.Close()
  var users []User
  for rows.Next(){
    var u User
    err = rows.Scan(&u.UserId,&u.OwnerId,&u.UserName,&u.Email,&u.Password,&u.Active,&u.Anonymous,&u.Verify,&u.Admin,&u.CreatedAt,&u.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("ESWU: %s",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all my minions.")
    }
    users = append(users,u)
  }
  return users,nil
}


func AdminListUsers(active bool)([]User,error){
  return nil,nil
}

func ListAllUsers(active bool)([]User,error){
  stmt := "SELECT * FROM `wheagle`.`user` WHERE (`active` = ?) ORDER BY updated_at ASC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := LogErrorToFile("sql",fmt.Sprintf("ELAU: %s",err))
    logError(e)
    return nil,errors.New("Server encountered an error while listing all specified users.")
  }
  defer rows.Close()
  var users []User
  for rows.Next(){
    var u User
    err = rows.Scan(&u.UserId,&u.OwnerId,&u.UserName,&u.Email,&u.Password,&u.Active,&u.Anonymous,&u.Verify,&u.Admin,&u.CreatedAt,&u.UpdatedAt)
    if err != nil{
      e := LogErrorToFile("sql",fmt.Sprintf("ESWU: %s",err))
      logError(e)
      return nil,errors.New("Server encountered an error while listing all my minions.")
    }
    users = append(users,u)
  }
  return users,nil
}

func MarkUserAsInActive(ownerId,userId string)error{
  return nil
}


func CheckIfOwner(userId,ownerId string)bool{
  var key,user string
  stmt := "SELECT apikey,ownerid FROM `wheagle`.`apikey` WHERE (`apikey`= ? AND `ownerid` = ?);"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&key,&user)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning apikey rows %s",err))
    utils.LogError(e)
    return false
  }
  return true
}

func CheckIfAdmin(adminId string)bool{
  var active,admin string
  stmt := "SELECT active,admin FROM `wheagle`.`user` WHERE (`userid`= ?);"
  row := db.QueryRow(stmt,adminId)
  err := row.Scan(&active,&admin)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning apikey rows %s",err))
    utils.LogError(e)
    return false
  }
  if !active && !admin{
    utils.Warning(fmt.Sprintf("Non admin with id %s tried accessing admin stuff",adminId))
    return false
  }
  return true
}


func Authenticate(password,email string)(bool,string,string){
  var userEmail,userName,hash,userId,role string
  stmt := "SELECT userid,username,email,password,role FROM `wheagle`.`user` WHERE email = ?;"
  row := db.QueryRow(stmt,email)
  err := row.Scan(&userId,&userName,&userEmail,&hash)
  if err != nil{
    if err == sql.ErrNoRows {
      utils.Waring(fmt.Sprintf("A none user with email %s tried accessing server",email))
      return false,userName,userId
    }
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    utils.LogError(e)
    return false,userName,userId
  }
  err = bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
  if err != nil{
    e := utils.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    utils.LogError(e)
    return false,userName,userId
  }
  return true,userName,userId
}
