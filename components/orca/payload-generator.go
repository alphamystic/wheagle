package orca

type PAYLOAD interface{
  RegisterPayloadGenerator() (pc *[]PayloadCommands,name string,hash string)
  GeneratePayload(agent *Agent,pc *[]PayloadCommands) ([]byte,error)
  RunPayload(name string,pc []*PayloadCommands) ([]byte,error)
  //packs a given payload in any way and returns the bytes to be save as a payload
  Packer(file []byte)(newFile []byte,error)
}

var commandParameter [][]string{}

type PayloadCommands struct{
  Commands *commandParameter
}

/* So what des a payload contain?
    1. Could be any give revshell connection
    2. Can be justa regular file like a ransomware file or a rootkit trojanbe.t.c
*/

/* what if I just want to return a connection
  Then use RunPayload() returning the connectionsin bytes or errror if it failed
*/
