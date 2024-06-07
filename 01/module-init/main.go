package main

//首先会按照依赖关系依次初始化pkg1、pkg2包，然后才会初始化main包。
import (
	"fmt"
	_ "mymodule03/pkg1"
	_ "mymodule03/pkg2"
)

const (
	c1 = "c1"
	c2 = "c2"
)

var (
	x  = constInitCheck()
	v1 = variableInit("v1")
	v2 = variableInit("v2")
)

func constInitCheck() string {
	if c1 != "" {
		fmt.Println("main: const c1 has been initialized")
	}
	if c2 != "" {
		fmt.Println("main: const c2 has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("main: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("main: first init func invoked")
}

func init() {
	fmt.Println("main: second init func invoked")
}

func main() {
	// init() // ERROR undefined: init
}
