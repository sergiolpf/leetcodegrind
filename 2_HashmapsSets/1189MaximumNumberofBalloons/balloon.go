package main

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("leetcode"))
}

func maxNumberOfBalloons(text string) int {
	const BALLON string = "balloon"

	m_ballon := make(map[rune]int, len(BALLON))

	m_text := make(map[rune]int, len(text))

	for _, letter := range BALLON {
		m_ballon[letter]++
	}

	for _, letter := range text {
		m_text[letter]++
	}

	numeroDeX := -1
	for letter, qtd := range m_ballon {
		if t_qtd, ok := m_text[letter]; ok {
			if numeroDeX < 0 || t_qtd/qtd < numeroDeX {
				numeroDeX = t_qtd / qtd
			}
		} else {
			return 0
		}

	}

	return numeroDeX
}
