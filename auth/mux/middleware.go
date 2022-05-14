package main

import "net/http"

// hardcore username dan password
const USERNAME = "Reza"
const PASSWORD = "123"

// struct untuk custommux
type CustomMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

// ====================== MUX DEFAULT ======================
// func Auth
func MiddlewareAuth(next http.Handler) http.Handler {
	// return untuk http.Handler yang mana di dalam nya terdapat pengecekan username dan password
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// cek username dan password terhadap basic auth
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Something went wrong"))
			return
		}

		// validasi username dan password
		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			w.Write([]byte("Username atau Password Salah"))
			return
		}

		// jika lolos validasi
		next.ServeHTTP(w, r)

	})
}

// func untuk mengecek apakah method yang dipakai adalah GET
func MiddlewareAllowOnlyGet(next http.Handler) http.Handler {
	// return untuk http.Handler
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Write([]byte("Only GET method is allowed"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

// ====================== MUX CUSTOM ======================
func (c *CustomMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CustomMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux
	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
