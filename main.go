package main

func main() {
	server := NewServer()
	StartServer(server)

	// p := pat.New()
	// p.Get("/ws", WSHandler)

	// n := negroni.Classic()
	// n.UseHandler(p)

	// http.ListenAndServe(":3000", n)
}
