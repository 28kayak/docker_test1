package main
//import {
//	fmt 
//	log 
//	net/http 
//}
import "fmt"
import "log"
import "net/http"

func main(){
 http.HandleFunc("/", 
 	func( w http.ResponseWriter, r *http.Request){ 
 		log.Println(" received request!") 
 		fmt.Println( w, "Hello World!!") 
 	}) 
 log.Println(" start server") 
 server := &http.Server{ Addr: ":8080"} 
 server.ListenAndServe() 
}
