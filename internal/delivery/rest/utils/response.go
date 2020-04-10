package utils

import (
	"encoding/json"
	"net/http"
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
)

func WriteResponse(w http.ResponseWriter, req *http.Request, start time.Time, status int, data interface{}, messages ...string) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
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
