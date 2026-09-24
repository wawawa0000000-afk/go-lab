# ▶ START HERE — 次回はここを読む

Go 学習リポジトリ。次にやることと、ディレクトリの意味だけ書いてある。

---

## 次回の始め方（コピペ用）

```bash
cd ~/projects/go-lab
claude
```

Claude は `CLAUDE.md` を自動で読むので、いきなり「続きやろう」でOK。

---

## いまどこ？（2026-09-24 時点）

| 進捗 | 状態 |
|---|---|
| 基本型（bool / 数値 / string） | ✅ 完了（`cmd/cleared/pra01-type-conversion`〜`pra03-string`） |
| 独自型・演算子・struct の生成 | ✅ 触った（`cmd/cleared/pra04-custom-type` `pra05-operators`） |
| bool 再演習・インクリメント/デクリメント | ✅ 完了（`cmd/cleared/pra06-boolean-review` `pra07-increment-decrement`） |
| 複数変数の同時宣言・複数値の戻り値 | ✅ 完了（`cmd/cleared/pra08-assignment`） |
| 変数宣言3形式（`var` / カンマ列挙 / `var (...)`）・`:=` | ✅ 完了（`cmd/cleared/pra09-variable-declaration`） |
| **struct + メソッド** | 🔜 **次はこれ**。お手本 `reference/02-structs-methods/` |
| interface / goroutine / http / testing | ⬜ その先（`reference/03`〜`07`） |

### 次の具体的アクション
1. `reference/02-structs-methods/README.md` と `main.go` を読む
2. `go run ./reference/02-structs-methods` で動かす
3. README 末尾の「練習課題」を **`cmd/pra10/main.go`**（作成済みの空の枠）に自分で書く
4. `go run ./cmd/pra10` で確認
5. Claude に「`pra10` を書いた。レビューして」と言う
6. 書き終わって確認できたら `git mv cmd/pra10 cmd/cleared/pra10-struct-method` のようにリネームして `cleared/` へ移動する

フォルダ名は `praNN-テーマ名` の形式（例: `pra01-type-conversion`）。
**クリア（完了）した練習は `cmd/cleared/` の中に集約する**。`cmd/` 直下は未着手の空枠だけにして見通しをよくする。

---

## ディレクトリの意味（これだけ覚える）

```
go-lab/
│
├── START_HERE.md   ← このファイル。次回の入口
├── CLAUDE.md       ← Claude が自動で読む状況説明（進捗の詳細もここ）
├── README.md       ← リポジトリの使い方・コマンド集
│
├── cmd/            ★ あなたが書く場所（練習プログラム）
│   ├── cleared/         ✅ クリア済みをまとめる箱（hello + 完了した praNN-テーマ名）
│   └── praNN/           ⬜ 未着手の空枠。書き終わったら cleared/ へ移動
│
├── internal/       ★ あなたが書く場所（再利用コード + テスト）
│   └── mathutil/       GCD・フィボナッチ（完成済み）
│
├── reference/      📖 Claude 作成のお手本（読む・動かす・真似る。ここは編集しない）
│   └── NN-xxx/         01基礎 02struct 03interface 04goroutine 05http 06cli 07testing
│
└── docs/           📝 学習メモ
    ├── go-basic-types.md        基本型まとめ
    ├── go-design-philosophy.md  設計思想（ポインタ安全性/多値/エラー/埋め込み/ビット演算子）
    └── 学習ログ.md              日ごとにやったことの記録
```

**迷ったらこの2択:**
- 自分で練習を書く → `cmd/` か `internal/`
- お手本を見たい → `reference/`

---

## よく使うコマンド

```bash
go run ./cmd/cleared/pra08-assignment     # 自分の練習を実行
go run ./reference/02-structs-methods   # お手本を実行
go test ./internal/...        # 自分のテスト
go test -race ./...           # 全テスト
gofmt -w .                    # 整形（一部未整形あり。気になれば実行）
go vet ./...                  # 静的チェック（現状クリーン）
```

---

## この会話より前の全記録

- Claude Code の生ログ: `~/.claude/projects/-home-wataru-fukukura/*.jsonl`
- 読みやすい要約（Linux/Docker/Go 横断）: `~/ai-coding-notes/conversation.md`
- 履歴ごと復元したい: `cd ~ && claude --resume`（この会話はホームで開始したため）
