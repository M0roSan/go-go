package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// Test the a GET request to the login page returns HTTP error with 401 for an authenticated user
func TestShowLoginPageAuthenticated(t *testing.T) {
	// Create response recorder
	w := httptest.NewRecorder()

	// get a new router
	r := getRouter(true)

	// set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})

	// Defin the orute similar to its definition
	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/u/login", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	// Create the service and process the above request
	r.ServeHTTP(w, req)

	// Test that the http status code is 401
	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a GET request to the login page returns the login page with the hTTP code 200 for unauthenticated user
func TestShowLoginPageUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.GET("/u/login", ensureNotLoggedIn(), showLoginPage)

	req, _ := http.NewRequest("GET", "/u/login", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Login</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a POST request to the login route returns an HTTP error with 401 for an authenticated user
func TestLoginAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)

	// Set the token cookie to simulate an authenticated user
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.POST("/u/login", ensureNotLoggedIn(), performLogin)

	loginPayload := getLoginPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fail()
	}
}

// Test that a POST request to the login route returns an success message for an unauthenticated user
func TestLoginUnauthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)

	r.POST("/u/login", ensureNotLoggedIn(), performLogin)

	loginPayload := getLoginPayload()
	req, _ := http.NewRequest("POST", "/u/login", strings.NewReader(loginPayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(loginPayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	// Test that the page title is "Successful Login"
	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Successful Login</title>") < 0 {
		t.Fail()
	}
}


func getLoginPayload() string {
	params := url.Values{}
	params.Add("username", "user1")
	params.Add("password", "pass1")

	return params.Encode()
}

func getRegistrationPOSTPayload() string {
	params := url.Values{}
	params.Add("username", "u1")
	params.Add("password", "p1")

	return params.Encode()
}