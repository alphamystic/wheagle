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

func CreateThreat(t Threat)error{
  return nil
}

//func ViewThreat()
