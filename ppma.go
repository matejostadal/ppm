/*
MATEJ OSTADAL
KMI UPOL
2024
*/

package main

import (
	"errors"
	"fmt"
)

var K int = 3

func main() {

	// EXAMPLE1
	// tested_string := "barbaraabarboraubaru"

	// appropriate alphabet for our example
	// alphabet := map[rune]bool{
	// 	'a': true,
	// 	'b': true,
	// 	'r': true,
	// 	'o': true,
	// 	'u': true,
	// 	'ε': true,
	// }

	// ------------
	// EXAMPLE 2
	tested_string := "pepapapapapu"

	alphabet := map[rune]bool{
		'p': true,
		'e': true,
		'a': true,
		'u': true,
		'ε': true,
	}

	// RUN
	context_map := ppm_a(tested_string, K, alphabet)
	print(context_map)
}

func ppm_a(input_string string, K int, alphabet map[rune]bool) map[string]int {
	// performing PPMA context analysis

	input := []rune(input_string)
	input_length := len(input)

	var symbol rune = 0
	// prepare map of contexts
	context_map := init_contexts(alphabet)

	for index := 0; index < input_length; index += 1 {
		symbol = input[index]

		fmt.Printf("SYMBOL READ: %s\n", string(symbol))

		// for correctness (optional)
		if !is_contained(symbol, alphabet) {
			panic(errors.New("given text contains a symbol that is not in the given alphabet"))
		}
		
		fmt.Printf("OUT: ")

		// checking occurence in contexts appearing before symbol being read
		for back_step := min(K, index); back_step >= 0; back_step -= 1 {

			// get recent context from input 
			context := string(input[index-back_step:index])
			// check the corresponding occurence count
			_, not_zero := context_map[string(symbol) + "|" + context]

			// OUTPUT
			if not_zero {
				fmt.Printf("<" + string(symbol) + ", " + context + "> ")
			} else {
				fmt.Printf("<" + string('ε') + ", " + context + "> ")
				if context == "" {
					fmt.Printf("<" + string(symbol) + ", c-1> ")
				}
			}
			// remember new occurence of symbol
			context_map[string(symbol) + "|" + context] += 1

		}
		fmt.Printf("\n\n")
	}
	return context_map
}

func is_contained(symbol rune, alphabet map[rune]bool) bool {
	// simply returns if symbol is contained in an alphabet (abstraction tool)
	_, ok := alphabet[symbol]
	return ok
}

func init_contexts(alphabet map[rune]bool) map[string]int {
	// sets the occurence count in the empty context to 1
	contexts := map[string]int{}

	for symbol, _ := range alphabet {
		contexts[string(symbol) + "|c-1"] = 1
	}

	return contexts
}