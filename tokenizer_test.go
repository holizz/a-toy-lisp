package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func compareTokens(t *testing.T, expected []Token, actual []*Token) {
	assert.Equal(t, len(expected), len(actual))
	for i := 0; i < len(expected); i++ {
		ex := expected[i]
		ac := actual[i]

		assert.Equal(t, ex.Type, ac.Type)
		assert.Equal(t, ex.Value, ac.Value)
	}
}

func TestTokenizerBlank(t *testing.T) {
	tokens, err := tokenizer("")
	assert.Nil(t, err)
	compareTokens(t, []Token{}, tokens)
}

func TestTokenizerParens(t *testing.T) {
	tokens, err := tokenizer("()")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "paren",
			Value: ")",
		},
	}, tokens)
}

func TestTokenizerName(t *testing.T) {
	tokens, err := tokenizer("(abc def)")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  "paren",
			Value: "(",
		},
		{
			Type:  "name",
			Value: "abc",
		},
		{
			Type:  "name",
			Value: "def",
		},
		{
			Type:  "paren",
			Value: ")",
		},
	}, tokens)
}

func TestTokensUnicode(t *testing.T) {
	tokens, err := tokenizer("\U0001F407")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  "name",
			Value: "\U0001F407",
		},
	}, tokens)
}

func TestTokensInteger(t *testing.T) {
	tokens, err := tokenizer("123 456")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  "integer",
			Value: "123",
		},
		{
			Type:  "integer",
			Value: "456",
		},
	}, tokens)
}
