# CLAUDE.md

Go の練習用リポジトリ（`github.com/wawawa0000000-afk/go-lab`, go 1.25）。オーナーは Go を学習中。
人間向けの入口は `START_HERE.md`、使い方は `README.md`。

---

## ▶ 次にやること（セッション開始時はまずここ）

最終更新: 2026-09-24

- **到達点**: 基本型（bool / 数値 / string）＋ 演算子 ＋ 独自型・struct の生成 ＋ インクリメント/デクリメント ＋ 複数変数の同時宣言・代入 ＋ 変数宣言3形式（`var` / 初期値付き `var` / `:=`）まで。
- **次のテーマ**: **struct + メソッド**（値/ポインタレシーバ、埋め込み）。お手本 = `reference/02-structs-methods/`。
- **次の練習枠**: `cmd/pra10`〜`cmd/pra20`（空の雛形）。どれか1つに `reference/02` の練習課題を書く。
- オーナーが `reference/02` を読んで `cmd/pra10`（など）に自分で書く → Claude はレビュー役。
  書き終わってビルドが通ったら `praNN-struct-method` のようにリネームして `cmd/cleared/` へ移動する。

### フォルダ運用ルール

- 命名: `cmd/praNN-テーマ名`（kebab-case の英単語）。中身を書き始めたら `git mv` でテーマ名を付ける。
  テーマ未定の空の枠（`pra10`〜`pra20` など）は番号だけのままで OK。
- **クリア（完了）済みは `cmd/cleared/` の中にまとめる**。`cmd/` 直下は未着手の空枠だけにして見通しをよくする。
  動作確認まで終わったら `git mv cmd/praNN-テーマ名 cmd/cleared/praNN-テーマ名`。

---

## ディレクトリ

| 場所 | 役割 | 誰が書く |
|---|---|---|
| `cmd/cleared/<name>/main.go` | クリア済みの実行プログラム | **オーナー** |
| `cmd/praNN/main.go` | 未着手の練習枠（空の雛形） | **オーナー** |
| `internal/<pkg>/` | 再利用コード + テーブル駆動テスト | **オーナー** |
| `reference/NN-xxx/` | Claude 作成の完成済みお手本（01〜07）。読む・動かす・真似る | Claude（**編集しない**。練習の答えにしない） |
| `docs/` | 学習メモ | 両方 |

- `docs/go-basic-types.md` … 基本型（bool / 数値 / string）まとめ
- `docs/go-design-philosophy.md` … 設計思想（ポインタ安全性 / 多値 / エラー / 埋め込み / ビット演算子）
- `docs/学習ログ.md` … 日ごとにやったことの記録
- `docs/進捗チェックリスト.md` … `pra01`〜`pra20` の完了チェックボックス
- git は初期化済み（2026-09-23〜）

## 作業の進め方

- 練習コードなので **先回りして完成させない**。ヒント・レビュー・詰まった箇所の説明が中心。
- 例外: オーナーが「コメントを付けて」「レビューして」「リネームして」「移動して」などと明示したら、そのファイル・フォルダへの追記・修正は可。
- **日本語コメントの挿入は基本的に Claude が代行する**（VSCode + Vim拡張機能の組み合わせで IME 変換が不安定なため）。
  「コメント入れて」（内容はコードから判断）、または「ここに〇〇って書いて」（文面を指定）と言われたら挿入する。
  コード本体（練習の答え）は書かない。詳細: [[vscode-vim-ime-bug]]（メモリ）。
- `cmd/` の既存ファイルと `docs/` の既存ファイルは、指示がない限り書き換えない。
- 新しい練習は「ディレクトリと `main.go` を作るだけ」で OK（`go` がディレクトリを走査する）。
- フォルダをリネーム・移動するときは `git mv` を使う（履歴を保つため）。
- 変更後は `go vet ./...` と `go test -race ./...` で確認。整形は `gofmt -w .`。
- **練習が完成してビルド・動作確認（`go build ./...` / `go run`）が通ったら、そのフォルダを
  `praNN` → `praNN-テーマ名` にリネームしつつ `cmd/cleared/` の中へ移動する**（`git mv` 一発で両方できる）。
  すでに `cleared/` に入っている・すでにテーマ名が付いている場合はやり直さなくてよい。

## よく使うコマンド

```bash
go run ./cmd/cleared/pra08-assignment     # オーナーの練習を実行
go run ./reference/02-structs-methods     # お手本を実行
go test ./internal/...                   # オーナーのテスト
go test -race ./...                      # 全テスト
gofmt -w .                               # 整形
go vet ./...                             # 静的チェック（現状クリーン）
```

`make` は未インストールのことがある。その場合は上の `go` コマンドを直接使う（Makefile と同等）。

## 進捗メモ（詳細）

**オーナーの学習**
- `internal/mathutil` … `GCD`, `FibSequence` + テスト・ベンチ。完成。
- `cmd/cleared/hello` … サンプル。
- `cmd/cleared/pra01-type-conversion` … 型変換（int→int64）、関数の戻り値。完了。
- `cmd/cleared/pra02-boolean` … bool、論理演算子 `&&` `||`。完了。
- `cmd/cleared/pra03-string` … string 連結、`strings.Replace`、`len` vs `utf8.RuneCountInString`。完了。
- `cmd/cleared/pra04-custom-type` … 独自型（`type myInteger int`）、struct の生成4種、型変換（`changer`）。完了。
  - メモ: `string(int)` は `go vet` が警告（コードポイント変換）。`string(rune(i))` か `strconv.Itoa` で回避。
- `cmd/cleared/pra05-operators` … 全演算子（算術 `+ - * / %` / ビット `& | ^ &^` / シフト `<< >>`）。演算子ごとの解説コメント付き。完了。
- `cmd/cleared/pra06-boolean-review` … bool、論理演算子 `&&` `||`（pra02 の再演習）。完了。
- `cmd/cleared/pra07-increment-decrement` … `inc++` / `dec--`（インクリメント・デクリメント）。完了。
- `cmd/cleared/pra08-assignment` … 複数変数の同時宣言（`:=`）、複数値を返す関数（`func fn() (int, int)`）。完了。
- `cmd/cleared/pra09-variable-declaration` … 変数宣言の3形式（`var x int` / カンマ列挙 `var a, b, c int` / 丸括弧グルーピング `var (...)`）、初期値付き `var`、`:=`（関数 `syouryaku` `tyokusetu` に分けて比較）。完了。
- `cmd/pra10`〜`pra20` … 空の雛形。`cmd/` 直下に待機中。次のテーマは struct + メソッド。テーマが決まって書き終わったら `cleared/` へ移動しリネーム。
- VS Code 設定: Go ファイルは保存時に自動整形（`[go] editor.formatOnSave`）。ただし `files.autoSave: afterDelay` の自動保存では整形されない（`Ctrl+S` で整形）。
- `docs/` … `go-basic-types.md`, `go-design-philosophy.md`, `学習ログ.md`, `進捗チェックリスト.md` 作成済み。
- 未整形: `cmd/cleared/pra01-type-conversion`〜`pra04-custom-type`, `cmd/cleared/pra07-increment-decrement`（`gofmt -w .` で直せる。動作には影響なし）。

**お手本教材（reference/、完成済み）**
`01` 基礎 / `02` struct・メソッド / `03` interface / `04` goroutine・channel・context /
`05` net/http REST API / `06` flag CLI / `07` testing（`package stats` + テスト）。
元は `~/go-lab`（ホーム直下）にあったものを 2026-09-07 にここへ統合。ホーム側は削除済み。

## 学習の推奨サイクル

1. `reference/NN-xxx/` の README と `main.go` を読んで `go run`
2. README 末尾の「練習課題」を空の練習枠（`cmd/praNN`）に**自分で**書く
3. テーマが固まったら `git mv cmd/praNN cmd/cleared/praNN-テーマ名` でリネーム兼移動
4. `go run ./cmd/cleared/praNN-テーマ名` で確認
5. 「`cmd/cleared/praNN-テーマ名` を書いた。レビューして」と Claude に依頼
6. この「▶ 次にやること」と「進捗メモ」「進捗チェックリスト.md」を更新する
