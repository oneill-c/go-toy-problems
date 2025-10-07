// To execute Go code, please declare a func main() in a package "main"

/*
Overview

You are tasked with building a simple in-memory users database that simulates importing and retrieving user data.
The goal is to design clear, minimal functions that handle basic CRUD-style operations.

---

PART 1:

Implement an in-memory database to store user data.
At minimum, it must implement the following operations:

1. import_users — accepts an input request and inserts the provided users into the database.
2. get_users — returns all user records as an array.
3. get_user_by_id — accepts a user ID and returns the corresponding user record.

---

Example input:
input = {
  "data": [
    {
      "name": "Batman",
      "phone": "555-555-1111",
      "email": "batman@gmail.com"
    },
    {
      "name": "Robin",
      "phone": "555-555-2222",
      "email": "robin@gmail.com"
    }
  ]
}

Example class sketch (Python-style pseudocode):
class UsersDatabase:
  def __init__(self):
    self.users = []

  def import_users(self, users_request):
    print("Not implemented")

  def get_user_by_id(self, user_id):
    print("Not implemented")

  def get_users(self):
    print("Not implemented")

---

PART 2:

Now consider conflicting data scenarios.
Your implementation should detect and handle duplicate or conflicting records (e.g., same email with different phone numbers).

Example conflicting inputs:

conflicting_input1 = {
  "data": [
    {
      "name": "Batman",
      "phone": "555-555-1111",
      "email": "batman@gmail.com"
    }
  ]
}

conflicting_input2 = {
  "data": [
    {
      "name": "Batman",
      "phone": "555-555-4950",
      "email": "batman@gmail.com"
    },
    {
      "name": "Joker",
      "phone": "555-555-3333",
      "email": "joker@gmail.com"
    }
  ]
}

// Import the conflicting user inputs
db.import_users(conflicting_input1)
db.import_users(conflicting_input2) // should fail due to conflict
*/

package main

import (
	"errors"
	"fmt"
)

type User struct {
	Name string
	Phone string
	Email string
}

type UserStore struct {
	UserDB map[string]User
	seenEmails map[string]struct{}
}

func NewUserStore() *UserStore {
	return &UserStore{
		UserDB: make(map[string]User), // [Name]{ Name, Phone, Email }
		seenEmails: make(map[string]struct{}),
	}
}

func (us *UserStore) ImportUsers(usersRequest []User) error {
	if len(usersRequest) == 0 {
		return errors.New("no users to import")
	}

	for _, u := range usersRequest {
		if u.Name == "" || u.Email == "" || u.Phone == "" {
			// move on to next record, we might want to retain via something like []RowError for logging and metric purposes
			continue
		}

		if _, ok := us.seenEmails[u.Email]; ok {
			// this is just for demonstration purposes. We should not print from library code like this
			fmt.Printf("invalid email address: %s\n", u.Email)
			continue
		}
	
		// we would want to validate fields like Email and Phone here as well

		us.seenEmails[u.Email] = struct{}{}
		us.UserDB[u.Name] = u
	}
	return nil
}

func (us *UserStore) GetUsers() []User {
	var temp []User
	for _, v := range us.UserDB {
		temp = append(temp, v)
	}
	return temp
}

func (us *UserStore) GetUserById(userName string) (*User, bool) {
	if userName == "" {
		return nil, false
	}
	u, ok := us.UserDB[userName]
	if !ok {
		return nil, false
	}
	return &u, ok
}

func main() {
	userStore := NewUserStore()
	importData := []User{
		{ Name: "Batman", Phone: "555-555-1111", Email: "batman@gmail.com" },
		{ Name: "Robin", Phone: "555-555-2222", Email: "robin@gmail.com" },
	}

	// import users
	imErr1 := userStore.ImportUsers(importData)
	if imErr1 != nil {
		panic(imErr1)
	}

	// test GetUsers()
	fmt.Println(userStore.GetUsers())

	// test GetUserById
	fmt.Println(userStore.GetUserById("Batman"))
	fmt.Println(userStore.GetUserById("Robin"))
	fmt.Println(userStore.GetUserById(""))
	fmt.Println(userStore.GetUserById("Jim"))

	// test duplication handling
	conflictingInput1 := []User{
		{ Name: "Batman", Phone: "555-555-1111", Email: "batman@gmail.com" },
	}

	imErr2 := userStore.ImportUsers(conflictingInput1)
	if imErr2 != nil {
		panic(imErr2)
	}

	fmt.Println(userStore.GetUsers())

	conflictingInput2 := []User{
		{ Name: "Batman", Phone: "555-555-1111", Email: "batman@gmail.com" },
		{ Name: "Joker", Phone: "555-555-3333", Email: "joker@gmail.com" },
	}

	imErr3 := userStore.ImportUsers(conflictingInput2)
	if imErr3 != nil {
		panic(imErr3)
	}

	fmt.Println(userStore.GetUsers())
	fmt.Println(userStore.GetUserById("Joker"))
}