package edr

/*
  Listing sysmon events documentations
  https://rootdse.org/posts/understanding-sysmon-events/
  https://systemweakness.com/list-of-sysmon-event-ids-for-threat-hunting-4250b47cd567
  https://learn.microsoft.com/en-us/sysinternals/downloads/sysmon
*/

/*
  REDING WINDOWS EVENTS IN POWERSHELL
  https://www.avecseclabs.com/bif/2019/07/15/using-powershell-to-read-sysmon-events/
    Get-WinEvent @{logname="Microsoft-Windows-Sysmon/Operational"} | Select-Object -Property TimeCreated,ID,Message -First 5 | Format-List
  https://devblogs.microsoft.com/scripting/use-powershell-to-query-all-event-logs-for-recent-events/
*/


type Event interface{
  ProcesCreate() *Details //eventid 1
  ProcessChanged() *Details //eventid 2
  NetworkConnection() *Details //event id 3
  ServiceStateChange() *Details //event id 4
  ProcessTerminated() *Details //event id 5
  DriverLoaded() *Details //event id 6
  ImageLoad() *Details //event id  7
  CreateRemoteThread() *Details //event id 8
  RawAccessRead() *Details //event id 9
  ProcessAccess() *Details //event id 10
  FileCreate() *Details //event id 11
  RegistryEventOCD() *Details //event id 12 Object Create Delete
  RegistryEventVS() *Details //event id 13 Value Set
  RegistryEventKVR() *Details //event id 14 (Key and Value Rename)
  FileCreateStreamHash() *Details //event id 15
  ServiceConfigurationChange() *Details //event id 16
  PipeEventPCr() *Details //event id 17 (Pipe Created)
  PipeEventPCn() *Details //event id 18 (Pipe Event Connected)
  WmiEventFAD() *Details //event id 19  WmiEventFilter activity Detected
  WmiEventCAD() *Details //event id 20 WminEvent COnsumer activity detected
  WmiEventCFAD() *Details //event id 21 wmiEvent  consumer to filter activity detected
  DnsEvent() *Details //event id 22
  FileDelete() *Details //event id 23
  ClipBoardChange() *Details //event id 24
  ProcessTampering() *Details //event id 25 Process image change
  Event255() *Details //event id 26 Sysmon error
  FileBlockExecutable() *Details // event id 27 (Still on test mode) https://www.huntandhackett.com/blog/bypassing-sysmon
}

type Detail struct{
  EventId string
  ThreatLevel int
  UserId string
  Description string
  EventType int
  Handled bool  //status
  eventData *EventData
  HandlerId string
  CreatedAt string
  UpdatedAt string
}

type EventData struct{
  EData ...interface{}
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

func (e *Event) ProcesCreate()error{
  return nil
}
