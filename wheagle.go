package main

import(
  "fmt"
	_ "image/jpeg"
	_ "image/png"
  "github.com/alphamystic/wheagle/cmd"
  "github.com/qeesung/image2ascii/convert"
)


func main(){
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	// Create the image converter
	converter := convert.NewImageConverter()
	fmt.Print(converter.ImageFile2ASCIIString("./wheagle.png", &convertOptions))
  fmt.Println("[+] Starting Wheagle")
  cmd.Execute()
}



//https://demo.dashboardpack.com/hospital-html/index.html
