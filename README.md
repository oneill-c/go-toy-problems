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
│   ├── hmac-with-retry/
│   │   └── main.go
│   └── sync-pagination/
│       └── main.go
├── in-memory-users-db/
│   └── main.go
├── bfs/
│   └── main.go
├── dfs/
│   └── main.go
├── dedupe-api/
│   └── main.go
└── README.md
```

---

## 🚀 Current Problems

### 1) Top Poster

Mock two endpoints (`/users` and `/posts`), fetch both, and return the user with the highest number of posts.

**Concepts:** chi mux routing, `httptest` mock server, JSON decode, simple aggregation.

---

### 2) CSV → Struct

Read a CSV file and parse each row into a strongly typed `User` struct.

**Concepts:** CSV parsing, type conversions (`strconv`, `time.Parse`), building slices of structs.

---

### 3) Flatten JSON

Flatten an arbitrarily nested JSON object or array into a flat `map[string]string` with joined keys.

**Concepts:** recursion, type switching, handling `map[string]any` and `[]any`, string conversion.

---

### 4) Concurrency — Worker Pool (WaitGroup)

**Path:** `concurrency/worker-pool-waitgroup/main.go`  
Process a finite batch (e.g., 10k blockchain transactions) in parallel using a fixed number of workers; emit only valid transactions.

**Concepts:** bounded concurrency, producer/consumer channels, `sync.WaitGroup`, fan-out/fan-in, backpressure via buffered channels.

---

### 5) HMAC-Verified JSON Fetch with Retry and Backoff

**Path:** `networking/hmac-with-retry/main.go`  
Build a resilient HTTP client that fetches a JSON payload, verifies authenticity with **HMAC-SHA256**, parses it into a typed struct, and retries failed requests with **exponential backoff and jitter**.

**Concepts:** secure message verification, JSON parsing, retry with jitter, timeout handling, rate-limit logic.

---

### 6) In-Memory Users Database

**Path:** `in-memory-users-db/main.go`  
Implement an in-memory database to manage user records. It should support basic operations for importing and retrieving users.

**Requirements:**

1. `import_users` — accepts an input request and inserts provided users into the database.
2. `get_users` — returns all user records.
3. `get_user_by_id` — returns a single record matching the provided ID.

Handle conflicting or duplicate records gracefully — for example, multiple entries with the same email but different phone numbers should trigger a validation error.

**Concepts:** in-memory data structures, deduplication, validation, simple data access patterns.

---

### 7) Customer Order Deduplication API

**Path:** `dedupe-api/main.go`

Expose a REST endpoint **POST /dedupe** that merges two systems’ order lists into a single clean, deduplicated JSON array.

**Requirements:**

1. Normalize input data
   - `email`: trim + lowercase
   - `customer_name`: trim + collapse internal spaces
2. Validate emails using a simple regex (`.+@.+\..+`); skip invalid ones.
3. Deduplicate by normalized email
   - Keep record with highest `amount`
   - Tie-break on equal `amount` using lexicographically smaller `id`
4. Sort alphabetically by `customer_name` (case-insensitive).
5. Return the cleaned list as JSON.

**Concepts:** basic HTTP handler, input validation, normalization, deduplication logic, JSON encoding.

**Example Request:**

```json
POST /dedupe
{
  "systemA": [
    {"id":"a1","customer_name":"Ada Lovelace","email":"ada@Example.com","amount":49.99},
    {"id":"a2","customer_name":"Alan Turing","email":"alan.turing@org","amount":29.50}
  ],
  "systemB": [
    {"id":"b1","customer_name":"ALAN  TURING","email":"  Alan.Turing@ORG ","amount":34.00},
    {"id":"b2","customer_name":"Grace Hopper","email":"grace.hopper@navy.mil","amount":99.99}
  ]
}
```

**Example Response:**

```json
[
  {
    "id": "a1",
    "customer_name": "Ada Lovelace",
    "email": "ada@example.com",
    "amount": 49.99
  },
  {
    "id": "b1",
    "customer_name": "Alan Turing",
    "email": "alan.turing@org",
    "amount": 34.0
  },
  {
    "id": "b2",
    "customer_name": "Grace Hopper",
    "email": "grace.hopper@navy.mil",
    "amount": 99.99
  }
]
```

**Concepts covered:**

- Input normalization and validation
- Deduplication logic using maps
- JSON encoding/decoding
- Case-insensitive sorting

---

### 8) BFS (Breadth-First Search)

Traverse a binary tree in level order and print node values.

---

### 9) DFS (Depth-First Search)

Traverse a binary tree in preorder and print node values.

---

### 10) Paginated API Fetch with Retry, Backoff, and Checkpointing

**Path:** `networking/sync-pagination/main.go`  
Implement a Go program that fetches event data from a paginated HTTP API, handles transient errors with retries, and supports resuming from a saved checkpoint between runs.

**Concepts:** pagination, retry logic, checkpointing, context cancellation.

---

## 🛠️ Requirements

- [Go 1.21+](https://go.dev/dl/)
- [chi router](https://github.com/go-chi/chi) (optional, for HTTP problems)

---

## ✅ Testing

Run all tests:

```bash
go test ./...
```

---

## 🎯 Roadmap (repo-wide)

- Add tests for new API-based problems
- Introduce context cancellation patterns
- Add error logging with `log/slog`
- Explore gRPC and WebSocket examples
- Add integration test harness for HTTP handlers

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
