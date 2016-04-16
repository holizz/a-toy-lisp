package main

import (
	"fmt"
	"io"
	"log"

	"gopkg.in/readline.v1"
)

func main() {
	rl, err := readline.New("Toy Lisp REPL> ")
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err == io.EOF || err == readline.ErrInterrupt {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		tokens, err := tokenizer(line)
		if err != nil {
			log.Fatal(err)
		}

		ast, err := parser(tokens)
		if err != nil {
			log.Fatal(err)
		}

		output, err := exec(ast, DefaultFunctions)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v\n", output)
	}
}
