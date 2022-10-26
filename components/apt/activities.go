package apt

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Activity struct{
  Name string `json:"activity_name"`
  AId string  `json:"activity_id"`
  CreatorId string `json:"creator_id"`
  Description string  `json:"activity_description"`
  Validated bool `json:"validated"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  ```json:"updatedat"`
}

//  	name 	activity_id 	creatorid 	description 	validated 	created_at 	updated_at
func CreateActivity(a Activity)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`threatactivity` (name,activity_id,creatorid,description,validated,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTIA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to insert activity. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.Name,&a.Aid,&a.CreatorId,&a.Description,&a.Validated,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EICA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating activity.")
  }
  return nil
}

func ViewActivity(aid string)(*Activity,error){
  var a Activity
  row := db.QueryRow("SELECT * FROM `wheagle`.`threatactivity` WHERE activity_id	 = ?;",aid)
  err := row.Scan(&a.Name,&a.Aid,&a.CreatorId,&a.Description,&a.Validated,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVA with id %s %s",aid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing activity with id of %s",aid))
  }
  return &a,nil
}

func ValidateActivity(aid string)error{
  upStmt := "UPDATE `wheagle`.`threatactivity` SET `validated` = ? WHERE (`activity_id` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPUA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while validating activity.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(true,aid)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EMNH: activity id: %s  %s",aid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing validate activity.")
  }
  return nil
}

func DeleteActivity(aid string)error{
  del,err := db.Prepare("DELETE FROM `wheagle`.`threatactivity` WHERE (`activity_id` = ?);")
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPDA with id: %s : %s",aid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while deleting activity.")
  }
  defer del.Close()
  var res sql.Result
  res,err = del.Exec(aid)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EEDA with id: %s : %s",aid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while deleting activity.")
  }
  return nil
}

func ListActivity(valid bool)([]Activity,error){
  stmt := "SELECT * FROM `wheagle`.`threatactivity` WHERE (`validated` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,valid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELVIA: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all handled activities.")
  }
  defer rows.Close()
  var actvts []Activity
  for rows.Next(){
    var a Activity
    err = rows.Scan(&a.Name,&a.Aid,&a.CreatorId,&a.Description,&a.Validated,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESFTA: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing activity.")
    }
    actvts = append(actvts,a)
  }
  return actvts,nil
}

func ListActivityByCreator(userId string,valid bool)([]Activity,error){
  stmt := "SELECT * FROM `wheagle`.`threatactivity` WHERE (`creatorid` = ? AND `validated` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,userId,valid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELABC: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing activity by creator.")
  }
  defer rows.Close()
  var activity []Activity
  for rows.Next(){
    var a Activity
    err = rows.Scan(&a.Name,&a.Aid,&a.CreatorId,&a.Description,&a.Validated,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESABC: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing activity by creator.")
    }
    actvts = append(actvts,a)
  }
  return actvts,nil
}
