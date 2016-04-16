package main

import (
	"fmt"
	"log"

	"github.com/gopherjs/gopherjs/js"
	"github.com/holizz/a-toy-lisp"
)

func main() {
	js.Global.Get("document").Call("addEventListener", "DOMContentLoaded", func() {
		inputField := js.Global.Get("document").Call("getElementById", "input")
		outputField := js.Global.Get("document").Call("getElementById", "output")
		submitButton := js.Global.Get("document").Call("getElementById", "submit")

		submitButton.Call("addEventListener", "click", func() {
			input := inputField.Get("value").String()

			tokens, err := lisp.Tokenizer(input)
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

			outputField.Set("value", fmt.Sprintf("%#v", output))
		})
	})
}
