package payload

type PayloadGenerator struct{}

type F func() ([]byte,error)

type Payload struct{
  Name string
  Extension string
}

func (p *PayloadGenerator)Encoder(f F) ([]byte,error){
  return nil,nil
}

func New() orca.GeneratePayload {
  return new(PayloadGenerator)
}
