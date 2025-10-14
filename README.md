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
│   ├── context-aware-worker-pool/
│   │   └── main.go
│   └── context-aware-token-bucket/
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

**Concepts:** context cancellation, timeout propagation, safe concurrent writes, worker synchronization.

---

### 6) Concurrency — Context-Aware Token Bucket

**Path:** `concurrency/context-aware-token-bucket/main.go`  
Implements a **rate limiter** combined with a **context-aware worker pool**.  
Workers process jobs only when a token is available, simulating controlled throughput (e.g., API rate limits).

**Features:**

- Token bucket refills `N` times per second (`tokensPerSecond`)
- Burst capacity defines the maximum tokens that can accumulate
- Integrated with context timeout for graceful shutdown
- Adjustable dials for bucket size, job count, and latency simulation

**Concepts:** token bucket algorithm, rate limiting, `context.Context`, worker coordination, graceful cancellation.

**Example output:**

```
worker 1 processed 2
worker 3 processed 5
worker 0 processed 1
context canceled: context deadline exceeded
collected 8 results before cancel: [2 4 6 8 10 12 14 16]
```

---

### 7) HMAC-Verified JSON Fetch with Retry and Backoff

**Path:** `networking/hmac-with-retry/main.go`  
Build a resilient HTTP client that fetches a JSON payload, verifies authenticity with **HMAC-SHA256**, parses it into a typed struct, and retries failed requests with **exponential backoff and jitter**.

**Concepts:** secure message verification, JSON parsing, retry with jitter, timeout handling, rate-limit logic.

---

### 8) In-Memory Users Database

**Path:** `in-memory-users-db/main.go`  
Implement an in-memory database to manage user records. It should support basic operations for importing and retrieving users.

**Concepts:** in-memory data structures, deduplication, validation, simple data access patterns.

---

### 9) Customer Order Deduplication API

**Path:** `dedupe-api/main.go`  
Expose a REST endpoint **POST /dedupe** that merges two systems’ order lists into a single clean, deduplicated JSON array.

**Concepts:** input normalization, deduplication logic, JSON encoding, simple REST handlers.

---

### 10) BFS (Breadth-First Search)

Traverse a binary tree in level order and print node values.

---

### 11) DFS (Depth-First Search)

Traverse a binary tree in preorder and print node values.

---

### 12) Paginated API Fetch with Retry, Backoff, and Checkpointing

**Path:** `networking/sync-pagination/main.go`  
Fetch paginated event data from an API, retry transient failures, and resume from a saved checkpoint.

**Concepts:** pagination, checkpointing, backoff, resumable syncs.

---

### 13) Time & Retries

**Path:** `time-and-retries/main.go`  
Exponential backoff retry mechanism with ±25% jitter.

**Concepts:** `time.Duration`, backoff, random jitter, retry limits.

---

### 14) Time & Retries with Context

**Path:** `time-and-retries-with-context/main.go`  
Same as above, but **context-aware** for graceful timeout/cancel handling.

**Concepts:** context cancellation, backoff, timing control.

---

### 15) String Manipulation — Email/Phone Normalization & Validation

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

- Add more concurrency control examples (token buckets, semaphores, rate limiters)
- Introduce structured logging (`log/slog`)
- Expand retry logic with metrics and tracing
- Explore gRPC, WebSocket, and streaming patterns
- Add integration tests for HTTP and concurrency cases

---

## 📚 Purpose

A lightweight Go playground for improving problem-solving, preparing for interviews, and building intuition for:

- HTTP + JSON
- Concurrency patterns
- Rate limiting and retry logic
- Data structures & algorithms
- Error handling and resilience
- Context & cancellation patterns
- Secure API interactions
- Data validation and normalization
