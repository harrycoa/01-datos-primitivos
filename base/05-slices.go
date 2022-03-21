package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(" ****************************** Condicionales ******************************")

	// 1.- Arrays
	var array [4]int

	array[0] = 1
	array[1] = 22

	fmt.Println(array)

	// Len = tamño array --- Cap = maxima capacidad array
	fmt.Println(array, len(array), cap(array))

	// 2.- Slices
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))

	// 3.1.- Slicing

	fmt.Println(slice, len(slice), cap(slice))

	// 4.- Recorriendo slices and arrys
	slicito := []string{"hola", "que", "ase"}

	for _, valor := range slicito {
		fmt.Println(valor)
	}

	// Funcion Palindromo
	isPalindromo("Anita Lava La tina")
	isPalindromo2("Anita Lava La tina")
	isPalindromo3("Anita Lava La tina")

	if isPalindromo3("reconocer") {
		fmt.Println("Is palindrome")
	} else {
		fmt.Println("Is not palindrome")
	}

}

func isPalindromo(text string) {
	var textReverse string
	text = strings.ToLower(text)
	fmt.Println(text)

	text = quitarEspacios(text)

	fmt.Println(text)

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		fmt.Println(" Si es Palindromo")
	} else {
		fmt.Println(" NO es Palindromo")
	}

}

func quitarEspacios(input string) string {
	transformed := func(r rune) rune {
		if r == ' ' {
			return 0
		}
		return r
	}
	return strings.Map(transformed, input)
}

func isPalindromo2(text string) {
	var n int = len(text) - 1
	for i := 0; i < n/2; i++ {
		if text[i] != text[n-1] {
			fmt.Println("No es palindromo")
			return
		}
	}
	fmt.Println("es palindromo ")
}

func isPalindromo3(text string) bool {
	palindrome := strings.ReplaceAll(strings.ToLower(text), " ", "")
	var backwardWord string
	for _, v := range palindrome {
		backwardWord = string(v) + backwardWord
	}
	return palindrome == backwardWord
}
