package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/ridwanakf/url-shortener-service/internal/app/config"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func New() *gin.Engine {
	g := gin.Default()

	//g.Use(secure.CORS(), secure.Headers(), middleware.Secure())

	g.Static("/", "public/static")
	tmpl := template.Must(template.ParseGlob("public/views/*.html"))
	g.SetHTMLTemplate(tmpl)

	return g
}

func Start(g *gin.Engine, cfg *config.Server) {
	srv := &http.Server{
		Handler:      g,
		Addr:         cfg.Port,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}

	// Start server
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
