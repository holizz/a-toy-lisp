package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecEmptyBody(t *testing.T) {
	ast := Node{Type: NodeTypeProgram, Body: []Node{}}

	functions := map[string]func(...LispValue) LispValue{}

	output, err := exec(ast, functions)
	assert.Nil(t, err)
	assert.Equal(t, LispValueTypeNil, output.Type)
}

func TestExecNumber(t *testing.T) {
	ast := Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type:     NodeTypeNumber,
			IntValue: 1,
		},
	}}

	functions := map[string]func(...LispValue) LispValue{}

	output, err := exec(ast, functions)
	assert.Nil(t, err)
	assert.Equal(t, LispValueTypeNumber, output.Type)
	assert.Equal(t, int64(1), output.IntValue)
}

func TestExecSingleCall(t *testing.T) {
	ast := Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type: NodeTypeCall,
			Name: "add",
			Arguments: []Node{
				{
					Type:     NodeTypeNumber,
					IntValue: 1,
				},
				{
					Type:     NodeTypeNumber,
					IntValue: 2,
				},
			},
		},
	}}

	functions := map[string]func(...LispValue) LispValue{
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

	output, err := exec(ast, functions)
	assert.Nil(t, err)
	assert.Equal(t, LispValueTypeNumber, output.Type)
	assert.Equal(t, int64(3), output.IntValue)
}
