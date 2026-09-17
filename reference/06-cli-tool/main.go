// Package main - 06: 単語カウント CLI ツール。
// 実行例:
//
//	echo "hello world hello" | go run ./reference/06-cli-tool
//	go run ./reference/06-cli-tool -mode lines ./reference/06-cli-tool/main.go
//	go run ./reference/06-cli-tool -top 5 ./reference/06-cli-tool/README.md
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {
	// flag でオプションを定義する
	mode := flag.String("mode", "words", "カウント単位: words | lines | chars")
	top := flag.Int("top", 0, "頻出語を上位 N 件表示 (0 で無効)")
	flag.Parse()

	// flag.Args() は オプション以外の引数（ファイル名など）
	input, name, err := openInput(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	defer input.Close()

	data, err := io.ReadAll(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "read error:", err)
		os.Exit(1)
	}
	text := string(data)

	switch *mode {
	case "words":
		words := strings.Fields(text)
		fmt.Printf("%s: %d words\n", name, len(words))
		if *top > 0 {
			printTop(words, *top)
		}
	case "lines":
		fmt.Printf("%s: %d lines\n", name, countLines(text))
	case "chars":
		fmt.Printf("%s: %d chars (runes: %d)\n", name, len(text), len([]rune(text)))
	default:
		fmt.Fprintf(os.Stderr, "unknown mode: %s\n", *mode)
		os.Exit(2)
	}
}

// 引数が無ければ標準入力、あればファイルを開く
func openInput(args []string) (io.ReadCloser, string, error) {
	if len(args) == 0 {
		return io.NopCloser(os.Stdin), "<stdin>", nil
	}
	f, err := os.Open(args[0])
	if err != nil {
		return nil, "", err
	}
	return f, args[0], nil
}

func countLines(text string) int {
	n := 0
	sc := bufio.NewScanner(strings.NewReader(text))
	for sc.Scan() {
		n++
	}
	return n
}

func printTop(words []string, n int) {
	freq := map[string]int{}
	for _, w := range words {
		freq[strings.ToLower(strings.Trim(w, ".,!?\"'()"))]++
	}

	type kv struct {
		word  string
		count int
	}
	pairs := make([]kv, 0, len(freq))
	for w, c := range freq {
		pairs = append(pairs, kv{w, c})
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count // 頻度降順
		}
		return pairs[i].word < pairs[j].word // 同数なら辞書順
	})

	fmt.Printf("--- top %d ---\n", n)
	for i := 0; i < n && i < len(pairs); i++ {
		fmt.Printf("%3d  %s\n", pairs[i].count, pairs[i].word)
	}
}
