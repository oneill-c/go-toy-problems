# go-toy-problems

A collection of small Go coding challenges (“toy problems”) with simple, working solutions.  
Each problem is designed to be easy to run, extend, and learn from.

---

## 📂 Structure

Each problem lives in its own folder and contains:

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
│   ├── worker-pool-waitgroup/
│   │   └── main.go
│   └── context-aware-worker-pool/
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
├── time-and-retries/
│   └── main.go
├── time-and-retries-with-context/
│   └── main.go
├── string-manipulation/
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

### 5) Concurrency — Context-Aware Worker Pool

**Path:** `concurrency/context-aware-worker-pool/main.go`  
Implements a worker pool that gracefully stops processing jobs when a **context timeout or cancellation** occurs. Each worker processes jobs concurrently until the context is canceled, at which point unfinished work is abandoned and the system shuts down cleanly.

**Features:**

- Uses `context.WithTimeout` to limit runtime duration
- Demonstrates graceful shutdown and cancellation propagation
- Safely collects results using a `JobStore` protected by a mutex
- Adjustable dials for global timeout and job processing delay

**Concepts:** context cancellation, timeout propagation, safe concurrent writes, worker synchronization.

**Example output:**

```
worker 2 processed 3
worker 0 processed 1
worker 1 processed 2
context canceled: context deadline exceeded
collected 6 results before cancel: [2 4 6 8 10 12]
```

---

### 6) HMAC-Verified JSON Fetch with Retry and Backoff

**Path:** `networking/hmac-with-retry/main.go`  
Build a resilient HTTP client that fetches a JSON payload, verifies authenticity with **HMAC-SHA256**, parses it into a typed struct, and retries failed requests with **exponential backoff and jitter**.

**Concepts:** secure message verification, JSON parsing, retry with jitter, timeout handling, rate-limit logic.

---

### 7) In-Memory Users Database

**Path:** `in-memory-users-db/main.go`  
Implement an in-memory database to manage user records. It should support basic operations for importing and retrieving users.

**Requirements:**

- `import_users` — accepts an input request and inserts provided users into the database.
- `get_users` — returns all user records.
- `get_user_by_id` — returns a single record matching the provided ID.

Handle conflicting or duplicate records gracefully — for example, multiple entries with the same email but different phone numbers should trigger a validation error.

**Concepts:** in-memory data structures, deduplication, validation, simple data access patterns.

---

### 8) Customer Order Deduplication API

**Path:** `dedupe-api/main.go`  
Expose a REST endpoint **POST /dedupe** that merges two systems’ order lists into a single clean, deduplicated JSON array.

**Concepts:** input normalization, deduplication logic, JSON encoding, simple REST handlers.

---

### 9) BFS (Breadth-First Search)

Traverse a binary tree in level order and print node values.

---

### 10) DFS (Depth-First Search)

Traverse a binary tree in preorder and print node values.

---

### 11) Paginated API Fetch with Retry, Backoff, and Checkpointing

**Path:** `networking/sync-pagination/main.go`  
Implement a Go program that fetches event data from a paginated HTTP API, handles transient errors with retries, and supports resuming from a saved checkpoint between runs.

**Concepts:** pagination, retry logic, checkpointing, context cancellation.

---

### 12) Time & Retries

**Path:** `time-and-retries/main.go`  
Implements an exponential backoff retry mechanism with ±25% jitter.  
Each retry doubles the delay up to a maximum cap, then applies random jitter to prevent synchronized retry storms.

**Concepts:** `time.Duration`, exponential backoff, random jitter, rate control, and retry safety limits.

---

### 13) Time & Retries with Context

**Path:** `time-and-retries-with-context/main.go`  
Implements an exponential backoff retry mechanism with ±25% jitter **and context cancellation**.  
Each retry doubles the delay up to a maximum cap, then applies random jitter to prevent synchronized retry storms.  
Demonstrates the use of `context.Context` for graceful cancellation and timeouts during retries.

**Concepts:** `context.Context`, exponential backoff, random jitter, retry loops, graceful timeout handling.

---

### 14) String Manipulation — Email/Phone Normalization & Validation

**Path:** `string-manipulation/main.go`  
Normalize and validate user contact info, then emit cleaned users and summary stats.

**Concepts:** string normalization, validation, light regex, stats aggregation, reporting.

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

- Add tests for new concurrency problems
- Expand retry logic with metrics
- Introduce structured logging (`log/slog`)
- Explore gRPC, WebSocket, and streaming examples
- Add integration test harness for HTTP handlers

---

## 📚 Purpose

A lightweight Go playground for improving problem-solving, preparing for interviews, and building intuition for:

- HTTP + JSON
- Concurrency patterns
- Data structures & algorithms
- Error handling
- Context & cancellation patterns
- Testing best practices
- Secure API interactions
- Parsing and data munging
