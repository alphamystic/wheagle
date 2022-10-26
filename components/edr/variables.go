package edr

/*
  * Contains variables for components functioning
*/

import(
  "time"
  "database/sql"
)

var db *sql.DB
var now = time.Now()
var currentTime = now.Format("2006-01-02 15:04:05")

// Use this in the db to identify the types of threats
const (
  processCreate = iota + 1
  processChanged
  networkConnection
  servicesStatusChanged
  processTerminated
  driverLoaded
  imageLoad
  createRemoteThread
  rawAccessRead
  processAccess
  fileCreate
  registryEventOCD
  registryEventVS
  registryEventKVR
  fileCreateStreamHash
  serviceConfiguratonChange
  pipeEventCpr
  pipeEventPCn
  wmiEventFAD
  wmiEventCAD
  wmiEventCFAD
  dnsEvent
  fileDelete
  clipBoardChange
  processTampering
  event255
  fileBlockExecutable
)
