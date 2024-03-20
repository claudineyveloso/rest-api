package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudineyveloso/rest-api.git/configs"
	"github.com/gorilla/mux"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Bem vindo a Rest API!"))
		if err != nil {
			http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
			return
		}
	}).Methods("GET")
	port := configs.GetServerPort()
	fmt.Printf("Servidor escutando na porta %s\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}

}
