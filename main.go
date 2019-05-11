package main
// try not to change the comments below:


//#cgo pkg-config: opencv 
// #cgo LDFLAGS: -lcurl
// #include "grab.hxx"
//typedef int (*intFunc) ();
//int intFunction(intFunc f)
//{
//  return f();
//}
import "C"

import (
   
     "github.com/gorilla/websocket"
    "net/http"
  "strconv"
    "fmt"
    "encoding/json"
    // "io/ioutil"
);



type parameters struct {
Img_url string   
Startx int      
Starty int      
Width int        
Height int            

}

var returned_value = 1100


var upgrader = websocket.Upgrader{ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {

http.HandleFunc("/image",func(w http.ResponseWriter , r * http.Request){
  
var conn , err =upgrader.Upgrade(w,r,nil)
if  err!=nil{
return

}
for {
mType,p,err:=conn.ReadMessage()

if err!=nil{

return 
}
// fmt.Println("(testing the data)here is the data: ",string(p))



returned_value := strconv.Itoa(returned_value)


// parsing json into parameters:
data := parameters{}
err =json.Unmarshal(p,&data)
if err!=nil{
fmt.Println("error parsing json: ",err.Error())
return

}
// fmt.Println("this is a test for parsed img_url:",data.Img_url)

 C.foregroundDetection(C.CString(data.Img_url),C.int(data.Startx),C.int(data.Starty),C.int(data.Width),C.int(data.Height))

b := []byte(returned_value)

err = conn.WriteMessage(mType,b)
if err!=nil{
  return
}


}



})
http.ListenAndServe(":3000", nil)


fmt.Println("I'm go and I'm finished")

}



