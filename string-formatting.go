package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// fmt.Printf("%+v\n", p)

	// fmt.Printf("%#v\n", p)

	// fmt.Printf("%T\n", p)

	// fmt.Printf("%t\n", true)

	// fmt.Printf("%d\n", 123)

	// fmt.Printf("%b\n", 14)

	fmt.Printf("%c\n", 33)

	// fmt.Printf("%x\n", 456)
	// fmt.Printf("%e\n", 123400000.0)
	// fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// To left-justify, use the - flag.
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// To left-justify use the - flag as with numbers.
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")

}
