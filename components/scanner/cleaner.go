package scanner

/*
  Enumerates the host looking for iocs and cleaning them out
*/

type Cleaner interface{
  Clean(ioc []*IOC) *Cleaned
}

type Cleaned struct{
  Cleaned false//if no ioc found return false
  Details []string
}
