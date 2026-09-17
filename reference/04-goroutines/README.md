# 04 - Goroutines & Concurrency

Go を学ぶ最大の理由。ここは時間をかけて。

## 目標
goroutine、channel、`select`、`sync.WaitGroup` / `sync.Mutex`、`context` を使える。

## 実行
```bash
go run ./reference/04-goroutines

# race detector 付きで実行（並行バグの検出）
go run -race ./reference/04-goroutines
```

## 要点

### goroutine
- `go f()` で起動する軽量スレッド。数万個でも動く
- `main` が終わると走行中の goroutine も強制終了する → 待つ仕組みが必要

### 待ち方
| 方法 | 用途 |
|---|---|
| `sync.WaitGroup` | N 個の goroutine の完了を待つ |
| channel 受信 | 結果を1つ受け取る |
| `errgroup`（準標準） | 複数 goroutine ＋エラー集約 |

### channel
```go
ch := make(chan int)      // 同期（バッファなし）: 送受信がそろうまでブロック
ch := make(chan int, 10)  // バッファあり: 10個まで詰められる
ch <- v                   // 送信
v := <-ch                 // 受信
v, ok := <-ch             // ok=false なら close 済み
close(ch)                 // 送信側が閉じる（受信側は閉じない）
for v := range ch { }     // close まで受信し続ける
```

### 格言
> **Don't communicate by sharing memory; share memory by communicating.**
> （メモリを共有して通信するな。通信してメモリを共有せよ ＝ できるだけ channel を使い、Mutex は最小限に）

### context
- タイムアウト・キャンセルを goroutine ツリー全体に伝播させる
- HTTP ハンドラや外部呼び出しでは第1引数に `ctx context.Context` を渡すのが慣習
- `ctx.Done()` を `select` で監視して中断する

## よくあるバグ
1. **ループ変数のキャプチャ**（Go 1.22 で挙動が改善されたが、明示的に引数で渡す癖を）
2. **WaitGroup.Add をgoroutine の中で呼ぶ** → 競合。外で呼ぶ
3. **close 済み channel への送信** → panic
4. **Mutex を値コピー** → ロックが効かない。`*T` レシーバで持つ

## 確認クイズ
1. バッファなし channel への送信はいつ返る? → **誰かが受信した瞬間**
2. `range ch` が終わるのは? → channel が **close** されたとき
3. `go run -race` は何を検出する? → 複数 goroutine が同じメモリを同期なしで読み書きしている箇所
4. `context.WithTimeout` の `cancel` を defer で呼ぶ理由は? → リソースリーク防止（必ず解放する）

## 練習課題
1. 3つの URL を並行取得し、最初に返ってきた本文を表示（`select` を使う）
2. 数値スライスを N 分割し、各区間の合計を goroutine で計算して合算
3. `context` を使い、2秒でタイムアウトする「重い計算」を書く
4. `-race` でエラーが出るコードを書き、Mutex で直す
