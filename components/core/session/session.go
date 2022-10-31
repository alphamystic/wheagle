package session

type Session struct {}

var Sessions []Session

func NewSession()*Session{}

func (s *Session) AddNewSession(s Session)error{
  return nil
}

func (s *Session) GetSession(sessionId string) (*Session,error){
  return nil,nil
}

func (s *Session) RemoveSession(sessionId string)(*Sessions,error){
  return nil,nil
}

func ListSessions()(*Sessions,error){
  return nil,nil
}
