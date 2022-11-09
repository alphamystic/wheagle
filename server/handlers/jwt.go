package handlers
/*
import (
  "fmt"
  "time"
  "net/http"
  "github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("USE TO TRUE RSND TO GENERATE TRUE RANDOM NUMBERS")

func CreateJWT()(string,error){
  token := jwt.New(jwt.SigningMethodHS256)
  claims := token.Claims.(jwt.MapClaims)
  claims["exp"] = time.Now().Add(time.Hour).Unix()
  tokenStr,err := tokem.SignedString(SECRET)
  if err != nil{
    return "",err
  }
  return tokenStr,nil
}

func ValidateJWT(next func(res http.ResponseWriter, req *http.Request){
  return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request){
    if req.Header["Token"] != nil{
      token,err := jwt.Parse(req.Header["Token"][0],func(t *jwt.Token)(interface{},error){
        _,ok := token.Method.(*jwt.SigningMethodHMAC)
        if !ok {
          res.WriteHeader(http.StatusUnauthorized)
          res.Write([]byte("Not Authorised"))
        }
        return SECRET,nil
      })
      if err != nil {
        res.WriteHeader(http.StatusUnauthorized)
        res.Write([]byte("Not Authorized")
      }
      if token.Valid{
        next(res,req)
      }
    } else {
      res.WriteHeader(http.StatusUnauthorized)
      res.Write([]byte("Not Authorized: "+ err.Error()))
    }
  })
})

var api_key = "098976trfcv"
func GetJWT(res http.ResponseWriter, req *http.Request){
  if req.Header["Access"] != nil{
    if req.Header["Access"][0] == api_key{
      return
    } else {
      token,err := CreateJWT()
      if err != nil{fmt.Println("[-] Error creating token: ",err);return}
      fmt.Fprint(res,token)
    }
  }
}

*/
