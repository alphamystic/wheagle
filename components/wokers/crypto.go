package wokers


import (
  "io"
  "os"
  "crypto/md5"
  "encoding/hex"
  "encoding/base64"
)


func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}


func ComputeMD5(filePath string) []byte {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return nil
	}
	return hash.Sum(result)
}


func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		NewDebugUpdate("Error: None-Valid Base64 String...")
		return ""
	}
	return string(data)
}

func HashStruct()string{
  var val string
  return val
}
