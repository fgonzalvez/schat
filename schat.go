package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var counter int = 0
var vector []message

type message struct {
	Id   int
	Name string
	Body string
}

func render(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func renderLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func messagesSelecter(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		displayMessages(w, r)
	} else if r.Method == "POST" {
		saveMessage(w, r)
	} else {
		return
	}
}

func saveMessage(w http.ResponseWriter, r *http.Request) {
	respon := r.Body
	if respon == nil {
		w.WriteHeader(400)
	} else {
		decoder := json.NewDecoder(r.Body)
		var m message
		decoder.Decode(&m)

		m.Id = counter
		counter++
		//vector := vector.append(m)
		//Guardar m en lista

		resA, _ := json.Marshal(m)
		log.Println(string(resA))
		fmt.Fprint(w, string(resA))
	}
}

func displayMessages(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprint(w, "success")
	/*m := Message{"Alice", "Hello"}
	//b, err := json.Marshal(m)
	enc := json.NewEncoder(w)
	enc.Encode(m)
	w.WriteHeader(200)
	fmt.Fprint(w, "success")
	/*w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, "success")
	return*/
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/index", render)
	http.HandleFunc("/", renderLogin)
	http.HandleFunc("/messages", messagesSelecter)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
