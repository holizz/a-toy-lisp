package main

import (
	"fmt"
	"log"
)

func exec(ast Node, functions map[string]func(...LispValue) LispValue) (LispValue, error) {
	evaluate := func(node Node) (LispValue, error) {
		return LispValue{}, fmt.Errorf("aaaa")
	}
	evaluate = func(node Node) (LispValue, error) {
		switch node.Type {
		case NodeTypeNumber:
			return LispValue{
				Type:     LispValueTypeNumber,
				IntValue: node.IntValue,
			}, nil
		case NodeTypeCall:
			values := []LispValue{}
			for _, argument := range node.Arguments {
				value, err := evaluate(argument)
				if err != nil {
					return LispValue{}, err
				}
				values = append(values, value)
			}

			value := functions[node.Name](values[0], values[1])
			return value, nil
		default:
			return LispValue{}, fmt.Errorf("halp")
		}
	}

	lastValue := LispValue{
		Type: LispValueTypeNil,
	}

	for _, node := range ast.Body {
		var err error
		lastValue, err = evaluate(node)
		if err != nil {
			return LispValue{}, err
		}
	}

	return lastValue, nil
}

type LispValue struct {
	Type     LispValueType
	IntValue int64
}

type LispValueType int

const (
	LispValueTypeNumber LispValueType = iota
	LispValueTypeNil
)

var DefaultFunctions = map[string]func(...LispValue) LispValue{
	"add": func(values ...LispValue) LispValue {
		output := LispValue{
			Type:     LispValueTypeNumber,
			IntValue: 0,
		}
		for _, value := range values {
			if value.Type != LispValueTypeNumber {
				log.Fatal("oh noes")
			}

			output.IntValue += value.IntValue
		}
		return output
	},
}
