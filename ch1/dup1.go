//Dup1 prints the text of each line that apppears more than once
// in the standard input, preceded by its count

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// the built in function make creates a new empty map, in this case
	// with the key as a string, and the value as an int
	counts := make(map[string]int)
	// bufio implements buffered I/O, wrapping an io.Reader or io.Writer object and creating
	// another object (Reeader or Writer ) that also implements the interface but provides buffering and help
	// for textual I/O
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential erros from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
