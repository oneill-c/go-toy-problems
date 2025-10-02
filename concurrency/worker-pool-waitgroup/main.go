package main

import (
	"fmt"
	"sync"
)

/**
*
*
* Here we are simulating working with a finite batch of independent tasks (e.g. validate 10k trades).
* This is an example of bounded concurrency with a clear "done" point, as opposed to streaming
*
*
* Problem:
*
* You’re given a list of 10,000 blockchain transactions that need to be validated (syntax, digital signature, and gas estimates).
* Each validation is independent. Build a system that processes these validations in parallel using a fixed number of workers, then outputs a list of only the valid transactions.
*
* 	1.	Check symbol is supported
*			•	Only allow a fixed whitelist: e.g. "BTC", "ETH", "SOL".
*		2.	Check price > 0
*			•	Transactions/trades must have a positive price.
*		3.	Check volume > 0
*			•	Volume must also be positive.
*		4.	Check gas estimate is reasonable
*			•	Require gas between 21000 and 1,000,000.
*			•	Just a range check — not a real blockchain query.
*		5.	Check signature is present
**/

type Transaction struct {
	ID int
	Symbol string
	Price float64
	Volume float64
	GasEstimate float64
	Signature string
}

type ValidTransaction = Transaction

// constants / validation points
const (
	workerCount = 5
	reasonableGasPriceMin = 21000.0
	reasonableGasPriceMax = 1_000_000.0
)
var supportedSymbolWhitelist = map[string]struct{}{"BTC": {}, "ETH": {}, "SOL": {}}

func validate(tx Transaction) (bool, error) {
	
	// check within price range
	if tx.Price <= 0 {
		return false, fmt.Errorf("price must be > 0, tx=%+v", tx)
	}

	// check within volume range
	if tx.Volume <= 0 {
		return false, fmt.Errorf("volume must be > 0, tx=%+v", tx)
	}

	// check within reasonable gas price min & max
	if tx.GasEstimate < reasonableGasPriceMin || tx.GasEstimate > reasonableGasPriceMax {
		return false, fmt.Errorf("gas estimate not within acceptable range, tx=%+v", tx)
	}

	// check if symbol in whitelist
	if _, ok := supportedSymbolWhitelist[tx.Symbol]; !ok {
		return false, fmt.Errorf("symbol not supported, tx=%+v", tx)
	}

	// check signature
	if tx.Signature == "" {
		return false, fmt.Errorf("invalid signature, tx=%+v", tx)
	}

	return true, nil
}

func worker(txs <-chan Transaction, out chan<- Transaction, wg *sync.WaitGroup) {
	defer wg.Done()
	for tx := range txs {
		if valid, err := validate(tx); valid {
			out <- tx
		} else {
			fmt.Println(err)
		}
	}
}

func main() {
	
	// producer channel
	jobs := make(chan Transaction, 128)

	// consumer channel
	out := make(chan ValidTransaction, len(transactions))

	// set up waitgroup to wait for n=workerCount workers
	var wg sync.WaitGroup
	wg.Add(workerCount)

	// fire off n=workerCount goroutines to process workload
	for i := 0; i < workerCount; i++ {
		go worker(jobs, out, &wg)
	}

	// queue up jobs
	for _, tx := range transactions {
		jobs <- tx
	}
	close(jobs)

	// waitgroup wait until processed and close out channel
	go func() {
		wg.Wait()
		close(out)
	}()

	// print result
	count := 0
	for range out {
		count++
	}
	fmt.Printf("Processed %d valid transactions of a possible %d\n", count, len(transactions))
}

// testdata
// transactions
var transactions = []Transaction{
	// ✅ valid
	{ID: 1, Symbol: "BTC", Price: 30000, Volume: 0.5, GasEstimate: 50000, Signature: "sig1"},
	{ID: 2, Symbol: "ETH", Price: 2000, Volume: 1.0, GasEstimate: 80000, Signature: "sig2"},
	{ID: 3, Symbol: "SOL", Price: 100, Volume: 10, GasEstimate: 60000, Signature: "sig3"},
	{ID: 4, Symbol: "ETH", Price: 1800, Volume: 2.0, GasEstimate: 70000, Signature: "sig4"},
	{ID: 5, Symbol: "BTC", Price: 31000, Volume: 0.8, GasEstimate: 90000, Signature: "sig5"},

	// ❌ invalid: negative price
	{ID: 6, Symbol: "ETH", Price: -2000, Volume: 1.0, GasEstimate: 60000, Signature: "sig6"},

	// ❌ invalid: negative volume
	{ID: 7, Symbol: "BTC", Price: 32000, Volume: -0.5, GasEstimate: 60000, Signature: "sig7"},

	// ❌ invalid: gas too low
	{ID: 8, Symbol: "BTC", Price: 30000, Volume: 1.0, GasEstimate: 1000, Signature: "sig8"},

	// ❌ invalid: gas too high
	{ID: 9, Symbol: "ETH", Price: 2000, Volume: 1.5, GasEstimate: 999999999, Signature: "sig9"},

	// ❌ invalid: symbol not in whitelist
	{ID: 10, Symbol: "DOGE", Price: 0.3, Volume: 1000, GasEstimate: 60000, Signature: "sig10"},

	// ❌ invalid: empty signature
	{ID: 11, Symbol: "SOL", Price: 150, Volume: 20, GasEstimate: 55000, Signature: ""},

	// ✅ valid
	{ID: 12, Symbol: "BTC", Price: 29500, Volume: 0.9, GasEstimate: 60000, Signature: "sig12"},
	{ID: 13, Symbol: "ETH", Price: 2200, Volume: 3.0, GasEstimate: 70000, Signature: "sig13"},
	{ID: 14, Symbol: "SOL", Price: 95, Volume: 15, GasEstimate: 62000, Signature: "sig14"},
	{ID: 15, Symbol: "BTC", Price: 33000, Volume: 0.2, GasEstimate: 64000, Signature: "sig15"},
	{ID: 16, Symbol: "ETH", Price: 2050, Volume: 2.5, GasEstimate: 58000, Signature: "sig16"},

	// ❌ invalid: price and volume both negative
	{ID: 17, Symbol: "BTC", Price: -100, Volume: -1, GasEstimate: 60000, Signature: "sig17"},

	// ❌ invalid: missing signature + bad symbol
	{ID: 18, Symbol: "XRP", Price: 0.5, Volume: 500, GasEstimate: 60000, Signature: ""},

	// ❌ invalid: negative volume + gas too low
	{ID: 19, Symbol: "ETH", Price: 2500, Volume: -5, GasEstimate: 2000, Signature: "sig19"},

	// ✅ valid
	{ID: 20, Symbol: "SOL", Price: 120, Volume: 8, GasEstimate: 59000, Signature: "sig20"},
}