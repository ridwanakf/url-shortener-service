package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
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

func (s *ShortenerService) IndexHandler(c echo.Context) error {
	return c.File("./web/build/index.html")
}

func (s *ShortenerService) RedirectHandler(c echo.Context) error {
	shortURL := c.Param("shortUrl")

	longURL, err := s.uc.GetLongURL(shortURL)
	if err != nil {
		return c.JSON(http.StatusNotFound, NotFound)
	}
	return c.Redirect(http.StatusMovedPermanently, longURL)
}

func (s *ShortenerService) GetListDataHandler(c echo.Context) error {
	userID := ""
	list, err := s.uc.GetAllURL(userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error()})
	}

	list = utils.ConvertBatchShortURL(c.Request(), list)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    list,
		"message": "success"},
	)
}

func (s *ShortenerService) CreateURLHandler(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error()})
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}
	if bodyReq.LongURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "longUrl not found!"})
	}

	var newURL entity.URL
	if bodyReq.ShortURL == "" {
		newURL, err = s.uc.CreateNewShortURL(bodyReq.LongURL)
	} else {
		newURL, err = s.uc.CreateNewCustomShortURL(bodyReq.ShortURL, bodyReq.LongURL)
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}

	res := utils.Request{
		ShortURL: utils.ConvertShortURL(c.Request(), newURL.ShortURL),
		LongURL:  newURL.LongURL,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":    res,
		"message": "success"},
	)
}

func (s *ShortenerService) UpdateURLHandler(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error()})
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}
	if bodyReq.LongURL == "" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "longUrl not found!"})
	}

	err = s.uc.UpdateShortURL(bodyReq.ShortURL, bodyReq.LongURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success"},
	)
}

func (s *ShortenerService) DeleteURLHandler(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error()})
	}

	var bodyReq utils.Request
	err = json.Unmarshal(body, &bodyReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}

	err = s.uc.DeleteURL(bodyReq.ShortURL)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success"},
	)
}
