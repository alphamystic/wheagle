package workers

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Task struct{
  OwnerId string
  TaskId string
  Name string
  Type bool//cmd or internal cmd
  Description string
  Handled bool
  CreatedAt string
  UpdatedAt string
}

// 	ownerid 	taskid 	name 	type 	description 	handled 	created_at 	updated_at
func CreateTask(t Task)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`task` (ownerid,taskid,name,type,description,handled,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTCT: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create task. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&t.OwnerId,&t.TaskId,&t.Name,&t.Type,&t.Description,&t.Handled,&t.CreatedAt,&t.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECT: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating task.")
  }
  return nil
}

func ListTask(ownerid string,complete bool)([]Task,error){
  stmt := "SELECT * FROM `wheagle`.`task` WHERE (`handled` = ? AND `ownerid` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,complete,ownerid)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOT: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing owner tasks.")
  }
  defer rows.Close()
  var tasks []Task
  for rows.Next(){
    var t Task
    err = rows.Scan(&t.OwnerId,&t.TaskId,&t.Name,&t.Type,&t.Description,&t.Handled,&t.CreatedAt,&t.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESFOT: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing owner task.")
    }
    tasks = append(tasks,t)
  }
  return tasks,nil
}

func ListTaskByType(ownerid,cmdtype,jobId string,complete bool)([]Task,error){
  stmt := "SELECT * FROM `wheagle`.`task` WHERE (`handled` = ? AND `ownerid` = ? AND type = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,complete,ownerid,cmdtype)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELOT: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing owner tasks by command.")
  }
  defer rows.Close()
  var tasks []Task
  for rows.Next(){
    var t Task
    err = rows.Scan(&t.OwnerId,&t.TaskId,&t.Name,&t.Type,&t.Description,&t.Handled,&t.CreatedAt,&t.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESFOT: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing owner tasks by command.")
    }
    tasks = append(tasks,t)
  }
  return tasks,nil
}

func MarkAsHandled(taskId,ownerId string,isAdmin bool)error{
  //remeber to check if user is admin to remove
  // Now remove this to root at the handler
  if !data.CheckIfAdmin(ownerId){
    utils.Danger(errors.New("Non admin tried marking task as handled"))
    return fmt.Errorf("Is not admin")
  }
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

func ViewTask(ownerId,taskId string)(*Task,error){
  var t Task
  row := db.QueryRow("SELECT * FROM `wheagle`.`task` WHERE (`ownerid` = ? AND `taskid` = ?);",ownerId,taskId)
  err := row.Scan(&t.OwnerId,&t.TaskId,&t.Name,&t.Type,&t.Description,&t.Handled,&t.CreatedAt,&t.UpdatedAt)
  if err != nil{
    if err == sql.ErrNoRows {
      utils.Danger(fmt.Errorf("Attempted access to task id of %s by user %s"+ taskId,ownerId))
      return errors.New("Requested taskid doesn't exist")
    }
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVT with id %s %s",taskId,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("EVT with id %s",taskId))
  }
  return &t,nil
}
