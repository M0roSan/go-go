package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
  
	"github.com/gin-gonic/gin"
)

var tmpArticleList []article
var tmpUserList []user

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper functoin to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		r.Use(setUserStatus())
	}
	return r
}

// Helper function to process a request and test its response
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool){
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request
	r.ServeHTTP(w, req)

	if !f(w){
		t.Fail()
	}
}

// Helper function to test middleware function
func testMiddlewareRequest(t *testing.T, r *gin.Engine, expectedHTTPCode int) {
	// Creates a request to send to the above route
	req, _ := http.NewRequest("GET", "/", nil)

	// Process the request and test the response
	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		return w.Code == expectedHTTPCode
	})
}

// This function is used to store the main lists into the tempirary one for testing
func saveLists() {
	tmpArticleList = articleList
	tmpUserList = userList
}

// This function is used to restore the main lists from the tmeporary one
func restoreLists() {
	articleList = tmpArticleList
	userList = tmpUserList
}
