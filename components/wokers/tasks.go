package workers

type Job struct{
  OwnerId string
  JobId string
  JobName string
  JobType bool//cmd or internal cmd
  Description string
  Handled bool
  CreatedAt string
  UpdatedAt string
}

func CreateJob()error{
  return nil
}

func ListJobs(complete bool)([]Job,error){
  return nil,nil
}

func ListMyJobs(userId,jobId string)([]Job,error){
  return nil,nil
}

func MarkAsHandled(jobId,ownerId string,isAdmin bool)error{
  //remeber to check if user is admin to remove
  return nil
}

func ViewJob(jobId string)(*Job,error){
  return nil,nil
}
