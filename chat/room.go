package chat

type Room struct{
  Name string
  Members map[net.Addr]*ChatClient
}
