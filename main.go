package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"saiga/middleware"
	h "saiga/pkg/helpers"
	"saiga/routes"
)

func main() {
	h.Initialization()

	router := mux.NewRouter()
	router.HandleFunc("/api/user/login", routes.UserLogin).Methods("POST")
	router.HandleFunc("/api/user/signup", routes.SignUp).Methods("POST")
	router.HandleFunc("/api/welcome", routes.Welcome).Methods("GET")
	router.HandleFunc("/api/tasks", middleware.IsAuthorized(routes.ListAllTasks)).Methods("GET")
	router.HandleFunc("/api/chats/{id}/messages", middleware.IsAuthorized(routes.ChatHistory)).Methods("GET")
	router.HandleFunc("/api/attachments/{id}/download", middleware.IsAuthorized(routes.DownloadChatMessage)).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
