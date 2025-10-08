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

**Roadmap (In-Memory Users DB):**

1. Implement error handling for duplicate/conflicting users.
2. Add table-driven tests for imports and lookups.
3. Introduce optional persistence to disk (JSON file).
4. Add filtering/sorting (by email, name, etc.).
5. (Stretch) Add a REST or gRPC layer to expose the API.

---

### 7) BFS (Breadth-First Search)

Traverse a binary tree in level order and print node values.

---

### 8) DFS (Depth-First Search)

Traverse a binary tree in preorder and print node values.

---

### 9) Paginated API Fetch with Retry, Backoff, and Checkpointing

**Path:** `networking/sync-pagination/main.go`  
Implement a Go program that fetches event data from a paginated HTTP API, handles transient errors with retries, and supports resuming from a saved checkpoint between runs.

**Requirements:**

- Fetch all pages until `next_cursor` is `null`
- Count total and unique events
- Save a `.checkpoint` file to resume progress
- Handle retries with exponential backoff and jitter
- Abort after 3 failed attempts or unrecoverable errors

**Example payload:**

```json
{
  "data": [
    {
      "id": "evt_123",
      "wallet_id": "wallet_abc",
      "kind": "deposit",
      "ts": "2025-10-06T15:04:05Z"
    }
  ],
  "next_cursor": "cursor_456"
}
```

**Example output:**

```json
{
  "processed": 1532,
  "unique_events": 1500,
  "last_cursor": "cursor_456"
}
```

**Concepts covered:**

- Pagination and cursor-based iteration
- Retry with exponential backoff and jitter
- Checkpointing and atomic file writes
- Idempotent design for resumable processes

**Roadmap (Sync Pagination):**

1. Add context-based cancellation for graceful shutdown.
2. Support concurrent page fetching (while preserving order).
3. Add testable interfaces (`fetchPage`, `saveCheckpointAtomic`).
4. Mock HTTP responses for integration testing.
5. Add structured logging (`log/slog`).

---

## 🛠️ Requirements

- [Go 1.21+](https://go.dev/dl/)
- [chi router](https://github.com/go-chi/chi) (only needed for HTTP-based problems)

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
- **In-Memory DB** → Conflict detection, Persistence, Filtering
- **Sync Pagination** → Context, Concurrency, Logging, Testing
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
