package main

import (
	"fmt"
	"io"
	"log"

	"gopkg.in/readline.v1"

	"github.com/holizz/a-toy-lisp"
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

		tokens, err := lisp.Tokenizer(line)
		if err != nil {
			log.Fatal(err)
		}

		ast, err := lisp.Parser(tokens)
		if err != nil {
			log.Fatal(err)
		}

		output, err := lisp.Exec(ast, lisp.DefaultFunctions)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v\n", output)
	}
}
