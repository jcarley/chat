package main

import (
	"log"

	r "github.com/dancannon/gorethink"
)

func messageChanges(ch chan interface{}) {
	go func() {
		for {
			res, err := r.Db("chat").Table("messages").Changes().Run(session)
			if err != nil {
				log.Fatalln(err)
			}

			var response r.WriteChanges
			for res.Next(&response) {
				ch <- response
			}

			if res.Err() != nil {
				log.Println(res.Err())
			}
		}
	}()
}
