# 03 - Interfaces

## 目標
インターフェースによる多態、型アサーション / type switch、標準インターフェース（`Stringer`, `error`, `io.Writer`）の実装。

## 実行
```bash
go run ./reference/03-interfaces
```

## 要点

### 構造的型付け（暗黙の実装）
```go
type Shape interface { Area() float64 }
type Circle struct{ R float64 }
func (c Circle) Area() float64 { ... }   // これだけで Circle は Shape
```
`implements Shape` と書かない。メソッドを満たせば自動的にその型として使える。

### 小さいインターフェースが良い
標準ライブラリの `io.Reader` / `io.Writer` は1メソッド。
**「関数が必要とするものだけ」を、使う側のパッケージで定義する**のが Go 流。

### 型を取り出す
| 方法 | 用途 |
|---|---|
| `v, ok := x.(Circle)` | 1つの型か確認 |
| `switch v := x.(type)` | 複数の型で分岐 |

### よく実装する標準インターフェース
| インターフェース | メソッド | 効果 |
|---|---|---|
| `fmt.Stringer` | `String() string` | `Println` での表示をカスタム |
| `error` | `Error() string` | エラー型になる |
| `io.Writer` | `Write([]byte) (int, error)` | `Fprintf` などの出力先になれる |
| `sort.Interface` | `Len/Less/Swap` | `sort.Sort` で並べ替え可能に |

## 確認クイズ
1. Go で「この型はこのインターフェースを実装します」と宣言する構文は? → **無い**（メソッドが揃えば自動）
2. `var s Shape` の `s` の値は? → **nil**
3. `any` の正体は? → `interface{}` の別名（メソッド0個 = 何でも入る）
4. type switch と型アサーションの違いは? → switch は複数型の分岐、アサーションは単一型の確認

## 練習課題
1. `Stringer` を `Money{Cents int}` に実装し `$12.34` と表示させる
2. `Notifier` interface（`Notify(msg string) error`）を定義し、`EmailNotifier` / `SlackNotifier` を作って `[]Notifier` にループで送る
3. `sort.Interface` を `ByLength []string` に実装し、文字列を長さ順にソート
4. 独自 `error` 型 `ValidationError{Field, Message string}` を作り、`errors.As` で取り出す
