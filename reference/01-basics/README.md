# 01 - Basics: Go の基本文法

## 目標
変数・型・制御構文・関数・スライス・マップ・エラー処理・defer を一通り書ける。

## 実行
```bash
cd ~/projects/go-lab
go run ./reference/01-basics
```

## Go の特徴（他言語との違い）
- **型変換は常に明示的**。`int` と `float64` を勝手に足せない
- **ループは `for` だけ**。`while` は `for cond {}`、無限ループは `for {}`
- **`:=` は関数の中だけ**。パッケージレベルは `var`
- **未使用の変数・import はコンパイルエラー**
- **例外がない**。エラーは戻り値で返す（`value, err := ...`）
- **整形スタイルは1つ**。`gofmt` が唯一の正解。タブインデント

## 覚えるイディオム
| イディオム | 例 |
|---|---|
| カンマ ok | `v, ok := m[key]` / `v, ok := x.(T)` |
| エラーチェック | `if err != nil { return err }` |
| エラーラップ | `fmt.Errorf("...: %w", err)` |
| 取り出し | `errors.As(err, &target)` / `errors.Is(err, ErrX)` |
| 無視 | `_ = someValue` |
| defer で後片付け | `f, _ := os.Open(...); defer f.Close()` |

## 確認クイズ
1. `var x int` の `x` の値は? → **0**（ゼロ値。string は `""`, bool は `false`, ポインタは `nil`）
2. `s := []int{1,2,3}; sub := s[0:2]; sub[0] = 99` の後、`s[0]` は? → **99**（スライスは配列を共有）
3. なぜ Go には `try/catch` が無い? → エラーを**値として**明示的に扱わせる設計思想
4. `defer` を3つ登録した順に a,b,c。実行順は? → **c, b, a**（LIFO）

## 練習課題
1. `fizzbuzz(n int) string` を書き、1〜20 で実行する
2. `map[string]int` で文章中の単語出現回数を数える関数を書く
3. `divide` を改造し、`b == 0` のとき独自エラー型 `DivByZeroError` を返すようにする
4. スライスから重複を除く `unique([]int) []int` を書く
