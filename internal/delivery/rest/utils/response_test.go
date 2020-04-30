package utils

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func Test_getHeader(t *testing.T) {
	processTime := time.Since(time.Now()).Seconds()
	messages := errors.New("error message").Error()

	header := ResponseHeader{
		ProcessTime: processTime,
		Messages:    []string{messages},
	}

	type args struct {
		processTime float64
		messages    []string
	}
	tests := []struct {
		name string
		args args
		want ResponseHeader
	}{
		{
			name: "Get Header Response Success",
			args: args{
				processTime: processTime,
				messages:    []string{messages},
			},
			want: header,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHeader(tt.args.processTime, tt.args.messages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBodyResponse(t *testing.T) {
	processTime := time.Since(time.Now()).Seconds()
	messages := []string{errors.New("error message").Error()}

	bodyWant := getBodyResponse(processTime, nil, messages)

	type args struct {
		processTime float64
		data        interface{}
		messages    []string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "get body response",
			args: args{
				processTime: processTime,
				data:        nil,
				messages:    messages,
			},
			want: bodyWant,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBodyResponse(tt.args.processTime, tt.args.data, tt.args.messages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBodyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
