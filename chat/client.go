package chat

type ChatClient struct{
  conn net.Conn
  nick string
  ClientRoom *Room
  ChatCommands chan<- ChatCommand
}


func (c *ChatClient) Run(){
  for cmd := range s.ChatCommands {
    switch cmd.Id{
    case CMD_NICK:
      s.Nick()
    case CMD_NICK:
      s.Nick()
    case CMD_NICK:
      s.Nick()
    case CMD_NICK:
      s.Nick()
    case CMD_NICK:
      s.Nick()
    }
  }
}

func (c *ChatClient) readInput(){
  for{
    msg,err := bufio.NewReader(c.Conn).ReadString('\n')
    if err != nil{
      return
    }
    msg = strings.Trim(msg,"\r\n")
    args := strings.Split(msg," ")
    cmd := strings.TrimSpace(args[0])
    switch cmd{
    case "/nick":
      c.ChatCommands <- ChatCommand{
        Id: CMD_NICK,
        chatClient: c,
        Args: args,
      }
    case "/join":
      c.ChatCommands <- ChatCommand{
        Id: CMD_JOIN,
        chatClient: c,
        Args: args,
      }
    case "/rooms":
      c.ChatCommands <- ChatCommand{
        Id: CMD_ROOMS,
        chatClient: c,
        Args: args,
      }
    case "/msg":
      c.ChatCommands <- ChatCommand{
        Id: CMD_MSG,
        chatClient: c,
        Args: args,
      }
    case "/quit":
      c.ChatCommands <- ChatCommand{
        Id: CMD_QUIT,
        chatClient: c,
        Args: args,
      }
    default:
      c.ClientError(fmt.Errorf("Unknown command: %s",cmd))
    }
  }
}


func (c *ChatClient) ClientError(err error){
  c.conn.Write([]byte("ERR: "+ err.Error() + "\n"))
}

func (c *ChatClient) ClientMessage(msg string){
  c.conn.Write([]byte("[Wheagle]: " + msg + "\n"))
}
