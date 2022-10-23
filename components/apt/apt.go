package apt

type Apt struct {
  Name string
  Code int
  Id string
  Description string
  Active true
  CreatedAt string
  UpdatedAt string
}


func CreateAPT(a Apt)error{
  return  nil
}

func ListAPT(active bool)([]Apt,error){
  return nil,nil
}

func DeactivateAPT(aptId string)error{
  return nil
}

func UpdateApt(a *Apt)error{
  return nil
}

func ListAPTs(active bool)([]Apt,error){
  return nil,nil
}
