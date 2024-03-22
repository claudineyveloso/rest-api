package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/claudineyveloso/rest-api.git/cmd/api"
	"github.com/claudineyveloso/rest-api.git/configs"
	"github.com/claudineyveloso/rest-api.git/db"
)

func main() {
	cfg := configs.Config{
		PublicHost: configs.Envs.PublicHost,
		Port:       configs.Envs.Port,
		DBUser:     configs.Envs.DBUser,
		DBPassword: configs.Envs.DBPassword,
		DBName:     configs.Envs.DBName,
	}
	db, err := db.NewPostgresSQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	// r := mux.NewRouter()
	// r.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
	// 	_, err := w.Write([]byte("Bem-vindo ao Rest-API-YT!"))
	// 	if err != nil {
	// 		http.Error(w, "Erro ao escrever resposta", http.StatusInternalServerError)
	// 		return
	// 	}
	// }).Methods("GET")
	// fmt.Printf("Servidor escutando na porta %s\n", "8080")
	// err := http.ListenAndServe(fmt.Sprintf(":%s", "8080"), r)
	// if err != nil {
	// 	log.Fatalf("Erro ao iniciar o servidor: %v", err)
	// }
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DB: Successfully connected!")
}
