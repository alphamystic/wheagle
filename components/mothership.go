package components

import (
  "net"
)

type Mothership struct{
  OwnerId string  `json:"ownerid"`
  Name string `json:"name"`
  MId string  `json:"mothershipid"`
  IpAddress net.IP  //`json:"mothershipaddres"`
  Port int `json:"port"`
  Description string  `json:"description"`
  Active bool `json:"active"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}


func CreateMothership(m Mothership)error{
  return nil
}

func ListMotherShips(active bool)([]Mothership,error){
  return nil,nil
}

func DeactivateMothership(msid string)error{
  return nil
}

func ViewMothership(msid string)(*Mothership,error){
  return nil,nil
}
