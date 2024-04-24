/*
MATEJ OSTADAL
KMI UPOL
2024
*/

package main

import "errors"

var K int = 3

func main() {

	// appropriate alphabet for our example
	alphabet := map[rune]bool{
		'a': true,
		'b': true,
		'r': true,
		'o': true,
		'u': true,
	}

	// example input
	tested_string := "barbaraabarboraubaru"

	ppm_a(tested_string, K, alphabet)
}

func ppm_a(input string, K int, alphabet map[rune]bool) int {

	// todo promyslet reprezentaci kontextů (resp. počtu výskytů v daném kontextu)
	// pravděpodobně vyzkoušet map indexovanou dle kontextu (slice původního stringu input)
	// aby to nebylo neefektivní... promyslet souvislost se stromem (pokud mám už spočítaný a uložený kontext, můžu to využít)

	input_length := len(input)
	var a uint8 = 0
	symbols_read := 0


	for symbols_read < input_length {
		a = input[symbols_read]

		// for correctness (optional)
		if !is_contained(a, alphabet) {
			panic(errors.New("Given text contains a symbol that is not contained in the given alphabet"))
		}

		symbols_read += 1
	}

	return 0
}

func is_contained(symbol uint8, alphabet map[rune]bool) bool {
	_, ok := alphabet[rune(symbol)]
	return ok
}
