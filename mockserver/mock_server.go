package mockserver

import (
	"net/http"
	"net/http/httptest"
)

var Message = []byte(`{
	"data": "Hello world!"
  }`)

func ThirdPartyAPIServerMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/v1/message/1", MockHandleGetMessage)

	return httptest.NewServer(handler)
}

func MockHandleGetMessage(w http.ResponseWriter, r *http.Request) {
	if r.URL.String()+"/api/v1/message/1" != r.RequestURI+"/api/v1/message/1" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"record not found"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(Message)
}
