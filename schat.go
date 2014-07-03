package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
    "github.com/gorilla/sessions"
   	"github.com/gorilla/context"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type message struct {
	Name string
	Body string
	Readed bool
}

type user struct {
	Name string
	Password string
}

var user1 user
var user2 user

var lastMessage message 

func checkLogin(u user) (bool){
	var ret bool 
	ret = false
	if u.Name == user1.Name {
		if u.Password == user1.Password {
			ret = true
		} else {
			ret = false
		}
	} else if u.Name == user2.Name {
		if u.Password == user2.Password {
			ret = true
		} else {
			ret = false
		}
	}
	return ret
}

func checkSession(r *http.Request) (bool){
	session, _ := store.Get(r, "loginSession")
    if val, ok := session.Values["username"].(string); ok {
        switch val {
        case "": return false
        default: return true
        }
    } else {
        return false
    }
}

func renderIndex(w http.ResponseWriter, r *http.Request) {
	if checkSession(r) {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		renderLogin(w,r)		
	}
}

func renderLogin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("login.html")
	t.Execute(w, nil)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	respon := r.Body
	if respon == nil {
		w.WriteHeader(400)
	} else {
		decoder := json.NewDecoder(r.Body)
		var u user
		decoder.Decode(&u)

		if checkLogin(u) {
			log.Println("buen login")
			session, _ := store.Get(r, "loginSession")
		    // Set some session values.
		    session.Values["username"] = u.Name
		    // Save it.
		    session.Save(r, w)
			http.Redirect(w,r,"index",http.StatusFound) 
		} else {
			w.WriteHeader(400)
			fmt.Fprint(w, "User or password incorrect")
		}
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

		if checkSession(r) {
			lastMessage.Name = m.Name
			lastMessage.Body = m.Body
			lastMessage.Readed = false

			resA, _ := json.Marshal(m)
			log.Println(string(resA))
			fmt.Fprint(w, string(resA))
		} else {
			w.WriteHeader(400)
			fmt.Fprint(w, "No session")
		}
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
		
		if checkSession(r) {
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
		} else {
			w.WriteHeader(400)
			fmt.Fprint(w, "No session")		
		}
	}
}

//Provisional data
func Initialize() {
	user1.Name = "mario"
	user1.Password = "1234"

	user2.Name = "luigi"
	user2.Password = "1234"
}

func main() {
	Initialize()
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/index", renderIndex)
	http.HandleFunc("/", renderLogin)
	http.HandleFunc("/messages", saveMessage)
	http.HandleFunc("/getMessages", displayMessages)
	http.HandleFunc("/login", loginUser)
	http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", context.ClearHandler(http.DefaultServeMux))
}
