package workers

import (
  "os"
  "fmt"
  "image"
  "image/png"

  "github.com/kbinani/screenshot"
)

type MyImages struct{
  ImageOne string
  ImageOneName string
  ImageTwo string
  ImageTwoName string
}
func Save(img *image.RGBA,filePath string){
  file,err := os.Create(filePath)
  if err != nil{
    fmt.Println("[-]    Unable to create image file: ",err)
  }
  defer file.Close()
  png.Encode(file,img)
}

func TakeScreenShot()(error,bool){
  var rp = "No Active display found"
  n := screenshot.NumActiveDisplays()
  if n <= 0{
    return fmt.Errorf(rp),false
  }
  var all image.Rectangle = image.Rect(0,0,0,0)
  for i := 0; i < n; i++{
    bounds := screenshot.GetDisplayBounds(i)
    bounds.Union(all)
    img, err := screenshot.CaptureRect(bounds)
    if err != nil {
      return err,false
    }
    fileName := fmt.Sprintf("%d_%dx%d.png",i,bounds.Dx(),bounds.Dy())
    Save(img,fileName)
    fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
  }
  //capture the whole desktop region into an image
  fmt.Printf("%v\n",all)
  img,err := screenshot.Capture(all.Min.X,all.Min.Y,all.Dx(),all.Dy())
  if err != nil{
    return err,false
  }
  Save(img,"all.png")
  return fmt.Errorf(rp),true
}
