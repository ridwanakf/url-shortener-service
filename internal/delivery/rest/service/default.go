package service

import (
	"github.com/ridwanakf/url-shortener-service/internal/delivery/rest/route"
	"net/http"
	"time"
)

type DefaultService struct {
}

func NewDefaultService() *DefaultService {
	return &DefaultService{
	}
}

var NotFound = struct {
	Message string `json:"message"`
}{
	Message: "404 Not Found!",
}

func (h *DefaultService) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()
	route.WriteResponse(w, req, start, http.StatusNotFound, NotFound)
}
