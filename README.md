# go-toy-problems

A collection of small Go coding challenges (“toy problems”) with simple, working solutions.  
Each problem is designed to be easy to run, extend, and learn from.

---

## 📂 Structure

- Each problem lives in its own folder.
- Every folder contains:
  - `main.go` → minimal runnable solution
  - (optional) `*_test.go` → table-driven tests
  - A short description in comments

Example:

```
go-toy-problems/
├── top-poster/
│   ├── main.go
│   └── main_test.go
├── bfs/
│   └── main.go
├── dfs/
│   └── main.go
└── README.md
```

---

## 🚀 Current Problems

### 1. Top Poster

- **Task**:  
  Mock two endpoints (`/users` and `/posts`), fetch both, and return the user with the highest number of posts.
- **Concepts covered**:
  - Using a mux (`chi`) for routing
  - Mocking HTTP servers with `httptest`
  - Fetching and decoding JSON
  - Counting & aggregating results

Run:

```bash
cd top-poster
go run main.go
```

Expected output:

```
Top poster: Bob with 2 posts
```

### 2. BFS (Breadth-First Search)

- **Task**:  
  Implement BFS traversal of a binary tree and print node values in level order.
- **Concepts covered**:
  - Queues in Go (using slices)
  - Iterative traversal
  - Exploring trees level by level

Run:

```bash
cd bfs
go run main.go
```

Expected output (example tree):

```
1 2 3 4 5
```

### 3. DFS (Depth-First Search)

- **Task**:  
  Implement DFS traversal of a binary tree and print node values in depth-first order.
- **Concepts covered**:
  - Recursion in Go
  - Preorder traversal
  - Exploring trees branch by branch

Run:

```bash
cd dfs
go run main.go
```

Expected output (example tree):

```
1 2 4 5 3
```

---

## 🛠️ Requirements

- [Go 1.21+](https://go.dev/dl/)
- [chi router](https://github.com/go-chi/chi) (only needed for HTTP-based problems)

Install chi:

```bash
go get github.com/go-chi/chi/v5
```

---

## ✅ Testing

Some problems include tests. Run all with:

```bash
go test ./...
```

---

## 🎯 Roadmap

Planned follow-ups (step-by-step complexity increase):

### Top Poster

1. **Retries + backoff**
2. **Pagination (fetch until completion)**
3. **Table-driven tests**

### BFS / DFS

- Variants with iterative + recursive approaches
- Unit tests for traversal correctness

---

## 📚 Purpose

This repo exists as a lightweight playground for practicing Go problem-solving skills, preparing for interviews, and building intuition for:

- HTTP + JSON
- Data structures
- Error handling
- Testing best practices
- Common algorithms (BFS, DFS)
