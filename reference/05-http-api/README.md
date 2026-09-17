# 05 - HTTP API

## 目標
標準ライブラリ `net/http` だけで REST API を書く。JSON、ルーティング、ミドルウェア、グレースフルな設計。

## 実行
```bash
go run ./reference/05-http-api
```
別ターミナルで:
```bash
curl localhost:8080/health
curl localhost:8080/todos
curl -X POST localhost:8080/todos -d '{"title":"buy milk"}'
curl localhost:8080/todos/1
curl -X DELETE localhost:8080/todos/1 -i
```

## 要点

### Go 1.22+ の新ルーティング
```go
mux.HandleFunc("GET /todos/{id}", handler)
id := r.PathValue("id")
```
以前は `gorilla/mux` などの外部ライブラリが要ったが、今は標準で足りることが多い。

### JSON
| 方向 | 方法 |
|---|---|
| struct → JSON | `json.NewEncoder(w).Encode(v)` |
| JSON → struct | `json.NewDecoder(r.Body).Decode(&v)` |
| タグ | `` `json:"id"` `` でフィールド名を指定、`,omitempty` で空を省略 |

### ミドルウェア
`func(http.Handler) http.Handler` の形。ロギング・認証・CORS などを重ねる。

### 本番で足すもの
- `http.Server` に `ReadTimeout` / `WriteTimeout`（DoS 対策・必須）
- グレースフルシャットダウン（`srv.Shutdown(ctx)` を signal で）
- 構造化ログ（`log/slog`）
- バリデーション、エラーレスポンスの統一

## 確認クイズ
1. `w.WriteHeader(200)` を呼んだ後に `w.Header().Set(...)` しても効かないのはなぜ? → ヘッダは**ステータス送信時に確定**するから。ヘッダ → WriteHeader → Body の順
2. `json.Decode` で未知のフィールドはどうなる? → 既定では**無視**（`DisallowUnknownFields()` で拒否できる）
3. なぜ map を返すハンドラで Mutex が要る? → 複数リクエストが**別 goroutine**で同時に走るから

## 練習課題
1. `PUT /todos/{id}` で `done` を更新できるようにする
2. `GET /todos?done=true` のクエリでフィルタする
3. API キー認証のミドルウェアを追加（`Authorization` ヘッダをチェック）
4. `os/signal` で Ctrl+C を捕まえ、`srv.Shutdown` で綺麗に終了する
5. ストアを `04` の PostgreSQL + Docker compose に差し替える
