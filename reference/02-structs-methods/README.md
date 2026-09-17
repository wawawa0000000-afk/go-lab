# 02 - Structs & Methods

## 目標
struct、メソッド、値/ポインタレシーバの使い分け、埋め込み（継承の代わり）を理解する。

## 実行
```bash
go run ./reference/02-structs-methods
```

## 要点

### struct はクラスではない
- データの集まり。メソッドは `func (レシーバ) 名前()` で外に書く
- **継承はない**。代わりに「埋め込み（composition）」を使う

### 値レシーバ vs ポインタレシーバ
| | `func (r T)` | `func (r *T)` |
|---|---|---|
| 受け取るもの | コピー | 元への参照 |
| 元を書き換え | できない | できる |
| 使う場面 | 小さい構造体・読み取り | 書き換える・大きい構造体 |
| 指針 | **どちらか一方に統一する**（混在させない） | |

呼び出し側は `r.Method()` と書けば Go が自動で `&r` / `*p` を補う。

### 埋め込み
```go
type Dog struct {
    Animal        // 匿名フィールド
    Breed string
}
```
`Dog` は `Animal` のフィールド・メソッドをそのまま持つ。同名メソッドを定義すれば上書き（ただし多態は interface で行う → 03）。

## 確認クイズ
1. 値レシーバのメソッド内で `r.X = 5` しても呼び出し元が変わらないのはなぜ? → **コピーを受け取っているから**
2. `a := T{1,2}; b := a` の後 `b` を変えると `a` は? → **変わらない**（struct は代入でコピー）
3. `Dog` に埋め込んだ `Animal` のメソッドを明示的に呼ぶには? → `d.Animal.Speak()`

## 練習課題
1. `Circle{Radius float64}` に `Area()` と `Perimeter()` を実装
2. `BankAccount{balance int}` に `Deposit` / `Withdraw`（残高不足はエラー）をポインタレシーバで実装
3. `Stack` 型（`[]int` を内包）に `Push` / `Pop`（空なら error）を実装
4. `Logger` struct を作り、`Server` に埋め込んで `server.Log("...")` と呼べるようにする
