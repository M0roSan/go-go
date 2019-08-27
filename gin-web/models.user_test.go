package main

import "testing"

// Test the validity of different combinatoins of username/password
func TestUserValidity(t *testing.T){
	if !isUserValid("user1", "pass1"){
		t.Fail()
	}
	
	if isUserValid("user2", "pass1") {
		t.Fail()
	}

	if isUserValid("user1", "") {
		t.Fail()
	}

	if isUserValid("", "user1"){
		t.Fail()
	}

	if isUserValid("User1", "pass1") {
		t.Fail()
	}
}

// Test if a new user can be registered with valid username/password
func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u, err := registerNewUser("newuser", "newpass")

	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()
}

// Test that a new user cannot be registered with invalid username/password
func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	// Existing user
	u, err := registerNewUser("user1", "pass1")
	if err == nil || u != nil {
		t.Fail()
	}

	// Register with blank password
	u, err = registerNewUser("newuser", "")
	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

// Test the functoin that checks for username availability
func TestUsernameAvailablity(t *testing.T) {
	saveLists()

	// new username should be avaiable
	if !isUsernameAvailable("newuser") {
		t.Fail()
	}

	// existing username
	if isUsernameAvailable("user1") {
		t.Fail()
	}

	// register a new user
	registerNewUser("newuser", "newpass")

	// this newly registered username shouldn't be available
	if isUsernameAvailable("newuser") {
		t.Fail()
	}

	restoreLists()
}