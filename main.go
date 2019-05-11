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
   
http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
if(r.Method=="POST")    {

 x,errx:=strconv.ParseInt(r.FormValue("x"),10,32)
 y,erry:=strconv.ParseInt(r.FormValue("y"),10,32)
 width,errw:=strconv.ParseInt(r.FormValue("width"),10,32)
 height,errh:=strconv.ParseInt(r.FormValue("height"),10,32)
img_url := r.FormValue("img_url")
if errx!=nil || erry !=nil || errw!=nil || errh!=nil{
    fmt.Println("error converting string to int")
}
 fmt.Println("this is x:",x," ",y," ",width," ",height) 



 fmt.Println("I'm go and just got an int form c++: ",C.foregroundDetection(C.CString(img_url),C.int(x),C.int(y),C.int(width),C.int(height)))




}
fmt.Println("salam")
})



// http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){

   

    


//     http.ServeFile(w,r,"C:/Users/mahyar/Desktop/opencv projects/go_api/index.html")


    



// })


// http.HandleFunc(

// "/upload",func(w http.ResponseWriter,r *http.Request){


// fmt.Println("File uploading")
// r.ParseMultipartForm(10<<20)



// file,handler,err :=r.FormFile("myFile")

// if err!=nil{
//     fmt.Println("Error Retrieving the File")
//     fmt.Println(err)
//     return

// }
// defer file.Close()
// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
// fmt.Printf("File Size: %+v\n", handler.Size)
// fmt.Printf("MIME Header: %+v\n", handler.Header)



// tempFile,err:= ioutil.TempFile("temp-images","upload-*.png")
// if err !=nil{
// fmt.Println("this is a tempfile error")
// fmt.Println(err)

// }

// defer tempFile.Close()

// fileBytes,err :=ioutil.ReadAll(file)
// if  err !=nil {

// fmt.Fprintf(w, "Successfully Uploaded File\n")


// }


// tempFile.Write(fileBytes)






// })




//  http.HandleFunc("/image",func(w http.ResponseWriter,r *http.Request){

//     var conn, _ = upgrader.Upgrade(w, r, nil)
  
// 		go func(conn *websocket.Conn) {
// 			for {

//                 mType, _, _ := conn.ReadMessage()

//                 var img64 [] byte
//                 mType++
                
                
                    
               
//                 files, _ := ioutil.ReadDir("C:/Users/mahyar/Desktop/opencv projects/go_api/temp-images");
//                 for _, f := range files {
//                 img64, _ = ioutil.ReadFile("C:/Users/mahyar/Desktop/opencv projects/go_api/temp-images/" + f.Name())
                         
//                 }
            
// 			conn.WriteMessage(2,img64)
// 			}
// 		}(conn)


// })
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




}



