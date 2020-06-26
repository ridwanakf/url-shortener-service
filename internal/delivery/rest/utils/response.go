package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

type (
	Response struct {
		Header ResponseHeader `json:"header"`
		Data   interface{}    `json:"data"`
	}

	ResponseHeader struct {
		ProcessTime float64  `json:"process_time"`
		Messages    []string `json:"messages"`
		ErrorCode   string   `json:"error_code"`
	}

	ResponseBoolean struct {
		Status string `json:"status"`
	}
)

func WriteResponse(c *gin.Context, start time.Time, status int, data interface{}, messages ...string) {
	w := c.Writer
	w.WriteHeader(status)
	w.Write(getBodyResponse(time.Since(start).Seconds(), data, messages))
	return
}

func getBodyResponse(processTime float64, data interface{}, messages []string) []byte {
	header := getHeader(processTime, messages)
	response := Response{
		Header: header,
		Data:   data,
	}

	dataBytes, _ := json.Marshal(response)
	return dataBytes
}

func getHeader(processTime float64, messages []string) ResponseHeader {
	header := ResponseHeader{
		ProcessTime: processTime,
		Messages:    []string{},
	}

	if len(messages) > 0 {
		header.Messages = messages
	}

	return header
}
