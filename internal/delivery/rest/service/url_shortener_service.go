package service

import (
	"log"
	"net/http"
	
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/app"
)

type ShortenerService struct {
	uc internal.ShortenerUC
}

func NewShortenerService(app *app.UrlShortenerApp) *ShortenerService {
	return &ShortenerService{
		uc: app.UseCases.ShortenerUC,
	}
}

func (s *ShortenerService) IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IndexHandler")
}

func (s *ShortenerService) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("RedirectHandler")
}

func (s *ShortenerService) GetListDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("GetListDataHandler")
}

func (s *ShortenerService) CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("CreateURLHandler")
	//custom and random based on query param
}

func (s *ShortenerService) UpdateURLHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("UpdateURLHandler")
}

func (s *ShortenerService) DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteURLHandler")
}
