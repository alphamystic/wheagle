package hb


type Sessions struct {
  sessions *[]stream.Session
}

// you can also initialize nthis at the root of your server for payload generation
func (ss ) InitializeSessions()[]*stream.Session{
  sessions := make([]*stream.Session)
  return &sessions
}

func (ss *Sessions) AddNewSession(s *stream.Session) *Sessions{
  sessions := append(&ss,s)
  return &session
}

func (ss *Sessions) CreateSession(s *stream.Session)(*stream.Session){
  session := *stream.Session{
    SessionId: utils.TrueRand(15),
  	ClientId: utils.TrueRand(15),
  	Motherships: []*stream.Mothership,
  	InteractMode: false,
  	Transport: tcp,
  	ProxyUrls : []*stream.ProxyUrl,
  	HostName:"",
  	Arch:"",
  	ActiveC2: "myIp",
  	IsAlive: false,
  	PID: int32(300),
  }
  return &session
}

func (ss *Session) UpdateSession(s *stream.Session) error{
  s1 := ss.GetSession(s.SessionId)
  if s1 == ""{
    return errors.New("Error, Session doesn't exist")
  }
  var s2 *stream.Session
  //set s2
  //remove s1
  var sns Sessions
  sns = append(&ss,s2)
  return nil
}

func (ss *stream.Session) RemoveSession(sessionId string)*Sessions{
  return nil
}
