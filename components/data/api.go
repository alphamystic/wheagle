package data

type Api struct{
  ApiKey string `json:"apikey"`
  OwnerId string  `json:"ownerid"`
  Active string `json:"active"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

func CreateApiKey(a Api)error{
  return nil
}

func ListApiKeys(active bool)([]Api,error){
  return nil,nil
}

/*if !true{
  //Do the Thing
}
*/
func ValidateApiKey(key,ownerId string)bool{
  return true
}

func UpdateKey(ownerId string)error{
  return nil
}

func DeactivateKey(ownerId string)error{
  return nil
}


func ViewApiKey(ownerId string)(*Api,error){
  return nil,nil
}
