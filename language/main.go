package main

import (
	"fmt"
	"github.com/comasz/GoLang-Learning/language/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, whats new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie ghet es innen"))
}
