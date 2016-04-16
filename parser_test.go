package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlank(t *testing.T) {
	ast, err := parser([]*Token{})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: NodeTypeProgram, Body: []Node{}}, ast)
}

func TestParseNumber(t *testing.T) {
	ast, err := parser([]*Token{
		{
			Type:  TokenTypeInteger,
			Value: "14",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type:     NodeTypeNumber,
			IntValue: 14,
		},
	}}, ast)
}

func TestParseSingleCallNoArguments(t *testing.T) {
	ast, err := parser([]*Token{
		{
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeName,
			Value: "foo",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type:      NodeTypeCall,
			Name:      "foo",
			Arguments: []Node{},
		},
	}}, ast)
}

func TestParseSingleCallWithArguments(t *testing.T) {
	ast, err := parser([]*Token{
		{
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeName,
			Value: "foo",
		},
		{
			Type:  TokenTypeInteger,
			Value: "123",
		},
		{
			Type:  TokenTypeInteger,
			Value: "456",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type: NodeTypeCall,
			Name: "foo",
			Arguments: []Node{
				{
					Type:     NodeTypeNumber,
					IntValue: 123,
				},
				{
					Type:     NodeTypeNumber,
					IntValue: 456,
				},
			},
		},
	}}, ast)
}

func TestParseCallRecursive(t *testing.T) {
	ast, err := parser([]*Token{
		{
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeName,
			Value: "foo",
		},
		{
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeName,
			Value: "bar",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: NodeTypeProgram, Body: []Node{
		{
			Type: NodeTypeCall,
			Name: "foo",
			Arguments: []Node{
				{
					Type:      NodeTypeCall,
					Name:      "bar",
					Arguments: []Node{},
				},
			},
		},
	}}, ast)
}
