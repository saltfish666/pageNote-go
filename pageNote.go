package main

import (
	"net/http"
	"fmt"
	"./router"
)

func indexHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "server is runing")
}

func main(){

	http.HandleFunc("/",indexHandler)
	http.HandleFunc("/user", router.UserHandler)
	http.HandleFunc("/note", router.NoteHandler)
	http.ListenAndServe(":8080",nil)
}
