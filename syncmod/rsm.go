package syncmod

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/thanhpk/randstr"
)

func Rsm_main() {
	urls := []string{}
	N := 10
	for range N {
		urls = append(urls, randstr.Hex(32))
	}
	rev := []chan *string{}
	n := 3
	wg := sync.WaitGroup{}
	wg1 := sync.WaitGroup{}
	for i := range n {
		rev = append(rev, make(chan *string, 1))
		wg.Add(1)
		go rever(rev[i], wg.Done)
	}
	ctx, cancel := context.WithCancel(context.Background())
	wg1.Add(1)
	go sender(ctx, rev, urls, wg1.Done)
	wg.Wait()
	cancel()
	wg1.Wait()
	fmt.Println("all done")
}

func sender(ctx context.Context, rev []chan *string, urls []string, done func()) {
	rev_i := 0
	is_exit := ctx.Done()
	isbreak := false
	for i := 0; i < len(urls) && !isbreak; {
		select {
		case rev[rev_i] <- &urls[i]:
			rev_i = (rev_i + 1) % len(rev)
			i++
		case <-is_exit:
			isbreak = true
		default:
			rev_i = (rev_i + 1) % len(rev)
		}
	}
	fmt.Println("close rev")
	for i := range rev {
		close(rev[i])
	}
	done()
	fmt.Println("sender done")
}

func rever(r <-chan *string, done func()) {
	isbreak := false
	for !isbreak {
		url, open := <-r
		if url != nil {
			rever_handle(*url)
		}
		if open {
			if url == nil {
				fmt.Println("!!!!!!!!!!!!!!!!!!")
			}
			continue
		} else {
			chanlen := len(r)
			for range chanlen {
				rever_handle(*(<-r))
			}
		}
		isbreak = true
	}
	done()
	fmt.Println("rever done")
}

func rever_handle(urls ...string) {
	time.Sleep(time.Second)
	for i := range urls {
		fmt.Println("handle", urls[i])
	}
}
