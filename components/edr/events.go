package edr

type Event struct{
  EventId string
  ThreatLevel int
  UserId string
  Description string
  Comment string
  Handled bool  //status
  HandlerId string
  CreatedAt string
  UpdatedAt string
}

/*
  *** SYSMON EVENT LOGS ****
        WINDOWS
  ProcessCreation ID 1
  NetworkConnection ID 3
  DLLLoaded ID 7
  ProcessHallowing ID 8
  ProcessAccess ID 10
  FileCreation ID 11
  RegistryKey ID 12/13/14
  AlternateDataStreams ID 15
  DNSEvents ID 22
*/
