package main

import (
	"fmt"
	"net/http"
)
func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"<h1>Hello GoLang</h1>")
	})
	http.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer,"<h1>WebRtC</h1>")
	})
	http.ListenAndServe(":8080",nil)

}
