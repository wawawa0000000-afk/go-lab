.PHONY: run build test cover bench lint fmt vet tidy clean

run:              ## cmd/cleared/hello を実行
	go run ./cmd/cleared/hello

build:            ## 全パッケージをビルド、バイナリは bin/ へ
	go build -o bin/ ./...

test:             ## 全テスト（-race 付き）
	go test -race ./...

cover:            ## カバレッジを計測して HTML で開く
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | tail -1
	go tool cover -html=coverage.out

bench:            ## ベンチマーク
	go test -bench=. -benchmem ./...

vet:              ## 標準の静的チェック
	go vet ./...

fmt:              ## 整形（差分があれば失敗）
	gofmt -l -w .

tidy:             ## go.mod / go.sum を整理
	go mod tidy

lint: fmt vet     ## fmt + vet

clean:
	rm -rf bin coverage.out
