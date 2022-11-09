package generator

type Generate interface{
  Generator() error
}

// An implant can be a mere agent or a fully featured minion
type Implant struct{
  Session *Session
  Name string
}

type Macro struct{}
func (m *Macro) Generator() error{
  //https://gist.github.com/mgeeky/9dee0ac86c65cdd9cb5a2f64cef51991
  return nil
}
type Binary struct{}
func (b *Binary) Generator() error{
  return nil
}
type ShellCode struct{}
func (s *ShellCode) Generator() error{
  return nil
}

//https://github.com/EddieIvan01/gld
func Generate(gen Generator) error{
  return nil
}
