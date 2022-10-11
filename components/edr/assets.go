package edr

/*
  * Assets defines hardware types. This are created by the admin as they can't exactly create themselves
  * They can range from keyboards to mouse and flashdisks or extra drives.
*/

type Asset struct{
  Name string
  AId string
  Description string
  Active string
  CreatedAt string
  UpdatedAt string
}

func CreateAsset(a Asset)error{
  return string
}

func ListAssets(active bool)([]Asset,error){
  return nil,nil
}

func ViewAsset(assId string)(*Asset,error){
  return nil,nil
}
