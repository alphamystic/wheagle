package workers

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Plugin struct {
  Owner string
  Name string
  Hash string
  PType  int
  Decsription string
  Active bool
  Signed bool
  CreatedAt string
  UpdatedAt string
}

//Plugin types (using integers to avoid user defining their own string type of plugin)
const  (
  cleaner = iota + 1
  malwareScanner
  vulnerabilityScanner
)

// owner 	name 	hash 	plugin_type 	description 	active 	signed 	created_at 	updated_at
func DBCreatePlugin(p Plugin)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`plugins` (owner,name,hash,palugin_type,description,active,signed,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPCP: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create plugin. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&p.Owner,&p.Name,&p.Name,&p.Hash,&p.Type,&p.Description,&p.Active,&p.Signed,&p.CreatedAt,&p.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ECDP: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating db plugin.")
  }
  return nil
}

func ListPlugins(ownerId string, admin bool)([]Plugin,error){
  stmt := "SELECT * FROM `wheagle`.`plugins` WHERE (`active` = ? AND `owner` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,valid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOP: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all owner plugins.")
  }
  defer rows.Close()
  var pls []Plugin
  for rows.Next(){
    var p Plugin
    err = rows.Scan(&p.Owner,&p.Name,&p.Name,&p.Hash,&p.Type,&p.Description,&p.Active,&p.Signed,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOP: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing owner plugins.")
    }
    pls = append(pls,p)
  }
  return pls,nil
}

func ListAllPlugins(active bool)([]Plugin,error){
  stmt := "SELECT * FROM `wheagle`.`plugins` WHERE (`active` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,valid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOP: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all plugins as admin.")
  }
  defer rows.Close()
  var pls []Plugin
  for rows.Next(){
    var p Plugin
    err = rows.Scan(&p.Owner,&p.Name,&p.Name,&p.Hash,&p.Type,&p.Description,&p.Active,&p.Signed,&p.CreatedAt,&p.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOP: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all plugins as admin.")
    }
    pls = append(pls,p)
  }
  return pls,nil
}

func ViewPlugin(hash string)(*Plugin,error){
  var p Plugin
  row := db.QueryRow("SELECT * FROM `wheagle`.`plugins` WHERE hash = ?;",hash)
  err := row.Scan(&p.Owner,&p.Name,&p.Name,&p.Hash,&p.Type,&p.Description,&p.Active,&p.Signed,&p.CreatedAt,&p.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVAPIK with id %s %s",hash,err))
    utils.Logerror(e)
    if err == sql.ErrNoRows {
      utils.Danger("No plugin with specified hash: "+ hash)
      return errors.New("Requested plugin doesn't exist")
    }
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing plugin of hash %s",hash))
  }
  return &p,nil
}

func DeactivatePlugin(currentTime,hash,ownerId string,admin bool)error{
  upStmt := "UPDATE `wheagle`.`plugins` SET (`active` = ? `updated_at` = ?) WHERE (`hash` = ? AND `owner` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPDP: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to deactivate plugin.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(false,currentTime,hash,ownerId)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EDP: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing deactivate plugin.")
  }
  return nil
}

//so does unsigning
func MarkAsSigned(ownerId,hash string,sign bool)error{
  upStmt := "UPDATE `wheagle`.`plugin` SET (`signed` = ? `updated_at` = ?) WHERE (`hash` = ? AND `owner` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPSAP: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to sign or unsign.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(sign,currentTime,ownerId,hash)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EES: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while marking plugin as signed.")
  }
  return nil
}
