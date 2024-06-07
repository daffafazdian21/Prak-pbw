package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "Muhammad Daffa Fazdian",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("Tugas_5_2_4521210003.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:5000/")
	http.ListenAndServe(":5000", nil)
}
