package orca

type BACKDOOR interface{
  //Reads a file backdoors it and returns a newer copy of the file
  //The reasons for implementing each is to ensure the file submitted
  // is of the requested type (Will probably change it to return just byte code)
  BDExe(legitfile,targetfile string) (*Backdoored,error)
  BDElf(legitfile,targetfile string) (*Backdoored,error)
  BDApk(legitfile,targetfile string) (*Backdoored,error)
  BDDll(legitfile,targetfile string) (*Backdoored,error)
  BDSo(legitfile,targetfile string) (*Backdoored,error)
  BDDeb(legitfile,targetfile string) (*Backdoored,error)
  BDImg(legitfile,targetfile string) (*Backdoored,error)
  BDMp4(legitfile,targetfile string) (*Backdoored,error)
  BDPdf(legitfile,targetfile string) (*Backdoored,error)
  BDZip(legitfile,targetfile string) (*Backdoored,error)
  BDIso(legitfile,targetfile string) (*Backdoored,error)
}

type Backdoored struct{
  legitfile,targetfile string
  Architecture string
  Data []byte
}
