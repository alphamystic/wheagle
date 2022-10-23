package apt

type Virus struct{
  VirusId string
  Hash string
  VirusType string //worm,rootkit,trojan
  CommunicationMode string //p2p or CnC
  OsType string
  Description string
  CreatedAt string
  UpdatedAt string
}

func CreateVirus(v Virus)error{
  return nil
}

func ListVirus()([]Virus,error){return nil,nil}

func ListOsVirus(ostype string)([]Virus,error){
  return nil,nil
}
