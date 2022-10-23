package apt

type Activity struct{
  Name string
  AId string
  Description string
  Validated bool
  CreatedAt string
  UpdatedAt string
}

func CreateActivity(a Activity)error{
  return nil
}

func ViewActivity(aid string)(*Activity,error){
  return nil,nil
}

func ValidateActivity(aid string)error{
  return nil
}

func ListActivity(valid bool)([]Activity,error){
  return nil,nil
}
