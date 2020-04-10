package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/service"
	"log"
	"net/http"
	"os"

	"github.com/google/gops/agent"
)

func initRouter(router *mux.Router, svc *service.Services) {
	router.HandleFunc("/", svc.IndexHandler).Methods("GET")
	router.HandleFunc("/{shortUrl}", svc.RedirectHandler).Methods("GET")
	router.HandleFunc("/list/", svc.GetListDataHandler).Methods("GET")
	router.HandleFunc("/create", svc.CreateURLHandler).Methods("POST")
	router.HandleFunc("/update", svc.UpdateURLHandler).Methods("PUT")
	router.HandleFunc("/delete", svc.DeleteURLHandler).Methods("DELETE")
	router.NotFoundHandler = svc.DefaultService
}

func Start(app *app.UrlShortenerApp) {
	port := os.Getenv("PORT")
	if port == "" {
		port = app.Cfg.Flag.Port
	}

	svc := service.GetServices(app)
	if err := agent.Listen(agent.Options{}); err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()
	initRouter(router, svc)

	fmt.Println("Apps served on :" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":" + port), router))
}
