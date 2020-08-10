package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type BasicHandler struct{}

func (b BasicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 5)
	w.Write([]byte("1"))
	w.WriteHeader(200)
}

func TestWaitForUrl(t *testing.T) {
	var testServer = httptest.NewServer(BasicHandler{})
	defer testServer.Close()

	type args struct {
		url      string
		status   int
		duration time.Duration
	}
	tests := []struct {
		name      string
		args      args
		wantReady bool
	}{
		{
			name: "1", args: args{
				url:      testServer.URL,
				status:   200,
				duration: time.Second * 6,
			},
			wantReady: true,
		},
		{
			name: "2",
			args: args{
				url:      testServer.URL,
				status:   300,
				duration: time.Second * 6,
			},
			wantReady: false,
		},
		{
			name: "3", args: args{
				url:      testServer.URL,
				status:   200,
				duration: time.Second * 4,
			},
			wantReady: false,
		},
		{
			name: "4", args: args{
				url:      testServer.URL,
				status:   300,
				duration: time.Second * 4,
			},
			wantReady: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReady := WaitForUrl(tt.args.url, tt.args.status, tt.args.duration); gotReady != tt.wantReady {
				t.Errorf("WaitForUrl() = %v, want %v", gotReady, tt.wantReady)
			}
		})
	}
}
