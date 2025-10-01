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

- **Task**: Read a CSV file and parse each row into a strongly typed `User` struct:
  ```go
  type User struct {
      ID     int
      Email  string
      Active bool
      DOB    time.Time
      Score  float64
  }
  ```
- **Notes**: Minimal parser using `encoding/csv`, manual field conversions (`strconv`, `time.Parse`).

Run:

```bash
cd csv-to-struct
go run main.go
```

Default input: `testdata/users.csv`

**Example output (truncated):**

```
CSV parsing complete.
Read N records

{ID:1 Email:alice@example.com Active:true DOB:1990-12-01 00:00:00 +0000 UTC Score:98.5}
...
```

**Roadmap (CSV → Struct):**

1. **Refactor for testability**
   - Extract a pure function `ParseUsers(r io.Reader) ([]User, error)`
   - Make date layout a parameter (e.g., `ParseUsers(r, layout string)`)
2. **Table-driven tests**
   - Happy path, bad int/bool/float, bad date, empty file, missing column
   - Golden-file test using `testdata/users.csv`
3. **Error handling**
   - Collect row-level errors with line numbers; continue on non-fatal rows
   - Return an aggregate error type (e.g., `type RowError struct { Line int; Err error }`)
4. **CSV options**
   - Support custom delimiter, comment char, trimming space, `FieldsPerRecord`
5. **I/O ergonomics**
   - Support reading from `os.Stdin` and a `-file` flag
   - Add `-date-layout` flag (default `2006-01-02`)
6. **Performance & safety**
   - Benchmarks with `testing.B`
   - Preallocate slice when `r.FieldsPerRecord` is known
7. **(Stretch)** Generic/reflect-based mapper or struct tags for column mapping

---

### 3) BFS (Breadth-First Search)

- **Task**: Traverse a binary tree in level order and print node values.
- **Concepts**: queue via slice, iterative traversal.

Run:

```bash
cd bfs
go run main.go
```

Expected (example tree):

```
1 2 3 4 5
```

---

### 4) DFS (Depth-First Search)

- **Task**: Traverse a binary tree in preorder and print node values.
- **Concepts**: recursion, call stack depth-first traversal.

Run:

```bash
cd dfs
go run main.go
```

Expected (example tree):

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

- **Top Poster**: Retries + backoff → Pagination (until completion) → Table-driven tests
- **CSV → Struct**: Refactor for `io.Reader` → Table-driven tests → Row-level error reporting → CSV options
- **BFS/DFS**: Add tests and iterative/recursive variants

---

## 📚 Purpose

This repo is a lightweight playground for practicing Go problem-solving skills, preparing for interviews, and building intuition for:

- HTTP + JSON
- Data structures & algorithms
- Error handling
- Testing best practices
- Parsing and data munging
