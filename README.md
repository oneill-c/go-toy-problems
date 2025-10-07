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
├── concurrency/
│   └── worker-pool-waitgroup/
│       └── main.go
├── networking/
│   └── hmac-with-retry/
│       └── main.go
├── bfs/
│   └── main.go
├── dfs/
│   └── main.go
└── README.md
```

---

## 🚀 Current Problems

### 1) Top Poster

Mock two endpoints (`/users` and `/posts`), fetch both, and return the user with the highest number of posts.

**Concepts:** chi mux routing, `httptest` mock server, JSON decode, simple aggregation.

Run:

```bash
cd top-poster
go run main.go
```

Expected output:

```
Top poster: Bob with 2 posts
```

---

### 2) CSV → Struct

Read a CSV file and parse each row into a strongly typed `User` struct.

**Concepts:** CSV parsing, type conversions (`strconv`, `time.Parse`), building slices of structs.

Run:

```bash
cd csv-to-struct
go run main.go
```

**Roadmap:**

1. Refactor into `ParseUsers(io.Reader)` for testability
2. Table-driven tests for conversions and error handling
3. Row-level error collection
4. Configurable CSV options (delimiter, layout)

---

### 3) Flatten JSON

Flatten an arbitrarily nested JSON object or array into a flat `map[string]string` with joined keys.

**Concepts:** recursion, type switching, handling `map[string]any` and `[]any`, string conversion.

Run:

```bash
cd flatten-json
go run main.go
```

**Example output (separator="."):**

```
user.id=42
user.name=jim
user.tags.0=eng
user.tags.1=guitar
user.prefs.darkMode=true
```

**Roadmap:**

1. Add unit tests with multiple input cases
2. Support custom key joiner (dot, underscore, etc.)
3. Handle null values distinctly (`null` vs empty string)
4. Optionally return `map[string]any` for type safety
5. Benchmarks for large JSON payloads

---

### 4) Concurrency — Worker Pool (WaitGroup)

**Path:** `concurrency/worker-pool-waitgroup/main.go`  
Process a finite batch (e.g., 10k blockchain transactions) in parallel using a fixed number of workers; emit only valid transactions.

**Concepts:** bounded concurrency, producer/consumer channels, `sync.WaitGroup`, fan-out/fan-in, backpressure via buffered channels.

Run:

```bash
cd concurrency/worker-pool-waitgroup
go run main.go
```

**Roadmap:**

1. Make `workerCount` configurable
2. Collect and return valid transactions
3. Add context cancellation & deadlines
4. Add table-driven tests (happy path, all invalid, mixed)
5. Add error metrics
6. Replace `WaitGroup` with `errgroup.Group`

---

### 5) HMAC-Verified JSON Fetch with Retry and Backoff

**Path:** `networking/hmac-with-retry/main.go`  
Build a resilient HTTP client that:

1. Fetches a JSON payload from an HTTP endpoint.
2. Verifies authenticity using **HMAC-SHA256** and a pre-shared secret.
3. Parses verified JSON into a typed struct.
4. Retries failed requests with **exponential backoff + jitter** (up to 5 retries, 10s timeout).

**Example payload:**

```json
{
  "data": {
    "event_id": "abc123",
    "timestamp": "2025-10-06T15:04:05Z",
    "user_id": "user_456",
    "action": "login"
  },
  "signature": "abc123deadbeef..."
}
```

**Concepts:**

- Secure message verification via HMAC-SHA256
- JSON decoding into typed structs
- Timeout and retry with backoff + jitter
- Error handling for 4xx/5xx and 429 (rate limiting)

Run:

```bash
cd networking/hmac-with-retry
go run main.go
```

**Roadmap:**

1. Use env/flag for secret key
2. Add context cancellation & deadlines
3. Add test server for signature validation
4. Add structured logging and metrics
5. Extend with signed `POST` request support
6. Use `errgroup.Group` for concurrent fetch/verify

---

### 6) BFS (Breadth-First Search)

Traverse a binary tree in level order and print node values.

**Concepts:** queue via slice, iterative traversal.

Run:

```bash
cd bfs
go run main.go
```

Expected output:

```
1 2 3 4 5
```

---

### 7) DFS (Depth-First Search)

Traverse a binary tree in preorder and print node values.

**Concepts:** recursion, call stack depth-first traversal.

Run:

```bash
cd dfs
go run main.go
```

Expected output:

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

Run all tests:

```bash
go test ./...
```

---

## 🎯 Roadmap (repo-wide)

- **Top Poster** → Retries + Pagination + Tests
- **CSV → Struct** → Reader refactor + Tests + Error reporting
- **Flatten JSON** → Tests + Configurable joiner + Null handling
- **Worker Pool** → Context + Tests + Metrics
- **HMAC with Retry** → Context, Tests, Logging, Metrics
- **BFS/DFS** → Table-driven tests + Iterative variants

---

## 📚 Purpose

A lightweight Go playground for improving problem-solving, preparing for interviews, and building intuition for:

- HTTP + JSON
- Concurrency patterns
- Data structures & algorithms
- Error handling
- Testing best practices
- Secure API interactions
- Parsing and data munging
