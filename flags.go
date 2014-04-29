package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("flagname", 1234, "Help message for flagname")
	// flag.Parse()
	fmt.Println("ip has value", *ip)

}
