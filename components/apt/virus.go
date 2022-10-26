package apt

import (
  "fmt"
  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/alphamystic/wheagle/utils"
)

type Virus struct{
  AptId string
  VirusId string
  Hash string
  VirusType string //worm,rootkit,trojan
  FileType string //lnk,zip.iso,
  CommunicationMode string //p2p or CnC
  OsType string
  Description string
  CreatedAt string
  UpdatedAt string
}

//aptid 	virusid 	hash 	virustype 	filetype 	communicationmode 	ostype 	description 	created_at 	updated_at
func CreateVirus(v Virus)error{
  var ins *sql.Stmt
  ins,err := db.Prepare("INSERT INTO `wheagle`.`virus` (aptid,virusid,hash,virustype,filetype,communicationmode,ostype,description,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPTIV: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while preparing to insert virus. Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(&v.AptId,&v.VirusId,&v.Hash,&v.VirusType,&v.FileType,&v.FileType,&v.CommunicationMode,&v.OsType,&v.Description,&v.CreatedAt,&v.UpdatedAt)
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EICV: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating virus.")
  }
  return nil
}

func ListVirus()([]Virus,error){
  stmt := "SELECT * FROM `wheagle`.`virus`  ORDER BY updated_at DESC;"
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
    err = rows.Scan(&v.AptId,&v.VirusId,&v.Hash,&v.VirusType,&v.FileType,&v.FileType,&v.CommunicationMode,&v.OsType,&v.Description,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESLV: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing virus.")
    }
    apts = append(apts,a)
  }
  return apts,nil
}

func ListOsVirus(ostype string)([]Virus,error){
  stmt := "SELECT * FROM `wheagle`.`apt` WHERE (`ostype` = ?) ORDER BY updated_at DESC;"
  rows,err := db.Query(stmt,ostype)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("ELAApt: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing virus by type.")
  }
  defer rows.Close()
  var apts []Apt
  for rows.Next(){
    var a Apt
    err = rows.Scan(&v.AptId,&v.VirusId,&v.Hash,&v.VirusType,&v.FileType,&v.FileType,&v.CommunicationMode,&v.OsType,&v.Description,&v.CreatedAt,&v.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("ESLoV: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing virus.")
    }
    apts = append(apts,a)
  }
  return apts,nil
}
