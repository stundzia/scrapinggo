package performance

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"

	"github.com/stundzia/scrapinggo/fetch"
	"github.com/stundzia/scrapinggo/parse"
)

func TestBashOrgWithParsing(parallel int) {
	okCount, errCount := &atomic.Uint64{}, &atomic.Uint64{}
	var quotes []string
	sem := semaphore.NewWeighted(int64(parallel))
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(422)
	mux := sync.Mutex{}
	testStart := time.Now()

	go func() {
		for i := 1; i < 423; i++ {
			_ = sem.Acquire(ctx, 1)
			go func(page int) {
				r, err := fetch.ResponseSimple("http://bash.org/?browse&p=" + strconv.Itoa(page))
				sem.Release(1)
				wg.Done()
				if err != nil || r.StatusCode != 200 {
					errCount.Add(1)
					return
				}
				okCount.Add(1)
				qs, err := parse.BashOrgQuotesGQ(r.Body)
				defer r.Body.Close()
				mux.Lock()
				quotes = append(quotes, qs...)
				mux.Unlock()
			}(i)
		}
	}()

	wg.Wait()
	quoteCount := len(quotes)
	testDuration := time.Now().Sub(testStart).Seconds()
	fmt.Println("Quote count: ", quoteCount)
	fmt.Println("OK: ", okCount.Load(), " Err: ", errCount.Load())
	fmt.Println("Avg seconds per quote: ", testDuration/float64(quoteCount))
	fmt.Println("Test completed in ", time.Now().Sub(testStart).Seconds(), " seconds")
	fmt.Println("First quote:\n", quotes[0])
	fmt.Println("Last quote:\n", quotes[len(quotes)-1])
}
