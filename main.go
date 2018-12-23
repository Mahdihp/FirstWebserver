package main

import (
	"fmt"
	"net/http"
)

const strConst string = "Salam"

var (
	wow string = "sss"
)

func main() {

	//var s string = "s"
	//var s2  = "s"
	// s3 := "s"
	//var a [5]int
	//var s = []int {1,2}
	//var s2 = [...]int {1,2}
	//myarray := [...]int{1,2}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		var att = "<h1>" + req.URL.Host + "</h1>" + "\n"
		fmt.Fprintf(w, "<h1>"+"Salam Azizam Fisrt Web Server"+"</h1>", att)
	})
	http.ListenAndServe(":8000", nil)

}
