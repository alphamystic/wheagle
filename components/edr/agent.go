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

func Createagent(a Aget)error{
  return nil
}

//Guess it should have a universal ID as the owner to allow many SOC operators on it
func ViewAgent(aid,ownerId string)(*Agent,error){
  return nil,nil
}

func ListAgents(ownerId string)([]Agent,error){
  return nil,nil
}

func DeleteAgent(isAdmin bool)error{
  return nil
}
