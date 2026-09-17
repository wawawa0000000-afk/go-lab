# 06 - CLI Tool

## 目標
`flag` でオプション解析、標準入力/ファイルの両対応、`os.Exit` とエラー出力、`go build` で単一バイナリ配布。

## 実行
```bash
# 標準入力から
echo "hello world hello go go go" | go run ./reference/06-cli-tool -top 3

# ファイルから
go run ./reference/06-cli-tool -mode lines ./reference/06-cli-tool/main.go
go run ./reference/06-cli-tool -mode chars ./go.mod

# ビルドして単体バイナリに
go build -o wc ./reference/06-cli-tool
./wc -top 5 ./reference/05-http-api/README.md
```

## 要点

### flag パッケージ
```go
mode := flag.String("mode", "words", "説明")   // *string が返る
top  := flag.Int("top", 0, "説明")
flag.Parse()                                    // ここで解析
rest := flag.Args()                             // オプション以外の引数
```
`-h` / `--help` は自動生成される。

### 標準入出力の作法
| 用途 | 書き先 |
|---|---|
| 通常の出力 | `os.Stdout`（`fmt.Println`） |
| エラー・ログ | `os.Stderr`（`fmt.Fprintln(os.Stderr, ...)`） |
| 終了コード | `os.Exit(1)`（0 以外 = 異常）。ただし defer が走らない点に注意 |

### パイプ対応
引数が無ければ `os.Stdin` を読む → `cmd1 | mytool` で使える。Unix ツールの基本。

### クロスコンパイル（Go の強み）
```bash
GOOS=windows GOARCH=amd64 go build -o wc.exe ./reference/06-cli-tool
GOOS=darwin  GOARCH=arm64 go build -o wc-mac ./reference/06-cli-tool
```
ランタイム不要の単一バイナリができる。

## 確認クイズ
1. `os.Exit(1)` の直前に登録した `defer` は実行される? → **されない**（`os.Exit` は即終了）
2. `len(text)` と `len([]rune(text))` が日本語で違うのはなぜ? → `len` は**バイト数**、rune 変換で**文字数**
3. エラーメッセージを `os.Stdout` に出すと何が困る? → パイプで次のコマンドにゴミが渡る。`os.Stderr` へ

## 練習課題
1. `-mode` に `top` を統合し、複数ファイルを引数で受けて合計する
2. `-ignore-case=false` フラグを足す
3. `wc` コマンド風に「行 単語 バイト」を1行で出す（引数なしのデフォルト動作）
4. 大きいファイル用に `io.ReadAll` をやめ、`bufio.Scanner` でストリーム処理にする
