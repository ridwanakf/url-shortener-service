package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"

	"github.com/ridwanakf/url-shortener-service/internal/entity"
)

func TestConvertBatchShortURL(t *testing.T) {
	req, urls := getDummyData()
	exp := getExpectedData()
	got := ConvertBatchShortURL(req, urls)

	assert.Equal(t, exp[0].ShortURL, got[0].ShortURL)
	assert.Equal(t, exp[1].ShortURL, got[1].ShortURL)
}

func TestConvertShortURL(t *testing.T) {
	req, urls := getDummyData()
	exp := getExpectedData()
	got := ConvertShortURL(req, urls[0].ShortURL)

	assert.Equal(t, exp[0].ShortURL, got)
}

func getDummyData() (*http.Request, []entity.URL){
	req, _ := http.NewRequest("GET", "/", nil)
	req.Host = "http://test.com"

	urls := []entity.URL{
		{
			ShortURL: "aBcDe",
		},
		{
			ShortURL: "a2c3e",
		},
	}

	return req, urls
}

func getExpectedData() []entity.URL{
	urls := []entity.URL{
		{
			ShortURL: "http://test.com/aBcDe",
		},
		{
			ShortURL: "http://test.com/a2c3e",
		},
	}

	return urls
}