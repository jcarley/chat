package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/pat"
)

var (
	router  *negroni.Negroni
	session *r.Session
)

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "localhost:28015",
		Database: "chat",
		MaxOpen:  40,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}

	r.DbCreate("chat").RunWrite(session)
	r.Db("chat").TableCreate("messages").Run(session)

	// cursor, err := r.DbList().Run(session)
	// if err != nil {
	// log.Fatalln(err.Error())
	// }
	// defer cursor.Close()

	// var response interface{}
	// found := false
	// for cursor.Next(&response) {
	// if response == "chat" {
	// found = true
	// }
	// }

	// if !found {
	// if err != nil {
	// log.Fatalln(err.Error())
	// }
	// }

}

func NewServer() *http.Server {
	router = initRouting()
	return &http.Server{
		Addr:    ":4000",
		Handler: router,
	}
}

func StartServer(server *http.Server) {
	log.Println("Starting server")

	err := server.ListenAndServeTLS("ssl/cert.pem", "ssl/key.pem")
	if err != nil {
		log.Fatalln("Error: %v", err)
	}

}

type Message struct {
	ID       string    `gorethink:"id,omitempty"`
	Username string    `gorethink:"username"`
	Message  string    `gorethink:"message"`
	Created  time.Time `gorethink:"created"`
}

func postMessage(w http.ResponseWriter, req *http.Request) {
	reader := req.Body

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var payload map[string]string
	err = json.Unmarshal(data, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := Message{
		Username: payload["username"],
		Message:  payload["message"],
	}
	message.Created = time.Now()

	_, err = r.Db("chat").Table("messages").Insert(message).RunWrite(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("{}"))
}

func initRouting() *negroni.Negroni {
	r := pat.New()

	// Add handlers for routes
	r.Post("/message", postMessage)

	// Add handlers for websockets
	r.Get("/ws", newChangesHandler(messageChanges))

	n := negroni.Classic()
	n.UseHandler(r)

	return n
}

func newChangesHandler(fn func(chan interface{})) http.HandlerFunc {
	h := newHub()
	go h.run()

	fn(h.broadcast)

	return wsHandler(h)
}
