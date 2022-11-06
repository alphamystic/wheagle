package orca

type BACKDOOR interface{
  //Reads a file backdoors it and returns a newer copy of the file
  //The reasons for implementing each is to ensure the file submitted
  // is of the requested type (Will probably change it to return an all for the inbuilt loader)
  BDExe(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDElf(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDApk(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDDll(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDSo(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDDeb(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDImg(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDMp4(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDPdf(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDZip(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  BDIso(legitfile,targetfile,payloadfile string) (*Backdoored,error)
  Default(filetype,legitfile,targetfile,payloadfile string) (*Backdoored,error)
}

type Backdoored struct{
  legitfile,targetfile string
  Architecture string
  Data []byte
}
