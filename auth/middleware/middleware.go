package main

import "net/http"

// hardcore username dan password
const USERNAME = "Reza"
const PASSWORD = "123"

// func Auth
func Auth(w http.ResponseWriter, r *http.Request) bool {
	// pangil built-in func BasicAuth untuk mendapatkan dan mengolah username dan password
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("Something wrong"))
		return false
	}

	// cek username dan password
	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("Invalid username or password"))
		return false
	}

	return true
}

// func untuk mengecek apakah method yang dipakai adalah GET
func AllowOnlyGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "GET" {
		w.Write([]byte("Only GET method is allowed"))
		return false
	}

	return true
}
