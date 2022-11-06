package session

type Session struct {}

type SessionManager struct{
  Sessions []*Session
}

var Sessions []Session

func init(){
  sessions := make([]Session)
  sm := SessionManager{sessions}
}

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

func (s *Session) ViewSession()(*Sessions,error){
  return nil,nil
}

func (sm *SessionManager) ListSessions()*Sessions{
  return &sessions
}

func (sm *SessionManager) ListPermanentSessions()(*Sessions){
  return nil
}
