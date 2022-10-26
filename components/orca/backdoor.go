package orca

type BACKDOOR interface{
  //Reads a file backdoors it and returns a newer copy of the file
  //The reasons for implementing each is to ensure the file submitted
  // is of the requested type (Will probably change it to return just byte code)
  BDExe(filename string) (*Backdoored,error)
  BDElf(filename string) (*Backdoored,error)
  BDApk(filename string) (*Backdoored,error)
  BDDll(filename string) (*Backdoored,error)
  BDSo(filename string) (*Backdoored,error)
  BDDeb(filename string) (*Backdoored,error)
  BDImg(filename string) (*Backdoored,error)
  BDMp4(filename string) (*Backdoored,error)
  BDPdf(filename string) (*Backdoored,error)
  BDZip(filename string) (*Backdoored,error)
  BDIso(filename string) (*Backdoored,error)
}

type Backdoored struct{
  Filename string
  Architecture string
  Data []byte
}
