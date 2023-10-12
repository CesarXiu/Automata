package main

import (
	"fmt"
	"regexp"
)

// Definición de los tokens
type Token struct {
	TokenType string
	Value     string
}

// Función para analizar léxicamente la entrada
func lexer(input string) []Token {
	var tokens []Token

	// Expresiones regulares para cada tipo de token
	tokenPatterns := map[string]string{
		"NUMERO":   "[0-9]+",
		"OPERADOR": "[+*-/]",
		"ESPACIO":  "\\s+",
	}

	// Construir la expresión regular combinada
	pattern := ""
	for _, p := range tokenPatterns {
		pattern += fmt.Sprintf("|(%s)", p)
	}

	// Iterar sobre los tokens encontrados
	re := regexp.MustCompile(pattern[1:])
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		for i, name := range re.SubexpNames()[1:] {
			if match[i+1] != "" {
				tokens = append(tokens, Token{name, match[i+1]})
			}
		}
	}

	return tokens
}

func main() {
	// Texto de entrada
	input := "3 + 4 * 2 - 1"

	// Analizar léxicamente la entrada
	tokens := lexer(input)

	// Imprimir los tokens encontrados
	for _, token := range tokens {
		fmt.Printf("Token: %s, Valor: %s\n", token.TokenType, token.Value)
	}
}
