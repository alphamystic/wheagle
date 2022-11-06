package workers

type Task struct{
  OwnerId string
  TaskId string
  Name string
  Type bool//cmd or internal cmd
  Description string
  Handled bool
  CreatedAt string
  UpdatedAt string
}

func CreateTask()error{
  return nil
}

func ListTask(complete bool)([]Task,error){
  return nil,nil
}

func ListMyTask(userId,jobId string)([]Task,error){
  return nil,nil
}

func MarkAsHandled(taskId,ownerId string,isAdmin bool)error{
  //remeber to check if user is admin to remove
  return nil
}

func ViewTask(taskId string)(*Task,error){
  return nil,nil
}
