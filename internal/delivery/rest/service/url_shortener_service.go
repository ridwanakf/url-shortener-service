package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ridwanakf/url-shortener-service/internal"
	"github.com/ridwanakf/url-shortener-service/internal/app"
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/utils"
	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

type ShortenerService struct {
	uc internal.ShortenerUC
}

func NewShortenerService(app *app.UrlShortenerApp) *ShortenerService {
	return &ShortenerService{
		uc: app.UseCases.ShortenerUC,
	}
}

func (s *ShortenerService) IndexHandler(c *gin.Context) {
	http.Error(c.Writer, "Forbidden", http.StatusForbidden)
}

func (s *ShortenerService) RedirectHandler(c *gin.Context) {
	// Check if route is /api/v1/list. This is limitation of Gin

	start := time.Now()

	shortURL := c.Param("shortUrl")

	longURL, err := s.uc.GetLongURL(shortURL)
	if err != nil {
		//or redirect to main page
		utils.WriteResponse(c, start, http.StatusNotFound, NotFound, err.Error())
		return
	}
	c.Redirect(http.StatusMovedPermanently, longURL)
}

func (s *ShortenerService) GetListDataHandler(c *gin.Context) {
	start := time.Now()
	userID := ""
	list, err := s.uc.GetAllURL(userID)
	if err != nil {
		log.Printf("[ShortenerService][GetListDataHandler] error getting list url :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	list = utils.ConvertBatchShortURL(c.Request, list)
	utils.WriteResponse(c, start, http.StatusOK, list, "success")
}

func (s *ShortenerService) CreateURLHandler(c *gin.Context) {
	start := time.Now()

	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][CreateURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][CreateURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, nil, err.Error())
		return
	}
	if bodyReq.LongURL == "" {
		log.Printf("[ShortenerService][CreateURLHandler] longUrl not found!")
		utils.WriteResponse(c, start, http.StatusBadRequest, nil, "longUrl not found!")
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
		utils.WriteResponse(c, start, http.StatusBadRequest, nil, err.Error())
		return
	}

	res := utils.Request{
		ShortURL: utils.ConvertShortURL(c.Request, newURL.ShortURL),
		LongURL:  newURL.LongURL,
	}

	utils.WriteResponse(c, start, http.StatusOK, res, "success")
}

func (s *ShortenerService) UpdateURLHandler(c *gin.Context) {
	start := time.Now()
	status := utils.ResponseBoolean{Status: "failed"}

	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}
	if bodyReq.LongURL == "" {
		log.Printf("[ShortenerService][UpdateURLHandler] longUrl not found!")
		utils.WriteResponse(c, start, http.StatusBadRequest, status, "longUrl not found!")
		return
	}

	err = s.uc.UpdateShortURL(bodyReq.ShortURL, bodyReq.LongURL)
	if err != nil {
		log.Printf("[ShortenerService][UpdateURLHandler] error updating short url :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}

	status.Status = "success"

	utils.WriteResponse(c, start, http.StatusOK, status, "success")
}

func (s *ShortenerService) DeleteURLHandler(c *gin.Context) {
	start := time.Now()
	status := utils.ResponseBoolean{Status: "failed"}

	body, err := ioutil.ReadAll(c.Request.Body)
	defer c.Request.Body.Close()
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error opening body :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error unmarshal :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}

	err = s.uc.DeleteURL(bodyReq.ShortURL)
	if err != nil {
		log.Printf("[ShortenerService][DeleteURLHandler] error deleting short url :%+v\n", err)
		utils.WriteResponse(c, start, http.StatusBadRequest, status, err.Error())
		return
	}

	status.Status = "success"

	utils.WriteResponse(c, start, http.StatusOK, status, "success")
}
