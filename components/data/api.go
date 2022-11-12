package data

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Api struct{
  ApiKey string `json:"apikey"`
  OwnerId string  `json:"ownerid"`
  Active string `json:"active"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

//  	apikey 	ownerid 	active 	created_at 	updated_at
func CreateApiKey(a Api)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`apikey` (apikey,ownerid,active,created_at,updated_at) VALUES(?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTCPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create apikey. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.ApiKey,&a.OwnerId,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EXIAPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating API Key.")
  }
  return nil
}

func ListApiKeys(active bool)([]Api,error){
  stmt := "SELECT * FROM `wheagle`.`apikey` WHERE (`active` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,valid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELAPIK: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all api keys.")
  }
  defer rows.Close()
  var keys []Api
  for rows.Next(){
    var a Api
    err = rows.Scan(&a.ApiKey,&a.OwnerId,&a.Active,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESAPIK: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing apikeys.")
    }
    keys = append(keys,a)
  }
  return keys,nil
}

/*if !true{
  //Do the Thing
}
*/
func ValidateApiKey(key,ownerId string)bool{
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

func UpdateKey(ownerId,apiKey string)error{
  upStmt := "UPDATE `wheagle`.`apikey` SET (`apikey` = ? AND `ownerid` = ? AND `updated_at` = ?) WHERE (`ownerid` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPUAPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to update API Key.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(apiKey,ownerId,currentTime)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EUAPIK: activity id: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing update apikey.")
  }
  return nil
}

func DeactivateKey(ownerId,key string)error{
  upStmt := "UPDATE `wheagle`.`apikey` SET (`active` = ? `updated_at` = ?) WHERE (`apikey` = ? AND `ownerid` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPDAPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to deactivate API Key.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(false,currentTime,apiKey,ownerId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EDAPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing deactivate apikey.")
  }
  return nil
}


func ViewApiKey(keyId string)(*Api,error){
  var a Api
  row := db.QueryRow("SELECT * FROM `wheagle`.`apikey` WHERE apikey	 = ?;",keyId)
  err := row.Scan(&a.ApiKey,&a.OwnerId,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVAPIK with id %s %s",keyId,err))
    utils.Logerror(e)
    if err == sql.ErrNoRows {
      utils.Danger(fmt.Errorf("Api key doen't exist: %s", keyId))
      return errors.New("Requested Apikey doesn't exist")
    }
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing apikey of %s",keyId))
  }
  return &a,nil
}
