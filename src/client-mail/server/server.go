package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/atomarxculo/client-mail/src/client-mail/smtp"
	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pagina principal")
}

func StartServer() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", hello)
	router.HandleFunc("/send", smtp.SendMail).Methods("GET")

	log.Println("Escuchando en el puerto", ":5555")
	log.Fatal(http.ListenAndServe(":5555", router))
}
