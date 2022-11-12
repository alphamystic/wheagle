package edr

import (
  "net"
)

type Agent struct{
  MinionId string
  AgentId string
  Name string
  Ip net.Ip
  OwnerId string
  CreatedAt string
  UpdatedAt string
}

// 	minionid 	agentid 	name 	agent_ip 	ownerid 	created_at 	updated_at
func Createagent(a Aget)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`agent` (minionid,agentid,name,agent_ip,ownerid,created_at,updated_at) VALUES(?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPCA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create agent. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.MinionId,&a.AgentId,&a.Name,&a.Ip,&a.OwnerId,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating agent.")
  }
  return nil
}

//Guess it should have a universal ID as the owner to allow many SOC operators on it
func ViewAgent(minionId,agentId string)(*Agent,error){
  var a Agent
  row := db.QueryRow("SELECT * FROM `wheagle`.`agent` WHERE (`minionid`	 = ? AND `agentid` = ?);",minionId,agentId)
  err := row.Scan(&a.MinionId,&a.AgentId,&a.Name,&a.Ip,&a.OwnerId,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVMA with id %s %s",agentId,err))
    utils.Logerror(e)
    if err == sql.ErrNoRows {
      utils.Danger(fmt.Errorf("Agent Id %s or Minion Id %s doen't exist: ", agentId,minionId))
      return errors.New("Requested agent doesn't exist")
    }
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing agent of %s",agentId))
  }
  return &a,nil
}

func ListAgents(minionId string)([]Agent,error){
  stmt := "SELECT * FROM `wheagle`.`agent` WHERE (`minionid` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,minionId)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTLA: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing agents.")
  }
  defer rows.Close()
  var agnts []Agents
  for rows.Next(){
    var a Agent
    err = rows.Scan(&a.MinionId,&a.AgentId,&a.Name,&a.Ip,&a.OwnerId,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESLA: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing agents.")
    }
    agnts = append(agnts,a)
  }
  return agnts,nil
}

func DeleteAgent(isAdmin bool)error{
  return nil
}
