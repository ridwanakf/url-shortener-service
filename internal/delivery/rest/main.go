package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/server"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/service"
	"log"
	"net/http"
	"os"
)

func initHandler(g *gin.Engine, svc *service.Services) {
	g.GET("/", svc.IndexHandler)
	g.GET("/{shortUrl}", svc.RedirectHandler).Methods("GET")
	g.GET("/list/", svc.GetListDataHandler).Methods("GET")
	g.POST("/create", svc.CreateURLHandler).Methods("POST")
	g.PUT("/update", svc.UpdateURLHandler).Methods("PUT")
	g.DELETE("/delete", svc.DeleteURLHandler).Methods("DELETE")
}

func Start(app *app.UrlShortenerApp) {
	port := os.Getenv("PORT")
	if port == "" {
		port = app.Cfg.Server.Port
	}

	srv := server.New()
	srv.Group("/api/v1")

	svc := service.GetServices(app)

	initHandler(srv, svc)

	fmt.Println("Apps served on :" + port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":"+port), router))
}
