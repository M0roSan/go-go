package main

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

// Test that a GET request to the home page returns the home page with HTTP code 200 for unauthenticated user
func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter(true)

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		// Test that the http status code is 200
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Home Page"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0
	
		return statusOK && pageOK
	})


}

// Test that a GET request to the home page returns the home page with HTTP code 200 for authenticated user
func TestShowIndexPageAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Home Page</title>") < 0 {
		t.Fail()
	}

}

// Test that a GET request an article page returns the article page with HTTP code 200 for an unauthenticated user
func TestArticleUnauthenticated(t *testing.T){
	r := getRouter(true)

	r.GET("/article/view/:article_id", getArticle)

	// Create a request to send to its definition in the routes file
	req, _ := http.NewRequest("GET", "/article/view/1", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		// Test that the page title is "Article 1"
		p, err := ioutil.ReadAll(w.Body)
		pageOK := err == nil && strings.Index(string(p), "<title>Article 1</title>") > 0

		return statusOK && pageOK
	})
}

// Test that a GET request to an article page returns the article page with HTTP code 200 for an authenticated user
func TestArticleAuthenticated(t *testing.T){
	w := httptest.NewRecorder()
	r := getRouter(true)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.GET("/article/view/:article_id", getArticle)

	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Article 1</title>") < 0 {
		t.Fail()
	}
}

// Test that a GET request to the home page returns the list of articles in JSON when the Accept header is set to application/json
func TestArticleListJSON(t *testing.T){
	r := getRouter(true)
	r.GET("/", showIndexPage)
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Add("Accept", "application/json")

	testHTTPResponse(t,r,req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var articles []article
		err = json.Unmarshal(p, &articles)
		return err == nil && len(articles) >= 2 && statusOK
	})
}

// Test that a GET request to an article page returns the article in XML when the Accept header is set to application/xml
func TestArticleListXML(t *testing.T){
	r := getRouter(true)
	r.GET("/article/view/:article_id", getArticle)
	req, _ := http.NewRequest("GET", "/article/view/1", nil)
	req.Header.Add("Accept", "application/xml")

	testHTTPResponse(t,r,req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK
		p, err := ioutil.ReadAll(w.Body)
		if err != nil {
			return false
		}
		var a article
		err = xml.Unmarshal(p, &a)
		return err == nil && a.ID == 1 && len(a.Title) >= 0 && statusOK
	})
}

// Test that a GET request to the article creatoin page returns the article reation page with the HTTP code 200 for authenticated user
func TestArticleCreationPageAuthenticated(t *testing.T){
	w := httptest.NewRecorder()
	r := getRouter(true)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.GET("/article/create", ensureLoggedIn(), showArticleCreationPage)
	req, _ := http.NewRequest("GET", "/article/create", nil)
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Create New Article</title>") < 0 {
		t.Fail()
	}
}

// Test that a GET request to the article creatoin page returns the article reation page with the HTTP code 401 for unauthenticated user
func TestArticleCreationPageUnauthenticated(t *testing.T){
	r := getRouter(true)
	r.GET("/article/create", ensureLoggedIn(), showArticleCreationPage)
	req, _ := http.NewRequest("GET", "/article/create", nil)
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusUnauthorized
	})
}


// Test that a POST request to create an article returns HTTP code 200 for authenticated user
func TestArticleCreationAuthenticated(t *testing.T) {
	w := httptest.NewRecorder()
	r := getRouter(true)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: "123"})
	r.POST("/article/create", ensureLoggedIn(), createArticle)

	articlePayload := getArticlePOSTPayload()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header = http.Header{"Cookie": w.HeaderMap["Set-Cookie"]}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fail()
	}

	p, err := ioutil.ReadAll(w.Body)
	if err != nil || strings.Index(string(p), "<title>Submission Successful</title>") < 0 {
		t.Fail()
	}
}

// Test that a POST request to create an article return HTTP code 401 for unauthenticated user
func TestArticleCreationUnauthenticated(t *testing.T) {
	r := getRouter(true)
	r.POST("/article/create", ensureLoggedIn(), createArticle)

	articlePayload := getArticlePOSTPayload()
	req, _ := http.NewRequest("POST", "/article/create", strings.NewReader(articlePayload))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(articlePayload)))	

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == http.StatusUnauthorized
	})

}
func getArticlePOSTPayload() string {
	params := url.Values{}
	params.Add("title", "Test Article Title")
	params.Add("content", "Test Article Content")

	return params.Encode()
}