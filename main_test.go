package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	cases := []string{
		"Tiago",
		"João",
		"Maria",
		"Fernanda",
	}

	for _, c := range cases {
		req, err := http.NewRequest("GET", "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		qs := req.URL.Query()
		qs.Add("name", c)

		req.URL.RawQuery = qs.Encode()

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(helloWorld)
		handler.ServeHTTP(rr, req)

		expected := fmt.Sprintf("Hello %s", c)

		if rr.Body.String() != expected {
			t.Errorf("Erro: esperado '%s', recebido: %s", expected, rr.Body.String())
		}
	}
}

func TestHealth(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(health)
	handler.ServeHTTP(rr, req)

	if rr.Body.String() != "OK" {
		t.Errorf("Erro: esperado 'OK', recebido: %s", rr.Body.String())
	}
}
