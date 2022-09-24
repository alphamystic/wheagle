package components

type Minion struct{
  MinionId string `json:"minionid"`
  Name string `json:"name"`
  UName string `json:"unsme"`
  UserId string  `json:"uid"`
  GroupId string  `json:"groupid"`
  HomeDir string `json:"homedir"`
  MinionType string `json:"miniontype"` //bot or agent
  Os  string  `json:"ostype"`
  Description string `json:"description"`
  Installed bool `json:"installed"`
  MotherShipId string `json:"mothershipid"`
  Ip string   `json:"ip"`
  OwnerId string  `json:'ownerid'`
  LastSeen string `json:"lastseen"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

func DBCreateMinion(b Bot)error{
  return nil
}

func DBViewMinion(minionId string){}


func DbListMinions(ownerId string)([]Minion,error){
  return nil,nil
}
