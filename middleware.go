package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

type (
	Options struct {
		AllowCredentials bool
		AllowHeaders     []string
		AllowMethods     []string
		ExposeHeaders    []string
	}

	CORSMiddleware struct {
		allowCredentials bool
		allowHeaders     []string
		allowMethods     []string
		exposeHeaders    []string
	}
)

func NewCORSMiddleware(opt *Options) *CORSMiddleware {
	return header(opt)
}

func (m *CORSMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.setHeader(w, r)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

func (m *CORSMiddleware) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.setHeader(w, r)

		switch r.Method {
		case http.MethodOptions:
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	})
}

func (m *CORSMiddleware) setHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	allowHeader := strings.Join(m.exposeHeaders, ",")
	allowCredentials := strconv.FormatBool(m.allowCredentials)
	allowMethods := strings.Join(m.allowMethods, ",")
	exposeHeaders := strings.Join(m.exposeHeaders, ",")

	if origin == "" && !m.allowCredentials {
		origin = "*"
	}

	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Headers", allowHeader)
	w.Header().Set("Access-Control-Allow-Credentials", allowCredentials)
	w.Header().Set("Access-Control-Allow-Methods", allowMethods)
	w.Header().Set("Access-Control-Expose-Headers", exposeHeaders)
}

func header(opt *Options) *CORSMiddleware {
	allowCredentials := opt.AllowCredentials
	allowHeaders := opt.AllowHeaders
	allowMethods := opt.AllowMethods
	exposeHeaders := opt.ExposeHeaders

	if !allowCredentials {
		allowCredentials = false
	}

	if allowHeaders == nil {
		allowHeaders = []string{"Content-Type", "X-CSRF-Token", "Authorization", "AccessToken", "Token"}
	}

	if allowMethods == nil {
		allowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}
	}

	if exposeHeaders == nil {
		exposeHeaders = []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"}
	}

	return &CORSMiddleware{
		allowCredentials: allowCredentials,
		allowHeaders:     allowHeaders,
		allowMethods:     allowMethods,
		exposeHeaders:    exposeHeaders,
	}
}
