package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/chrisaandes/twitter-clone/middleware"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Router ...*/
func Router() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middleware.CheckDb()).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
