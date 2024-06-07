package pkg2

import (
	"fmt"
	// _ "mymodule03/pkg3"
)

var (
	_ = constInitCheck()
	v = variableInit("v")
)

const c = "c"

func constInitCheck() string {
	if c != "" {
		println("pkg2: const c has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg2: var %s has been initialized\n", name)
	return name
}

func init() {
	println("pkg2: init func invoked")
}
