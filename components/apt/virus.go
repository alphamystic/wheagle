package apt

type Virus struct{
  VirusId string
  Hash string
  VirusType string //worm,rootkit,trojan
  CommunicationMode string //p2p or CnC
  OsType string
  Description string
}
