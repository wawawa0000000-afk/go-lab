# 07 - Testing

## 目標
`testing` パッケージでユニットテスト、テーブル駆動テスト、ベンチマーク、カバレッジ。外部フレームワーク不要。

## 実行
```bash
go test ./reference/07-testing              # 実行
go test -v ./reference/07-testing           # 各テスト名を表示
go test -run TestMedian ./reference/07-testing        # 名前で絞る
go test -run TestMedian/偶数個 ./reference/07-testing   # サブテストで絞る
go test -cover ./reference/07-testing       # カバレッジ率
go test -bench=. ./reference/07-testing     # ベンチマーク
go test -race ./...               # 競合検出つき

# カバレッジを HTML で見る
go test -coverprofile=cov.out ./reference/07-testing && go tool cover -html=cov.out
```

## 要点

### ファイルの約束
- テストは `*_test.go` に書く（ビルド対象外になる）
- 関数は `func TestXxx(t *testing.T)`
- 同じパッケージ名にすると内部にアクセス可、`パッケージ名_test` にすると外部利用者の視点でテスト

### t のメソッド
| メソッド | 意味 |
|---|---|
| `t.Errorf(...)` | 失敗を記録して**続行** |
| `t.Fatalf(...)` | 失敗を記録して**そのテストを中断** |
| `t.Run(name, fn)` | サブテスト（テーブル駆動で使う） |
| `t.Helper()` | このフレームを stack trace から隠す（ヘルパー関数用） |
| `t.Cleanup(fn)` | テスト終了時の後片付け |
| `t.Parallel()` | 他の Parallel テストと並行実行 |

### テーブル駆動テスト
入力と期待値を `[]struct{...}` に並べ、ループで `t.Run` する。Go で最も一般的なスタイル。ケース追加が1行で済む。

### 比較のコツ
- 数値・文字列: `==`
- スライス: `slices.Equal`
- map: `maps.Equal`
- 複雑な構造体: `reflect.DeepEqual` か `github.com/google/go-cmp`

### ベンチマーク
```go
func BenchmarkX(b *testing.B) {
    for b.Loop() {   // Go 1.24+。旧: for i := 0; i < b.N; i++
        X()
    }
}
```

## 確認クイズ
1. `t.Error` と `t.Fatal` の違いは? → Fatal はそのテスト関数を即中断（後続のアサーションを実行しない）
2. テストファイルは本番バイナリに含まれる? → **含まれない**（`_test.go` はビルド対象外）
3. テーブル駆動で `t.Run` を使う利点は? → 失敗したケース名が分かる・個別に絞って実行できる

## 練習課題
1. `Mode`（最頻値）関数を追加し、テーブル駆動でテスト
2. `Median` に「NaN が混ざったらエラー」を追加し、そのテストを書く
3. `t.Parallel()` を付けて実行時間が変わるか確認
4. カバレッジ 100% を目指して不足ケースを埋める
5. `ExampleIsPalindrome` に `// Output:` 行を足して doctest 化する
