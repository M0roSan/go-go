package main

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that the length of the list of articles returned is the same as the length ofthe global variable holding a list
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
		v.ID != articleList[i].ID ||
		v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}

}

// Test the function that fetches an article by its ID
func TestGetArticleByID(t *testing.T){
	a, err := getArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}

// Test the function that creates a new article
func TestCreateNewArticle(t *testing.T){
	// get the original count of articles 
	originalLength := len(getAllArticles())

	// add another article
	title := "New test title"
	content := "New test content"
	a, err := createNewArticle(title, content)

	// get the new count of articles
	newLength := len(getAllArticles())

	if err != nil || newLength != originalLength + 1 ||
		a.Title != title || a.Content != content {
			t.Fail()
		}
}

