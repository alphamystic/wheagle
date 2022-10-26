package orca

// this is more of a detection interface
type EDR interface{
  Static(file []byte) (*StaticData,error)
  Dynamic(file []byte) (*StaticData,error)
  Behavioral(file []byte) (*StaticData,error)
}


/*
  The runner for this will read a file and dump it incide the returned interface
*/

type StaticData struct{
  Good bool
  CPatterns []CommonPatterns
}

type CommonPatterns struct{
  Title string
  Details string
}
