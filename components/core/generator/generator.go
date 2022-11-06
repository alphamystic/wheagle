package generator

type Generator interface{
  ShellCodeGenerate() error
  MacroGenerate() error
  BinaryGenerate() error
}

// An implant can be a mere agent or a fully featured minion
type Implant struct{
  Session *Session
  Name string
}

//https://github.com/EddieIvan01/gld
func Generate(gen Generator) error{
  return nil
}
