package rest

import (
	"os"

	"github.com/labstack/echo"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/server"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/service"
)

func initAPIHandler(eg *echo.Group, svc *service.Services) {
	eg.GET("/list", svc.GetListDataHandler)
	eg.POST("/create", svc.CreateURLHandler)
	eg.PUT("/update", svc.UpdateURLHandler)
	eg.DELETE("/delete", svc.DeleteURLHandler)
}

func Start(app *app.UrlShortenerApp) {
	port := os.Getenv("PORT")
	if port != "" {
		app.Cfg.Server.Port = port
	}

	srv := server.New()
	svc := service.GetServices(app)

	srv.GET("/", svc.IndexHandler)
	srv.GET("/:shortUrl", svc.RedirectHandler)

	api := srv.Group("/api/v1")

	// API Handler
	initAPIHandler(api, svc)

	server.Start(srv, &app.Cfg.Server)
}
