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
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
	}, tokens)
}

func TestTokenizerName(t *testing.T) {
	tokens, err := tokenizer("(abc def)")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  TokenTypeParen,
			Value: "(",
		},
		{
			Type:  TokenTypeName,
			Value: "abc",
		},
		{
			Type:  TokenTypeName,
			Value: "def",
		},
		{
			Type:  TokenTypeParen,
			Value: ")",
		},
	}, tokens)
}

func TestTokensUnicode(t *testing.T) {
	tokens, err := tokenizer("\U0001F407")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  TokenTypeName,
			Value: "\U0001F407",
		},
	}, tokens)
}

func TestTokensInteger(t *testing.T) {
	tokens, err := tokenizer("123 456")
	assert.Nil(t, err)
	compareTokens(t, []Token{
		{
			Type:  TokenTypeInteger,
			Value: "123",
		},
		{
			Type:  TokenTypeInteger,
			Value: "456",
		},
	}, tokens)
}
