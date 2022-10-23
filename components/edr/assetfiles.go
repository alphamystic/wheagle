package edr

type AssetFile struct{
  AssetId string `json:"assetid"`
  Filename string `json:"filename"`
  FileId string `json:"fileid"`
  FileHash string `json:"filehash"`
  CreatedAt string  `json:"createdat"`
  UpdatedAt string  `json:"updatedat"`
}

func CreateAssetFile(af AssetFile)error{
  return nil
}

func ListAssetFiles(assetId string)([]AssetFile,error){
  return nil,nil
}
