package ws

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"
)

// go test -v -timeout 30s -count=1 -run TestSendOnClosedChannel AIL2-ProveNode/ws
func TestSendOnClosedChannel(t *testing.T) {
	wg := sync.WaitGroup{}

	// sender := func(ch chan int) {
	// 	wg.Add(1)
	// 	defer wg.Done()
	// 	for i := 0; i < 100; i++ {
	// 		// ch <- i
	// 		select {
	// 		case ch <- i:
	// 		default:
	// 			log.Println("sender: channel closed")
	// 			return
	// 		}
	// 		time.Sleep(100 * time.Millisecond)
	// 	}
	// }
	senderDone := func(ch chan int, done chan bool) {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 100; i++ {
			// ch <- i
			select {
			case <-done:
				log.Println("sender: channel closed")
				return
			default:
				ch <- i
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	senderCtx := func(ctx context.Context, ch chan int) {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 100; i++ {
			// ch <- i
			select {
			case <-ctx.Done():
				log.Println("sender: channel closed")
				return
			default:
				ch <- i
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	receiver := func(ch chan int) {
		wg.Add(1)
		defer wg.Done()
		for {
			val, ok := <-ch
			if ok {
				log.Printf("read channel %v\n", val)
			} else {
				log.Println("receiver: channel closed")
				return
			}
		}
	}

	ch := make(chan int, 10)
	done := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())

	go receiver(ch)
	for i := 0; i < 3; i++ {
		// go sender(ch)
		go senderDone(ch, done)
		go senderCtx(ctx, ch)
	}

	time.Sleep(1 * time.Second)
	close(done)
	cancel()
	close(ch)

	wg.Wait()
}
