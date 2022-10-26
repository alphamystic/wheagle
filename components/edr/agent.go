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
