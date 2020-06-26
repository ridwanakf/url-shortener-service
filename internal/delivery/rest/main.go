package rest

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/server"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/service"
)
func initAPIHandler(g *gin.RouterGroup, svc *service.Services) {
	g.POST("/create", svc.CreateURLHandler)
	g.PUT("/update", svc.UpdateURLHandler)
	g.DELETE("/delete", svc.DeleteURLHandler)
}

func Start(app *app.UrlShortenerApp) {
	port := os.Getenv("PORT")
	if port == "" {
		port = app.Cfg.Server.Port
	}

	srv := server.New()
	svc := service.GetServices(app)

	srv.GET("/", svc.IndexHandler)
	srv.GET("/:shortUrl/", svc.RedirectHandler)

	api := srv.Group("/api/v1")

	// API Handler
	initAPIHandler(api, svc)

	server.Start(srv, &app.Cfg.Server)
}
