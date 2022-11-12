package edr

type AssetFile struct{
  AssetId string `json:"assetid"`
  Filename string `json:"filename"`
  FileId string `json:"fileid"`
  FileHash string `json:"filehash"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

// assetid 	filename 	fileid 	filehash 	created_at 	updated_at
func CreateAssetFile(af AssetFile)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`assetfiles` (assetid,filename,fileid,filehash,created_at,updated_at) VALUES(?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPCAF: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to create assetfile. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&a.AssetId,&a.Filename,&a.FileId,&a.FileHash,&a.CreatedAt,&a.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EECAF: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating asset file.")
  }
  return nil
}

func ListAssetFiles(assetId string)([]AssetFile,error){
  stmt := "SELECT * FROM `wheagle`.`assetfiles` WHERE (`assetid` = ? ) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,assetId)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELAF: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all api keys.")
  }
  defer rows.Close()
  var af []AssetFile
  for rows.Next(){
    var a AssetFile
    err = rows.Scan(&a.AssetId,&a.Filename,&a.FileId,&a.FileHash,&a.CreatedAt,&a.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESF: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing assetfiles.")
    }
    af = append(af,a)
  }
  return af,nil
}

func ViewAssetFiles() error{
  return nil
}

func DeleteAssetFile()error{
  return nil
}

func DeleteAssetFiles()error{
  return nil
}
