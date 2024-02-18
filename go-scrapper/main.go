package main

import (
	handlers "main/handlers"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){

	r := mux.NewRouter()


	r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		handlers.ScrapHandler(w,r)
	})
	
	http.ListenAndServe(":8003",r)
}