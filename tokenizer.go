package main

import "regexp"

func tokenizer(program string) ([]*Token, error) {
	tokens := []*Token{}

	for _, _char := range program {
		char := string(_char)
		previous := &Token{}
		if len(tokens) > 0 {
			previous = tokens[len(tokens)-1]
		}

		switch {
		case char == "(" || char == ")":
			tokens = append(tokens, &Token{
				Type:  TokenTypeParen,
				Value: char,
			})
		case char == " ":
			previous.Finished = true
		case regexp.MustCompile("[0-9]").MatchString(char):
			if previous.Type == TokenTypeInteger && previous.Finished == false {
				previous.Value += char
			} else {
				tokens = append(tokens, &Token{
					Type:  TokenTypeInteger,
					Value: char,
				})
			}
		default:
			// Treat everything as a name unless otherwise specified
			if previous.Type == TokenTypeName && previous.Finished == false {
				previous.Value += char
			} else {
				tokens = append(tokens, &Token{
					Type:  TokenTypeName,
					Value: string(char),
				})
			}
		}
	}

	return tokens, nil
}

type Token struct {
	Type     TokenType
	Value    string
	Finished bool
}

type TokenType int

const (
	TokenTypeParen TokenType = iota
	TokenTypeInteger
	TokenTypeName
)
