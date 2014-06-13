package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
    "github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type message struct {
	Name string
	Body string
	Readed bool
}

type user struct {
	Name string
}

var lastMessage message 

func render(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func renderLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	log.Println("ok")
	http.Redirect(w,r,"index",http.StatusFound) 
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

		lastMessage.Name = m.Name
		lastMessage.Body = m.Body
		lastMessage.Readed = false

		resA, _ := json.Marshal(m)
		log.Println(string(resA))
		fmt.Fprint(w, string(resA))
	}
}

func displayMessages(w http.ResponseWriter, r *http.Request) {
	respon := r.Body
	if respon == nil {
		w.WriteHeader(400)
	} else {
		decoder := json.NewDecoder(r.Body)
		var u user
		decoder.Decode(&u)

		if u.Name == lastMessage.Name {
			w.WriteHeader(204)
		} else if lastMessage.Readed == true {
			w.WriteHeader(204)
		} else {
			lastMessage.Readed = true
			resA, _ := json.Marshal(lastMessage)
			log.Println(string(resA))
			fmt.Fprint(w, string(resA))
		}
	}
}

func main() {
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", renderLogin)
	http.HandleFunc("/index", render)
	http.HandleFunc("/messages", messagesSelecter)
	http.HandleFunc("/getMessages", displayMessages)
	http.HandleFunc("/login", loginUser)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
}
