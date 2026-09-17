# CLAUDE.md

Go の練習用リポジトリ（`github.com/wawawa0000000-afk/go-lab`, go 1.25）。オーナーは Go を学習中。
人間向けの入口は `START_HERE.md`、使い方は `README.md`。

---

## ▶ 次にやること（セッション開始時はまずここ）

最終更新: 2026-09-07

- **到達点**: 基本型（bool / 数値 / string）＋ 演算子 ＋ 独自型・struct の生成まで。
- **次のテーマ**: **struct + メソッド**（値/ポインタレシーバ、埋め込み）。お手本 = `reference/02-structs-methods/`。
- **次の練習枠**: `cmd/pra06/`（未作成。`mkdir` して `main.go` を新規作成する）。
- オーナーが `reference/02` を読んで `cmd/pra06` に自分で書く → Claude はレビュー役。

---

## ディレクトリ

| 場所 | 役割 | 誰が書く |
|---|---|---|
| `cmd/<name>/main.go` | 実行プログラム。1 テーマ = 1 ディレクトリ | **オーナー** |
| `internal/<pkg>/` | 再利用コード + テーブル駆動テスト | **オーナー** |
| `reference/NN-xxx/` | Claude 作成の完成済みお手本（01〜07）。読む・動かす・真似る | Claude（**編集しない**。練習の答えにしない） |
| `docs/` | 学習メモ | 両方 |

- `docs/go-basic-types.md` … 基本型（bool / 数値 / string）まとめ
- `docs/go-design-philosophy.md` … 設計思想（ポインタ安全性 / 多値 / エラー / 埋め込み / ビット演算子）
- `docs/学習ログ.md` … 日ごとにやったことの記録
- git はまだ初期化していない

## 作業の進め方

- 練習コードなので **先回りして完成させない**。ヒント・レビュー・詰まった箇所の説明が中心。
- 例外: オーナーが「コメントを付けて」「レビューして」と明示したら、そのファイルへの追記・修正は可。
- `cmd/` の既存ファイルと `docs/` の既存ファイルは、指示がない限り書き換えない。
- 新しい練習は「ディレクトリと `main.go` を作るだけ」で OK（`go` がディレクトリを走査する）。
- 変更後は `go vet ./...` と `go test -race ./...` で確認。整形は `gofmt -w .`。

## よく使うコマンド

```bash
go run ./cmd/pra06                    # オーナーの練習を実行
go run ./reference/02-structs-methods # お手本を実行
go test ./internal/...               # オーナーのテスト
go test -race ./...                  # 全テスト
gofmt -w .                           # 整形
go vet ./...                         # 静的チェック（現状クリーン）
```

`make` は未インストールのことがある。その場合は上の `go` コマンドを直接使う（Makefile と同等）。

## 進捗メモ（詳細）

**オーナーの学習**
- `internal/mathutil` … `GCD`, `FibSequence` + テスト・ベンチ。完成。
- `cmd/hello` … サンプル。
- `cmd/pra01` … 型変換（int→int64）、関数の戻り値。完了。
- `cmd/pra02` … bool、論理演算子 `&&` `||`。完了。
- `cmd/pra03` … string 連結、`strings.Replace`、`len` vs `utf8.RuneCountInString`。完了。
- `cmd/pra04` … 独自型（`type myInteger int`）、struct の生成4種、型変換（`changer`）。完了。
  - メモ: `string(int)` は `go vet` が警告（コードポイント変換）。`string(rune(i))` か `strconv.Itoa` で回避。
- `cmd/pra05` … 全演算子（算術 `+ - * / %` / ビット `& | ^ &^` / シフト `<< >>`）。演算子ごとの解説コメント付き。完了。
- `docs/` … `go-basic-types.md`, `go-design-philosophy.md`, `学習ログ.md` 作成済み。
- 未整形: `cmd/pra01`〜`pra04`（`gofmt -w .` で直せる。動作には影響なし）。

**お手本教材（reference/、完成済み）**
`01` 基礎 / `02` struct・メソッド / `03` interface / `04` goroutine・channel・context /
`05` net/http REST API / `06` flag CLI / `07` testing（`package stats` + テスト）。
元は `~/go-lab`（ホーム直下）にあったものを 2026-09-07 にここへ統合。ホーム側は削除済み。

## 学習の推奨サイクル

1. `reference/NN-xxx/` の README と `main.go` を読んで `go run`
2. README 末尾の「練習課題」を `cmd/praNN`（または `internal/`）に**自分で**書く
3. `go run ./cmd/praNN` で確認
4. 「`cmd/praNN` を書いた。レビューして」と Claude に依頼
5. この「▶ 次にやること」と「進捗メモ」を更新する
