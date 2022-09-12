package components

type Bot struct{
  BotId string
  Name string
  Ip string
  OwnerId string
  CreatedAt string
  UpdatedAt string
}

func DBCreateBot(b Bot)error{
  return nil
}

func DBViewBot(botId string){}
