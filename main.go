package main

import (
	"fmt"

	"github.com/ikaem/stringmod/quotes"
	"github.com/ikaem/stringmod/strings"
)

func main() {

	o, e := strings.CountOddEven("12345")

	fmt.Println(o, e)
	fmt.Println(quotes.GetEmoji("smiley"))

}
