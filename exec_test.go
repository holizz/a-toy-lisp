package lisp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecEmptyBody(t *testing.T) {
	ast := Node{Type: NodeTypeProgram, Body: []Node{}}

	output, err := Exec(ast, DefaultFunctions)
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

	output, err := Exec(ast, DefaultFunctions)
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

	output, err := Exec(ast, DefaultFunctions)
	assert.Nil(t, err)
	assert.Equal(t, LispValueTypeNumber, output.Type)
	assert.Equal(t, int64(3), output.IntValue)
}

func TestExecSingleCallWithManyValues(t *testing.T) {
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
				{
					Type:     NodeTypeNumber,
					IntValue: 3,
				},
				{
					Type:     NodeTypeNumber,
					IntValue: 4,
				},
				{
					Type:     NodeTypeNumber,
					IntValue: 5,
				},
			},
		},
	}}

	output, err := Exec(ast, DefaultFunctions)
	assert.Nil(t, err)
	assert.Equal(t, LispValueTypeNumber, output.Type)
	assert.Equal(t, int64(15), output.IntValue)
}
