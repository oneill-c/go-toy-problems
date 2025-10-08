// Title: Customer Order Deduplication API

// Endpoint
//   POST /dedupe
//   Content-Type: application/json
//   Accept: application/json

// Request Body
// {
//   "systemA": [Order, ...],
//   "systemB": [Order, ...]
// }

// Order
// {
//   "id": string,
//   "customer_name": string,
//   "email": string,
//   "amount": number
// }

// Normalization & Rules
// 1) Normalize
//    - email: trim + lowercase
//    - customer_name: trim outer spaces; collapse internal runs of spaces to a single space
// 2) Validate emails: must contain “@” and “.” (use a simple check like `.+@.+\..+`); if invalid → skip record
// 3) Deduplicate by normalized email
//    - Keep the record with the higher amount
//    - Tie-breaker on equal amount: keep the lexicographically smaller `id`
// 4) Sort output alphabetically by `customer_name` (case-insensitive)
// 5) Return the merged, cleaned list as JSON array

// Responses
// - 200 OK
//   Body: [Order, ...]  // cleaned records (with normalized fields) sorted by customer_name
// - 400 Bad Request
//   - Invalid JSON shape, wrong types, or arrays missing
// - 415 Unsupported Media Type
//   - If `Content-Type` is not `application/json`
// - 413 Payload Too Large (optional, if enforcing a limit)

// Example Request
// POST /dedupe
// Content-Type: application/json

// {
//   "systemA": [
//     {"id":"a1","customer_name":"Ada Lovelace","email":"ada@Example.com","amount":49.99},
//     {"id":"a2","customer_name":"Alan Turing","email":"alan.turing@org","amount":29.50}
//   ],
//   "systemB": [
//     {"id":"b1","customer_name":"ALAN  TURING","email":"  Alan.Turing@ORG ","amount":34.00},
//     {"id":"b2","customer_name":"Grace Hopper","email":"grace.hopper@navy.mil","amount":99.99}
//   ]
// }

// Example 200 Response
// [
//   {"id":"a1","customer_name":"Ada Lovelace","email":"ada@example.com","amount":49.99},
//   {"id":"b1","customer_name":"Alan Turing","email":"alan.turing@org","amount":34.00},
//   {"id":"b2","customer_name":"Grace Hopper","email":"grace.hopper@navy.mil","amount":99.99}
// ]

// Notes
// - Sorting is by full `customer_name` (case-insensitive), not just the first character.
// - All fields in the response are the normalized versions.
// - If both arrays are absent or empty, return `[]` with 200.
package main

import (
	"encoding/json"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

type Order struct {
	ID           string  `json:"id"`
	CustomerName string  `json:"customer_name"`
	Email        string  `json:"email"`
	Amount       float64 `json:"amount"`
}

type ReqBody struct {
	SystemA []Order `json:"system_a"`
	SystemB []Order `json:"system_b"`
}

func NormalizeOrder(order Order) Order {
	return Order{
		ID:           order.ID,
		Amount:       order.Amount,
		Email:        NormalizeEmail(order.Email),
		CustomerName: NormalizeName(order.CustomerName),
	}
}

func NormalizeName(s string) string {
	return strings.Trim(s, " ")
}

func NormalizeEmail(s string) string {
	return strings.Trim(strings.ToLower(s), " ")
}

func ValidateEmail(s string) bool {
	re := regexp.MustCompile(`.+@.+\..+`)
	return re.MatchString(s)
}

func handleOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var rq ReqBody
	if err := json.NewDecoder(r.Body).Decode(&rq); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	all := append(append([]Order{}, rq.SystemA...), rq.SystemB...)
	byEmail := make(map[string]Order)

	for _, o := range all {
		// 1. Normalize email, customer name
		no := NormalizeOrder(o)

		// 2. skip record if invalid
		if !ValidateEmail(no.Email) {
			continue
		}

		// 3. Dedupe by email
		existing, ok := byEmail[no.Email]

		// 3a. Only place record at byEmail[no.Email] if doesn't yet exist OR if no.Amount > existing.Amount
		if !ok || no.Amount > existing.Amount {
			byEmail[no.Email] = no
		}
	}

	// 4. sort alphabetically by customer name and return json
	out := make([]Order, 0, len(byEmail))
	for _, r := range byEmail {
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].CustomerName) < strings.ToLower(out[j].CustomerName)
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(out)
}

func main() {
	http.HandleFunc("/dedupe", handleOrders)
	http.ListenAndServe(":8080", nil)
}
