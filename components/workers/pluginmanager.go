package workers

type Plugins struct {
  Owner string
  Name string
  Hash string
  PType  int
  Decsription string
  Active bool
  CreatedAt string
  UpdatedAt string
}

const  (
  cleaner = iota + 1
  malwareScanner
  vulnerabilityScanner
)

func DBCreatePlugin(p Plugin)error{
  return nil
}

func ListPlugins(ownerId string, admin bool)([]Plugin,error){
  return nil,nil
}

func ViewPlugin(pid string)(*Plugin,error){
  return nil,nil
}

func DeactivatePlugin(pid,owner string,admin bool)error{
  return nil
}
