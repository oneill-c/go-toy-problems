package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

func main() {
	// 1. Setup data structures for users + posts
	type User struct {
		ID string `json:"id"`
		Name string `json:"name"`
	}

	type Post struct {
		ID string `json:"id"`
		Title string `json:"title"`
		UserID string `json:"userId"`
	}

	users := []User{
		{ ID: "u123", Name: "Jim" },
		{ ID: "u456", Name: "Joe" },
		{ ID: "u789", Name: "George" },
		{ ID: "u1011", Name: "Sally" },
	}

	posts := []Post{
		{ ID: "p123", Title: "post 1", UserID: "u123" },
		{ ID: "p456", Title: "post 2", UserID: "u123" },
		{ ID: "p789", Title: "post 3", UserID: "u456" },
		{ ID: "p1011", Title: "post 4", UserID: "u789" },
		{ ID: "p1213", Title: "post 5", UserID: "u1011" },
		{ ID: "p1415", Title: "post 6", UserID: "u1011" },
		{ ID: "p1718", Title: "post 7", UserID: "u123" },
		{ ID: "p1920", Title: "post 8", UserID: "u123" },
	}

	router := chi.NewRouter()

	// 2. Define GET users endpoint
	router.Get("/users", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(users)
	})
	// 3. Define GET posts endpoint
	router.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(posts)
	})

	// 4. Start test HTTP server
	srv := httptest.NewServer(router)
	defer srv.Close()

	// 5. Call /users endpoint
	var fetchedUsers []User
	respUsers, err := http.Get(fmt.Sprintf("%s%s", srv.URL, "/users"))
	if err != nil {
		panic(err)
	}
	defer respUsers.Body.Close()
	json.NewDecoder(respUsers.Body).Decode(&fetchedUsers)

	// 6. Call /posts endpoint
	var fetchedPosts []Post
	respPosts, err := http.Get(fmt.Sprintf("%s%s", srv.URL, "/posts"))
	if err != nil {
		panic(err)
	}
	defer respPosts.Body.Close()
	json.NewDecoder(respPosts.Body).Decode(&fetchedPosts)

	// 7. Count posts with userID
	counts := make(map[string]int)
	for _, p := range fetchedPosts {
		counts[p.UserID]++
	}

	// 8. Get top poster
	var maxCount int
	var topPoster User
	
	for _, u := range users {
		if counts[u.ID] > maxCount {
			maxCount = counts[u.ID]
			topPoster = u
		}
	}
	fmt.Printf("Top poster: %s with %d posts\n", topPoster.Name, counts[topPoster.ID])
}