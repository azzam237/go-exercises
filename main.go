package main

import (
	"fmt"
	"go-exercise/modules"
)

func main() {
	var option string
	fmt.Scanf("%s\n", &option)

	if option == "1" {
		fmt.Println(modules.Astring)
	}

}
