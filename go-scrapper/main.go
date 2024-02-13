package goscrapper

import (
	"main/handlers"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){

	r := mux.NewRouter()


	r.HandleFunc('/',func(w http.ResponseWriter, r *http.Request) {
		handlers.ScrapHandler(w,r)
	})
}