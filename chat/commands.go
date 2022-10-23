package chat

type commandId int

const (
  CMD_NICK commandId = iota
  CMD_JOIN
  CMD_ROOMS
  CMD_MSG
  CMD_QUIT
)
type ChatCommand struct{
  Id commandId
  chatClient (ChatClient)
  Args []string
}
