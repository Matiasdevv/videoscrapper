package main

import (
	handlers "main/handlers"
)


func main(){

	// r := mux.NewRouter()


	// r.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.ScrapHandler(w,r)
	// })
	
	handlers.ScrapHandler()
}