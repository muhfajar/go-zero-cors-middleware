package middleware

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewCORSMiddleware(t *testing.T) {
	allowHeadersExpected := []string{"Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"}
	allowMethodsExpected := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	exposeHeadersExpected := []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}

	cors := NewCORSMiddleware(&Options{})

	if cors.allowCredentials {
		t.Errorf("Allow-Credentials: %v; want %v", cors.allowCredentials, false)
	}

	if !reflect.DeepEqual(cors.allowHeaders, allowHeadersExpected) {
		t.Errorf("Allow-Header: %v; want %v", cors.allowHeaders, allowHeadersExpected)
	}

	if !reflect.DeepEqual(cors.allowMethods, allowMethodsExpected) {
		t.Errorf("Allow-Methods: %v; want %v", cors.allowMethods, allowMethodsExpected)
	}

	if !reflect.DeepEqual(cors.exposeHeaders, exposeHeadersExpected) {
		t.Errorf("Expose-Headers: %v; want %v", cors.exposeHeaders, exposeHeadersExpected)
	}
}

func TestNewCORSMiddlewareOptions(t *testing.T) {
	allowHeadersExpected := []string{"Authorization"}
	allowMethodsExpected := []string{"GET"}
	exposeHeadersExpected := []string{"Content-Length"}

	cors := NewCORSMiddleware(&Options{
		AllowCredentials: true,
		AllowHeaders:     allowHeadersExpected,
		AllowMethods:     allowMethodsExpected,
		ExposeHeaders:    exposeHeadersExpected,
	})

	if !cors.allowCredentials {
		t.Errorf("Allow-Credentials: %v; want %v", cors.allowCredentials, false)
	}

	if !reflect.DeepEqual(cors.allowHeaders, allowHeadersExpected) {
		t.Errorf("Allow-Header: %v; want %v", cors.allowHeaders, allowHeadersExpected)
	}

	if !reflect.DeepEqual(cors.allowMethods, allowMethodsExpected) {
		t.Errorf("Allow-Methods: %v; want %v", cors.allowMethods, allowMethodsExpected)
	}

	if !reflect.DeepEqual(cors.exposeHeaders, exposeHeadersExpected) {
		t.Errorf("Expose-Headers: %v; want %v", cors.exposeHeaders, exposeHeadersExpected)
	}
}

func TestCORSMiddleware_Handle(t *testing.T) {
	allowHeadersExpected := "Content-Type,X-CSRF-Token,Authorization,AccessToken,Token"
	allowOriginExpected := "https://muhfajar.id"
	allowMethodsExpected := "GET,POST,PUT,DELETE,OPTIONS,PATCH"

	cors := NewCORSMiddleware(&Options{})

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodOptions, "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Origin", allowOriginExpected)

	cors.Handle(func(w http.ResponseWriter, r *http.Request) {}).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Response status: %d; want %d", status, http.StatusNoContent)
	}

	if header := rr.Header(); header == nil {
		t.Errorf("Response header: nil; want %v", header)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Headers") != allowHeadersExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Headers"), allowHeadersExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Methods") != allowMethodsExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Methods"), allowMethodsExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Origin") != allowOriginExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Origin"), allowOriginExpected)
	}
}

func TestCORSMiddleware_HandleOptions(t *testing.T) {
	allowHeadersExpected := "Authorization"
	allowOriginExpected := "*"
	allowMethodsExpected := "GET"

	cors := NewCORSMiddleware(&Options{
		AllowHeaders:  []string{"Authorization"},
		AllowMethods:  []string{"GET"},
		ExposeHeaders: []string{"Content-Length"},
	})

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodOptions, "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	cors.Handle(func(w http.ResponseWriter, r *http.Request) {}).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Response status: %d; want %d", status, http.StatusNoContent)
	}

	if header := rr.Header(); header == nil {
		t.Errorf("Response header: nil; want %v", header)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Headers") != allowHeadersExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Headers"), allowHeadersExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Methods") != allowMethodsExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Methods"), allowMethodsExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Origin") != allowOriginExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Origin"), allowOriginExpected)
	}
}

func TestCORSMiddleware_Handler(t *testing.T) {
	allowHeadersExpected := "Content-Type,X-CSRF-Token,Authorization,AccessToken,Token"
	allowOriginExpected := "https://muhfajar.id"
	allowMethodsExpected := "GET,POST,PUT,DELETE,OPTIONS,PATCH"

	cors := NewCORSMiddleware(&Options{})

	rr := httptest.NewRecorder()
	handler := cors.Handler()
	req, err := http.NewRequest(http.MethodOptions, "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Origin", "https://muhfajar.id")

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Response status: %d; want %d", status, http.StatusNoContent)
	}

	if header := rr.Header(); header == nil {
		t.Errorf("Response header: nil; want %v", header)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Headers") != allowHeadersExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Headers"), allowHeadersExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Methods") != allowMethodsExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Methods"), allowMethodsExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Origin") != allowOriginExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Origin"), allowOriginExpected)
	}
}

func TestCORSMiddleware_HandlerOptions(t *testing.T) {
	allowHeadersExpected := "Authorization"
	allowOriginExpected := "*"
	allowMethodsExpected := "GET"

	cors := NewCORSMiddleware(&Options{
		AllowHeaders:  []string{"Authorization"},
		AllowMethods:  []string{"GET"},
		ExposeHeaders: []string{"Content-Length"},
	})

	rr := httptest.NewRecorder()
	handler := cors.Handler()
	req, err := http.NewRequest(http.MethodOptions, "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Response status: %d; want %d", status, http.StatusNoContent)
	}

	if header := rr.Header(); header == nil {
		t.Errorf("Response header: nil; want %v", header)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Headers") != allowHeadersExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Headers"), allowHeadersExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Methods") != allowMethodsExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Methods"), allowMethodsExpected)
	}

	if header := rr.Header(); header.Get("Access-Control-Allow-Origin") != allowOriginExpected {
		t.Errorf("Response header: %s; want %v", header.Get("Access-Control-Allow-Origin"), allowOriginExpected)
	}
}
