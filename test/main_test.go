package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
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

type CError struct {
}

func (c CError) Error() string {
	return "CError"
}

func TestA(t *testing.T) {
	var i []byte = nil
	fmt.Println(len(i))
	/* cause := errors.New("whoops")
	fmt.Printf("%+v\n", cause)

	err := errors.WithStack(cause)
	fmt.Printf("%+v\n", err)

	fmt.Printf("%+v\n", CError{}) */

	// fmt.Printf("%+v\n", f())
}

func f() error {
	cause := errors.New("whoops")
	err := errors.WithStack(cause)
	return err
}
