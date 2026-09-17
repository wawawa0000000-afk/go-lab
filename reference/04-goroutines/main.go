// Package main - 04: 並行処理。Go の目玉機能。
// 実行: go run ./reference/04-goroutines
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 1. goroutine と WaitGroup ===")
	waitGroupDemo()

	fmt.Println("\n=== 2. channel（値の受け渡し） ===")
	channelDemo()

	fmt.Println("\n=== 3. worker pool ===")
	workerPoolDemo()

	fmt.Println("\n=== 4. select（複数 channel を待つ） ===")
	selectDemo()

	fmt.Println("\n=== 5. context でキャンセル ===")
	contextDemo()

	fmt.Println("\n=== 6. Mutex で共有状態を守る ===")
	mutexDemo()
}

// goroutine は go キーワードで起動する軽量スレッド。
// WaitGroup で「全部終わるまで待つ」
func waitGroupDemo() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(time.Duration(id) * 20 * time.Millisecond)
			fmt.Printf("  worker %d 完了\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("  全員終わった")
}

// channel は goroutine 間で安全に値を渡すパイプ
func channelDemo() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i * 10 // 送信
		}
		close(ch) // 送信終了を通知
	}()

	// range は channel が close されるまで受信し続ける
	for v := range ch {
		fmt.Println("  受信:", v)
	}
}

// 決まった数の worker で大量のジョブを処理する定番パターン
func workerPoolDemo() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				results <- j * j
			}
		}(w)
	}

	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for r := range results {
		sum += r
	}
	fmt.Println("  1..6 の2乗の合計:", sum) // 91
}

func selectDemo() {
	fast := make(chan string)
	slow := make(chan string)

	go func() { time.Sleep(10 * time.Millisecond); fast <- "fast" }()
	go func() { time.Sleep(50 * time.Millisecond); slow <- "slow" }()

	for i := 0; i < 2; i++ {
		select {
		case m := <-fast:
			fmt.Println("  fast から:", m)
		case m := <-slow:
			fmt.Println("  slow から:", m)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("  タイムアウト")
		}
	}
}

// context: キャンセル・タイムアウトを goroutine ツリー全体に伝える
func contextDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Millisecond)
	defer cancel()

	result := make(chan int)
	go func() {
		time.Sleep(100 * time.Millisecond) // 重い処理のつもり
		result <- 42
	}()

	select {
	case r := <-result:
		fmt.Println("  結果:", r)
	case <-ctx.Done():
		fmt.Println("  キャンセルされた:", ctx.Err()) // context deadline exceeded
	}
}

func mutexDemo() {
	var (
		mu      sync.Mutex
		counter int
		wg      sync.WaitGroup
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++ // この行を mu で守らないと race condition
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("  counter (期待値 1000):", counter)
}
