package edr

import (
  "net"
)

type Threat struct{
  Name string
  Ip net.Ip
  VirusId string
  Details string
  Active bool
  CreatedAt string
  UpdatedAt string
}

// name 	threat_ip 	virusid 	details 	active 	created_at 	updated_at
func CreateThreat(t Threat)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`threats` (name,threat_ip,virusid,details,active,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTCT: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create threat. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&t.Name,&t.Ip,&t.VirusId,&t.Details,&t.Active,&t.CreatedAt,&t.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECT: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating threat.")
  }
  return nil
}

func ViewThreat(vid string)(*Threat,error){
  var t Threat
  row := db.QueryRow("SELECT * FROM `wheagle`.`threat` WHERE (`virusid`	 = ?);",vid)
  err := row.Scan(&t.Name,&t.Ip,&t.VirusId,&t.Details,&t.Active,&t.CreatedAt,&t.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EQVT with id %s",vid,err))
    utils.Logerror(e)
    if err == sql.ErrNoRows {
      utils.Danger(fmt.Errorf("Threat Id %s doen't exist: ",vid))
      return errors.New("Requested asset agent doesn't exist")
    }
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing threat id of %s",vid))
  }
  return &t,nil
}

func ListThreats(active bool) ([]Threat,error){
  stmt := "SELECT * FROM `wheagle`.`threats` WHERE (`active` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELAPIK: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing threats.")
  }
  defer rows.Close()
  var ths []Threats
  for rows.Next(){
    var t Theat
    err = rows.Scan(&t.Name,&t.Ip,&t.VirusId,&t.Details,&t.Active,&t.CreatedAt,&t.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESFLT: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing threats.")
    }
    ths = append(ths,t)
  }
  return ths,nil
}


func MarkThreatAsIActive(vid string,active bool) error{
  upStmt := "UPDATE `wheagle`.`threats` SET (`active` = ? AND `updated_at` = ?) WHERE (`virusid` = ?);";
  stmt,err := db.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTMTAA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to mark threat as active inactive.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(active,vid,currentTime)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EUAPIK: activity id: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while executing update apikey.")
  }
  return nil
}
