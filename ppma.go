/*
MATEJ OSTADAL
KMI UPOL
2024
*/

package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var K int = 3

func main() {

	// // EXAMPLE1
	// tested_string := "barbaraabarboraubaru"

	// // appropriate alphabet for our example
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

	// RUN ENCODE
	fmt.Printf("IN -  String being encoded: %s\n\n", tested_string)
	context_map_in, output_string := ppm_a_in(tested_string, K, alphabet)

	fmt.Printf("STEP - Encoded output: %s\n\n", output_string)

	// RUN DECODE
	context_map_out, got_back_string := ppma_a_out(output_string, alphabet)
	fmt.Printf("OUT - String got from decoding: %s\n\n", got_back_string)

	fmt.Printf("CHECK - Result of two context maps deep equality: %t\n\n", reflect.DeepEqual(context_map_in, context_map_out))	
}

func ppm_a_in(input_string string, K int, alphabet map[rune]bool) (map[string]int, string)  {
	// performing PPMA encoding

	input := []rune(input_string)

	var context_string string = ""
	var output_string strings.Builder
	// prepare map of contexts
	context_map := init_contexts(alphabet)

	for index, symbol := range(input){
	
		// fmt.Printf("SYMBOL READ: %s\n", string(symbol))

		// for correctness (optional)
		if !is_contained(symbol, alphabet) {
			panic(errors.New("given text contains a symbol that is not in the given alphabet"))
		}
		
		// fmt.Printf("OUT: ")

		// checking occurence in contexts appearing before symbol being read
		for back_step := min(K, index); back_step >= 0; back_step -= 1 {

			// get recent context from input 
			context := string(input[index-back_step:index])
			// check the corresponding occurence count
			_, not_zero := context_map[string(symbol) + "|" + context]

			// OUTPUT
			if not_zero {
				context_string = "<" + string(symbol) + "," + context + ">"

				output_string.WriteString(context_string)
				// fmt.Printf(context_string + " ")
			} else {
				context_string = "<" + string('ε') + "," + context + ">"

				output_string.WriteString(context_string)
				// fmt.Printf(context_string + " ")
			
				if context == "" {
					context_string = "<" + string(symbol) + ",-1>"
					output_string.WriteString(context_string)
					// fmt.Printf(context_string + " ")
				}
			}
			// remember new occurence of symbol
			context_map[string(symbol) + "|" + context] += 1
		}
		// fmt.Printf("\n\n")
	}
	return context_map, output_string.String()
}

func ppma_a_out(output_string string, alphabet map[rune]bool) (map[string]int, string) {
	// performing PPMA decoding

	// split the output string into bracketed parts
	output := strings.Split(output_string, "<")[1:]

	var part []string
	var current_contexts []string
	var input_string strings.Builder
	// prepare map of contexts - different initialization than in encoding
	context_map := map[string]int{"ε|-1": 1,}

	out_buffer := ""

	for _, out_part := range(output) {
		// read one bracket produced while encoding
		part = strings.Split(out_part, ",")

		// split the bracket
		symbol := part[0]
		context := strings.TrimSuffix(part[1], ">")

		// for correctness (optional)
		if !is_contained([]rune(symbol)[0], alphabet) {
			panic(errors.New("given text contains a symbol that is not in the given alphabet"))
		}

		// if we are reading a new symbol, the symbol in the buffer goes into the output string
		if symbol != out_buffer {
			input_string.WriteString(out_buffer)
			out_buffer = ""
		}
		
		if symbol != "ε" {
			// new adept for output string (but still can be found later)
			out_buffer = symbol
			
			// new context
			current_contexts = append(current_contexts, context)

			// increase symbol occurence in all currently stored contexts + empty the contexts
			store_contexts(symbol, current_contexts, context_map)
			current_contexts = current_contexts[:0]

		} else {
			// just remember the context that occured and skip
			current_contexts = append(current_contexts, context)
			continue
		}
	}
	// empty buffer
	input_string.WriteString(out_buffer)

	return context_map, input_string.String()
}

func store_contexts(symbol string, current_contexts []string, context_map map[string]int) {
	// increases counter in context map for given symbol in all given contexts
	for _, context := range(current_contexts) {
		context_map[symbol + "|" + context] += 1
	}
}

func is_contained(symbol rune, alphabet map[rune]bool) bool {
	// simply returns if symbol is contained in an alphabet (abstraction tool)
	_, ok := alphabet[symbol]
	return ok
}

func init_contexts(alphabet map[rune]bool) map[string]int {
	// sets the occurence count in the empty context to 1
	contexts := map[string]int{}

	for symbol := range alphabet {
		contexts[string(symbol) + "|-1"] = 1
	}

	return contexts
}