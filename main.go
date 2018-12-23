package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	PORT = ":8000"
)

func main() {

	//var s string = "s"
	//var s2  = "s"
	// s3 := "s"
	//var a [5]int
	//var s = []int {1,2}
	//var s2 = [...]int {1,2}
	//myarray := [...]int{1,2}

	/*http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var att = "<h1>" + req.URL.Host + "</h1>" + "\n"
		att += req.URL.Query().Get("token")
		att += req.FormValue("username")
		fmt.Fprintf(w, "<h1>"+"Salam Azizam Fisrt Web Server"+"</h1>", att)
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Start Server On Port... "+PORT)
	http.ListenAndServe(PORT,nil)*/

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var att = "<h1>" + req.URL.Host + "</h1>" + "\n"
		vars := mux.Vars(req)

		att += vars["token"]
		att += vars["username"]
		fmt.Fprintf(w, "<h1>"+"Salam Azizam Fisrt Web Server"+"</h1>", att)
	})
	fmt.Println("Start Server On Port... " + PORT)
	http.ListenAndServe(":8000", r)

}
