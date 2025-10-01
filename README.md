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
├── csv-to-struct/
│   ├── main.go
│   └── testdata/
│       └── users.csv
├── flatten-json/
│   └── main.go
├── bfs/
│   └── main.go
├── dfs/
│   └── main.go
└── README.md
```

---

## 🚀 Current Problems

### 1) Top Poster

- **Task**: Mock two endpoints (`/users` and `/posts`), fetch both, and return the user with the highest number of posts.
- **Concepts**: chi mux routing, `httptest` mock server, JSON decode, simple aggregation.

Run:

```bash
cd top-poster
go run main.go
```

Expected:

```
Top poster: Bob with 2 posts
```

---

### 2) CSV → Struct

- **Task**: Read a CSV file and parse each row into a strongly typed `User` struct.
- **Concepts**: CSV parsing, type conversions (`strconv`, `time.Parse`), building slices of structs.

Run:

```bash
cd csv-to-struct
go run main.go
```

**Roadmap (CSV → Struct):**

1. Refactor into `ParseUsers(io.Reader)` for testability
2. Table-driven tests for conversions and error handling
3. Row-level error collection
4. Configurable CSV options (delimiter, layout)

---

### 3) Flatten JSON

- **Task**: Flatten an arbitrarily nested JSON object/array into a flat `map[string]string` with joined keys.
- **Concepts**: recursion, type switching, handling `map[string]any` and `[]any`, string conversion.

Run:

```bash
cd flatten-json
go run main.go
```

**Example input:**

```json
{
  "user": {
    "id": 42,
    "name": "jim",
    "tags": ["eng", "guitar"],
    "prefs": { "darkMode": true }
  }
}
```

**Example output (separator="."):**

```
user.id=42
user.name=jim
user.tags.0=eng
user.tags.1=guitar
user.prefs.darkMode=true
```

**Roadmap (Flatten JSON):**

1. Add unit tests with multiple input cases
2. Support custom key joiner (dot, underscore, etc.)
3. Handle null values distinctly (`null` vs empty string)
4. Optionally return `map[string]any` for type safety, with string conversion helper
5. Benchmarks for large JSON payloads

---

### 4) BFS (Breadth-First Search)

- **Task**: Traverse a binary tree in level order and print node values.
- **Concepts**: queue via slice, iterative traversal.

Run:

```bash
cd bfs
go run main.go
```

Expected:

```
1 2 3 4 5
```

---

### 5) DFS (Depth-First Search)

- **Task**: Traverse a binary tree in preorder and print node values.
- **Concepts**: recursion, call stack depth-first traversal.

Run:

```bash
cd dfs
go run main.go
```

Expected:

```
1 2 4 5 3
```

---

## 🛠️ Requirements

- [Go 1.21+](https://go.dev/dl/)
- [chi router](https://github.com/go-chi/chi) (only needed for HTTP-based problems)

Install chi (if running _Top Poster_):

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

## 🎯 Roadmap (repo-wide)

- **Top Poster**: Retries + backoff → Pagination → Table-driven tests
- **CSV → Struct**: Refactor to `io.Reader` → Tests → Row-level errors → Options
- **Flatten JSON**: Add tests → Configurable joiner → Null handling → Type-safe map
- **BFS/DFS**: Add table-driven tests + iterative/recursive variants

---

## 📚 Purpose

This repo is a lightweight playground for practicing Go problem-solving skills, preparing for interviews, and building intuition for:

- HTTP + JSON
- Data structures & algorithms
- Error handling
- Testing best practices
- Parsing and data munging
