package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

type A struct {
}

type B struct {
	A A
}

func (p *B) String() string {
	return fmt.Sprintf("B[A[%v]]", p.A)
}

func TestBasics(t *testing.T) {
	var b B
	fmt.Printf("v is: %v\n", &b)

	var i interface{}
	fmt.Printf("i is: %v\n", i)

}
