package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBlank(t *testing.T) {
	ast, err := parse([]*Token{})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: "program", Body: []Node{}}, ast)
}

func TestParseNumber(t *testing.T) {
	ast, err := parse([]*Token{
		{
			Type:  "integer",
			Value: "14",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: "program", Body: []Node{
		{
			Type:     "number",
			IntValue: 14,
		},
	}}, ast)
}

func TestParseSingleCallNoArguments(t *testing.T) {
	ast, err := parse([]*Token{
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "name",
			Value: "foo",
		},
		{
			Type:  "paren",
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: "program", Body: []Node{
		{
			Type:      "call",
			Name:      "foo",
			Arguments: []Node{},
		},
	}}, ast)
}

func TestParseSingleCallWithArguments(t *testing.T) {
	ast, err := parse([]*Token{
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "name",
			Value: "foo",
		},
		{
			Type:  "integer",
			Value: "123",
		},
		{
			Type:  "integer",
			Value: "456",
		},
		{
			Type:  "paren",
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: "program", Body: []Node{
		{
			Type: "call",
			Name: "foo",
			Arguments: []Node{
				{
					Type:     "number",
					IntValue: 123,
				},
				{
					Type:     "number",
					IntValue: 456,
				},
			},
		},
	}}, ast)
}

func TestParseCallRecursive(t *testing.T) {
	ast, err := parse([]*Token{
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "name",
			Value: "foo",
		},
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "name",
			Value: "bar",
		},
		{
			Type:  "paren",
			Value: ")",
		},
		{
			Type:  "paren",
			Value: ")",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, Node{Type: "program", Body: []Node{
		{
			Type: "call",
			Name: "foo",
			Arguments: []Node{
				{
					Type:      "call",
					Name:      "bar",
					Arguments: []Node{},
				},
			},
		},
	}}, ast)
}
