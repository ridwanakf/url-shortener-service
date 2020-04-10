package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/utils"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func (s *ShortenerService) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	params := mux.Vars(r)
	shortURL := params["shortUrl"]

	longURL, err := s.uc.GetLongURL(shortURL)
	if err != nil {
		//or redirect to main page
		utils.WriteResponse(w, r, start, http.StatusNotFound, NotFound)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

func (s *ShortenerService) GetListDataHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	list, err := s.uc.GetAllURL()
	if err != nil {
		log.Printf("[ShortenerService][GetListDataHandler] error getting list url :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	utils.WriteResponse(w, r, start, http.StatusOK, list, "success")
}

func (s *ShortenerService) CreateURLHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][CreateURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][CreateURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, nil, err.Error())
		return
	}
	if bodyReq.LongURL == "" {
		log.Printf("[ShortenerService][CreateURLHandler] longUrl not found!")
		utils.WriteResponse(w, r, start, http.StatusBadRequest, nil, "longUrl not found!")
		return
	}

	var newURL entity.URL
	if bodyReq.ShortURL == "" {
		newURL, err = s.uc.CreateNewShortURL(bodyReq.LongURL)
	} else {
		newURL, err = s.uc.CreateNewCustomShortURL(bodyReq.ShortURL, bodyReq.LongURL)
	}
	if err != nil {
		log.Printf("[ShortenerService][CreateURLHandler] error creating short url :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	res := utils.Request{
		ShortURL: newURL.ShortURL,
		LongURL:  newURL.LongURL,
	}

	utils.WriteResponse(w, r, start, http.StatusOK, res, "success")
}

func (s *ShortenerService) UpdateURLHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status := utils.ResponseBoolean{Status: "failed"}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}
	if bodyReq.LongURL == "" {
		log.Printf("[ShortenerService][UpdateURLHandler] longUrl not found!")
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, "longUrl not found!")
		return
	}

	err = s.uc.UpdateShortURL(bodyReq.ShortURL, bodyReq.LongURL)
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error updating short url :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}

	status.Status = "success"

	utils.WriteResponse(w, r, start, http.StatusOK, status, "success")
}

func (s *ShortenerService) DeleteURLHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	status := utils.ResponseBoolean{Status: "failed"}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}

	err = s.uc.DeleteURL(bodyReq.ShortURL)
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error deleting short url :%+v\n", err)
		utils.WriteResponse(w, r, start, http.StatusBadRequest, status, err.Error())
		return
	}

	status.Status = "success"

	utils.WriteResponse(w, r, start, http.StatusOK, status, "success")
}
