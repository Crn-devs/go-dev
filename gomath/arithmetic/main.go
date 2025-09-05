package main

import (

	// notebook-only: markdown/latex output

	"fmt"

	"github.com/janpfeifer/gonb/gonbui"
)

func main() {

	x, y := 5, 4
	result := 5 + 4

	eq := fmt.Sprintf("$%v + %v = %v$", x, y, result)

	// quick demo so you can test
	gonbui.DisplayMarkdown(eq)
}
