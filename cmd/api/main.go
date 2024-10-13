package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/Jonathan0823/API-with-GO/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main(){
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting API service...")

	fmt.Println("API service started")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		log.Error(err)
	}
	
}