package data

import (
  "net"
)

/* A minion is rat of your choosing, from rootkit to a normal agent
  This is what wheagle uses it as an agent
*/
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
  Ip net.Ip  `json:"ip"`
  OwnerId string  `json:'ownerid'`
  LastSeen string `json:"lastseen"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}
/*
string MinionId = 1;
string MSession = 2;
string HostName = 3;
string Username = 4;
string Userid = 5;
string GroupId = 6;
string Homedir = 7;
string MinionType = 8;
string OsType = 9;
string Description = 10;
bool Installed = 11;
string MothershipId = 12;
string MinionIp = 13;
string OwnerId = 14;
string Lastseen = 15;
string PID = 16;
bool Persistance = 17;
string PersistanceMode = 18;
*/


func DBCreateMultipleImplants(minions []Minion)error{
  for _,minion := range minions{
    err := DBCreateMinion(minion)
    if err != nil{
      return err
    }
  }
  return nil
}

func DBCreateMinion(b Bot)error{
  return nil
}

func DBViewMinion(minionId string){}


func DbListMinions(ownerId string)([]Minion,error){
  return nil,nil
}
