// Package textstat conta palavras e linhas de um texto.
package textstat

import "strings"

// CountWords devolve o número de palavras de s, separadas por espaço em
// branco Unicode.
func CountWords(s string) int {
	return len(strings.Fields(s))
}

// CountLines devolve o número de linhas de s.
func CountLines(s string) int {
	return strings.Count(s, "\n")
}
