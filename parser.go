package lisp

import (
	"fmt"
	"strconv"
)

func Parser(tokens []*Token) (Node, error) {
	current := 0

	walk := func() (Node, error) { return Node{}, fmt.Errorf("ugh") }

	walk = func() (Node, error) {

		switch {
		case tokens[current].Type == TokenTypeInteger:
			value, err := strconv.Atoi(tokens[current].Value)
			if err != nil {
				return Node{}, err
			}

			current++

			return Node{
				Type:     NodeTypeNumber,
				IntValue: int64(value),
			}, nil
		case tokens[current].Type == TokenTypeParen && tokens[current].Value == "(":
			// Skip opening paren
			current++

			// Get function name
			name := tokens[current].Value
			current++

			// Get arguments
			arguments := []Node{}

			for tokens[current].Type != TokenTypeParen || tokens[current].Value != ")" {
				node, err := walk()
				if err != nil {
					return Node{}, err
				}
				arguments = append(arguments, node)
			}

			// Skip closing paren
			current++

			return Node{
				Type:      NodeTypeCall,
				Name:      name,
				Arguments: arguments,
			}, nil
		}

		return Node{}, fmt.Errorf("oh no")
	}

	ast := Node{
		Type: NodeTypeProgram,
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
	Type      NodeType
	Body      []Node
	Name      string
	Arguments []Node
	IntValue  int64
}

type NodeType int

const (
	NodeTypeProgram NodeType = iota
	NodeTypeNumber
	NodeTypeCall
)
