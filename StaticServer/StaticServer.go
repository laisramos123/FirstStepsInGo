package StaticServer

import (
	"log"
	"net/http"
)

func StaticServer() {
	fs := http.FileServer(http.Dir("./Static"))
	http.Handle("/", fs)
	log.Println("Listening on: 3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
