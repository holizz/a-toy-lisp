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
				Type:  "paren",
				Value: char,
			})
		case char == " ":
			previous.Finished = true
		case regexp.MustCompile("[0-9]").MatchString(char):
			if previous.Type == "integer" && previous.Finished == false {
				previous.Value += char
			} else {
				tokens = append(tokens, &Token{
					Type:  "integer",
					Value: char,
				})
			}
		default:
			// Treat everything as a name unless otherwise specified
			if previous.Type == "name" && previous.Finished == false {
				previous.Value += char
			} else {
				tokens = append(tokens, &Token{
					Type:  "name",
					Value: string(char),
				})
			}
		}
	}

	return tokens, nil
}

type Token struct {
	Type     string
	Value    string
	Finished bool
}
