package edr

/*
  * Assets defines hardware types. This are created by the admin as they can't exactly create themselves
  * They can range from keyboards to mouse and flashdisks or extra drives.
*/

type Asset struct{
  AgentId string `json:"agentid"`
  Name string `json:"asset_name"`
  AId string  `json:"asset_id"`
  Description string  `json:"description"`
  Active bool `json:"active"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

//  	agentid 	name 	asset_id 	description 	active 	created_at 	updated_at
func CreateAsset(a Asset)error{
  return var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`assets` (agentid,name,asset_id,description,active,created_at,updated_at) VALUES(?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTCA: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create asset. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.AgentId,&a.Name,&a.Aid,&a.Description,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EXIAPIK: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating asset.")
  }
  return nil
}

func ListAssets(active bool)([]Asset,error){
  stmt := "SELECT * FROM `wheagle`.`assets` WHERE (`active` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,active)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EQA: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing assets.")
  }
  defer rows.Close()
  var asts []Asset
  for rows.Next(){
    var a Asset
    err = rows.Scan(&a.AgentId,&a.Name,&a.Aid,&a.Description,&a.Active,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESA: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing asset.")
    }
    asts = append(asts,a)
  }
  return asts,nil
}

func ViewAsset(assId,agentId string)(*Asset,error){
  var a Asset
  row := db.QueryRow("SELECT * FROM `wheagle`.`assets` WHERE (`agentid`	 = ? AND `asset_id` = ?);",minionId,agentId)
  err := row.Scan(&a.AgentId,&a.Name,&a.Aid,&a.Description,&a.Active,&a.CreatedAt,&a.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EVMA with id %s %s",agentId,err))
    utils.Logerror(e)
    if err == sql.ErrNoRows {
      utils.Danger(fmt.Errorf("Agent Id %s or asset Id %s doen't exist: ", agentId,assId))
      return errors.New("Requested asset agent doesn't exist")
    }
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing asset of %s",assId))
  }
  return &a,nil
}
