package main

import (
	"log"
	"net/http"
	"tugas-golang/database"
	"tugas-golang/routes"
)

func main() {
	database.ConnectDatabase()
	mux := routes.SetAPI()
	log.Println("Go server running on port 8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
