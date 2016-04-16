package main

import (
	"fmt"
	"strconv"
)

func parse(tokens []*Token) (Node, error) {
	current := 0

	walk := func() (Node, error) { return Node{}, fmt.Errorf("ugh") }

	walk = func() (Node, error) {

		switch {
		case tokens[current].Type == "integer":
			value, err := strconv.Atoi(tokens[current].Value)
			if err != nil {
				return Node{}, err
			}

			current++

			return Node{
				Type:     "number",
				IntValue: int64(value),
			}, nil
		case tokens[current].Type == "paren" && tokens[current].Value == "(":
			// Skip opening paren
			current++

			// Get function name
			name := tokens[current].Value
			current++

			// Get arguments
			arguments := []Node{}

			for tokens[current].Type != "paren" || tokens[current].Value != ")" {
				node, err := walk()
				if err != nil {
					return Node{}, err
				}
				arguments = append(arguments, node)
			}

			// Skip closing paren
			current++

			return Node{
				Type:      "call",
				Name:      name,
				Arguments: arguments,
			}, nil
		}

		return Node{}, fmt.Errorf("oh no")
	}

	ast := Node{
		Type: "program",
		Body: []Node{},
	}

	for current < len(tokens) {
		node, err := walk()
		if err != nil {
			return Node{}, err
		}

		ast.Body = append(ast.Body, node)
	}

	return ast, nil
}

type Node struct {
	Type      string
	Body      []Node
	Name      string
	Arguments []Node
	IntValue  int64
}
