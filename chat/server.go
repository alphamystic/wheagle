package chat

type ChatServer struct{
  Rooms map[string]*Room
  commands chan ChatCommand
}

func NewChatServer()*ChatServer{
  return &server := {
    rooms: make(map[string]*Room),
    commands: make(chan ChatCommand),
  }
}


func (s *ChatServer) NewClient(conn net.Conn){
  log.Printf("[+] New Client has connected: %s",conn.RemoteAddr().String())
  c := &client{
    conn: conn,
    nick: "anonymous",
    ChatCommands: s.ChatCommands,
  }
  
}
