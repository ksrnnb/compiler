package main

import (
	"fmt"
	"os"

	"github.com/ksrnnb/compiler/tokenizer"
)

func main() {
	// 暫定的に、1つのファイル名をコマンドの引数で指定する方式をとる
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "%s: 1 argument should be given", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: file cannot open", os.Args[0])
		os.Exit(1)
	}

	tokenizer.NewTokenizer(file)
}
