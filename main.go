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
		var att = req.URL.Host + "\n"
		fmt.Fprintf(w, "Salam Azizam Fisrt Web Server", att)
	})
	http.ListenAndServe(":8000", nil)

}
