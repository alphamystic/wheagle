package apt

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Apt struct {
  Name string
  Code int
  Id string
  Description string
  Active true
  CreatedAt string
  UpdatedAt string
}

// 	aptname 	code 	id 	description 	active 	created_at 	updated_at

func CreateAPT(a Apt)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`apt` (aptname,code,id,description,active,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTIApt: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to insert apt. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.Code,&a.Id,&a.Description,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECApt: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating apt.")
  }
  return nil
}

func ListAPT(active bool)([]Apt,error){
  stmt := "SELECT * FROM `wheagle`.`apt` WHERE (`active` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELAApt: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing apt.")
  }
  defer rows.Close()
  var apts []Apt
  for rows.Next(){
    var a Apt
    err = rows.Scan(&a.Code,&a.Id,&a.Description,&a.Active,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESApt: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing apt.")
    }
    apts = append(apts,a)
  }
  return apts,nil
}

func ViewApt(aid string)(*Apt,error){
  var a Apt
  row := db.QueryRow("SELECT * FROM `wheagle`.`apt` WHERE id = ?;",aid)
  err := row.Scan(&a.Name,&a.Aid,&a.CreatorId,&a.Description,&a.Validated,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVApt with id %s %s",aid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing apt with id of %s",aid))
  }
  return &a,nil
}

func MarkAPTAsInactive(aptId string)error{
  upStmt := "UPDATE `wheagle`.`apt` SET `active` = ? AND `updated_at` ? WHERE (`id` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPAptIn: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while marking APT an inactive.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(true,aptId,currentTime)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EMNH: activity id: %s  %s",aid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing apt deactivate.")
  }
  return nil
}

func UpdateApt(a *Apt)error{
  upStmt := "UPDATE `wheagle`.`apt` SET (`aptname` = ? AND `code` = ? AND `id`= ? AND `decription` = ? AND `updated_at` ?) WHERE (`id` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTUApt: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to update APT.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(&a.Name,&a.Code,&a.Id,&a.Description,currentTime)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EMNH: activity id: %s  %s",aid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing update apt.")
  }
  return nil
}
