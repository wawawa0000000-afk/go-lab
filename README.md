# go-lab

Go の練習用リポジトリ。Go はツールチェインにビルド・テスト・整形・カバレッジが全部入っているので、
外部ツールはほぼ不要。

**次回はまず `START_HERE.md` を読む。** Claude を使うなら `cd ~/projects/go-lab && claude`。

## レイアウト

```
go-lab/
├── START_HERE.md   次回の入口（次にやること）
├── CLAUDE.md       Claude が自動で読む状況説明・進捗メモ
├── README.md       このファイル（使い方・コマンド）
│
├── cmd/            ★ 自分で書く: 実行プログラム。1 テーマ = 1 ディレクトリ
│   ├── hello/          サンプル
│   ├── pra01/          型変換・関数の戻り値
│   ├── pra02/          bool・論理演算子
│   ├── pra03/          string・Replace・len/rune
│   ├── pra04/          独自型・struct の生成・型変換
│   └── pra05/          全演算子（算術・ビット・シフト）＋解説コメント
│
├── internal/       ★ 自分で書く: 再利用コード + テスト（モジュール外から import 不可）
│   └── mathutil/       GCD・フィボナッチ（完成）
│
├── reference/      📖 Claude 作成のお手本（読む・動かす・真似る。編集しない）
│   └── 01..07/         基礎 / struct / interface / goroutine / http / cli / testing
│
└── docs/           📝 学習メモ
    ├── go-basic-types.md
    ├── go-design-philosophy.md
    └── 学習ログ.md
```

**迷ったら**: 自分で練習を書く → `cmd/` か `internal/`／お手本を見る → `reference/`

## 必要なもの（Ubuntu 26.04）

```bash
sudo apt-get update
sudo apt-get install -y golang-go   # Go 1.26
# 任意: 追加のリンター
# go install honnef.co/go/tools/cmd/staticcheck@latest
```

## 基本コマンド

```bash
go run ./cmd/hello        # 実行
go build ./...            # 全部ビルド（コンパイル通るか確認）
go test ./...             # 全テスト
go test -race ./...       # データ競合検出つき（おすすめ）
go test -v ./internal/mathutil   # 1 パッケージだけ詳細表示
go test -run TestGCD ./...        # 名前で絞り込み
```

Makefile 経由でも同じことができる:

```bash
make test      # go test -race ./...
make cover     # カバレッジ計測 + HTML 表示
make bench     # ベンチマーク
make lint      # gofmt + go vet
```

## 練習を 1 個増やす手順

**ライブラリ的なコード:**
1. `internal/foo/foo.go` に `package foo` で関数を書く
2. `internal/foo/foo_test.go` に `TestXxx(t *testing.T)` を書く（テーブル駆動が定番）
3. `go test ./internal/foo`

**動かして確かめるプログラム:**
1. `cmd/bar/main.go` に `package main` + `func main()`
2. `go run ./cmd/bar`

新規ファイル・ディレクトリを作るだけで OK。設定ファイルの編集は不要
（`go` コマンドがディレクトリを走査する）。

## テストの書き方メモ

- テストファイルは `*_test.go`、同じパッケージに置く
- `t.Run(name, func(t *testing.T){...})` でサブテスト（テーブル駆動）
- `t.Errorf` は続行、`t.Fatalf` は即中断
- ベンチは `func BenchmarkXxx(b *testing.B)`、`for i := 0; i < b.N; i++`
- `go test -race` を習慣に。並行処理の練習で効く

## デバッグ

```bash
sudo apt-get install -y delve   # または: go install github.com/go-delve/delve/cmd/dlv@latest
dlv test ./internal/mathutil
dlv debug ./cmd/hello
```
