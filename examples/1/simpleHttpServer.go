package main

import(
        "fmt"
	"net/http"
	"github.com/gorilla/mux"

)

func Ping(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Pong")
}

func main(){
	fmt.Println("Here we go")
	route:= mux.NewRouter()
	route.HandleFunc("/ping", Ping)
	http.ListenAndServe("localhost:8000", route)
}
