# reference/ — Go 学習リファレンス教材

Claude が用意した「読んで動かす」教材。`cmd/praNN/`（あなた自身の練習）とは分離してある。
**ここのファイルはお手本。あなたの練習は今まで通り `cmd/` と `internal/` に書く。**

## 位置づけ

| 場所 | 役割 |
|---|---|
| `cmd/praNN/` | あなたが自分で書く練習（触るのはここ） |
| `internal/<pkg>/` | あなたが書く再利用コード + テスト |
| `docs/` | あなたの学習メモ |
| `reference/` ← これ | 完成済みのお手本コード + 解説。読む・実行する・真似る用 |

## 実行方法

各フォルダは独立した `main` パッケージ（`07` を除く）。リポジトリのルートから:

```bash
cd ~/projects/go-lab
go run ./reference/01-basics
go run ./reference/04-goroutines
go run ./reference/05-http-api      # 別ターミナルで curl localhost:8080/todos
go test ./reference/07-testing
```

## カリキュラム

| # | フォルダ | テーマ | あなたの進捗との対応 |
|---|---|---|---|
| 1 | `01-basics/` | 変数・型・制御構文・関数・スライス・マップ・エラー・defer | `cmd/pra01`〜`pra03` で bool/数値/string をやった続き |
| 2 | `02-structs-methods/` | struct・値/ポインタレシーバ・埋め込み | 次のステップ |
| 3 | `03-interfaces/` | 多態・型アサーション・`Stringer`/`error`/`io.Writer` | `docs/go-design-philosophy.md` ④ の実コード版 |
| 4 | `04-goroutines/` | goroutine・channel・select・Mutex・context | Go の目玉。時間をかける |
| 5 | `05-http-api/` | `net/http` だけで REST API | 実用アプリ |
| 6 | `06-cli-tool/` | `flag`・標準入出力・ファイル操作で CLI | 実用アプリ |
| 7 | `07-testing/` | テーブル駆動テスト・ベンチ・Example | `internal/mathutil` のテストと同じスタイル |

## 使い方の流れ

1. `reference/NN-xxx/README.md` を読む
2. `reference/NN-xxx/main.go` を読んで `go run` する
3. README 末尾の「練習課題」を **`cmd/` または `internal/` に自分で書く**
   ```bash
   # 例: 01 の課題を cmd/pra04 でやる（pra04 は空枠）
   $EDITOR ~/projects/go-lab/cmd/pra04/main.go
   go run ./cmd/pra04
   ```
4. 詰まったら「`reference/03` を見ながら `cmd/pra04` を書いた。レビューして」と依頼

## 元の場所

以前 `~/go-lab`（ホーム直下）にあった内容をここへ移動・統合したもの。ホーム側は削除済み。
Go のモジュールキャッシュ `~/go` は Go ツールチェインが管理する標準ディレクトリなので残してある。
