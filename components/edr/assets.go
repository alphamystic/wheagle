package edr

/*
  * Assets defines hardware types. This are created by the admin as they can't exactly create themselves
  * They can range from keyboards to mouse and flashdisks or extra drives.
*/

type Asset struct{
  Name string `json:"asset_name"`
  AId string  `json:"asset_id"`
  Description string  `json:"description"`
  Active bool `json:"active"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
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
