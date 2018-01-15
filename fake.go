package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Binary say Hello World1")
	fmt.Fprintf(os.Stderr, "Std err 1\n")
	fmt.Println("Binary say Hello World2")
	fmt.Fprintf(os.Stderr, "Std err 2\n")
	fmt.Println("Binary say Hello World3")
	fmt.Fprintf(os.Stderr, "Std err 3\n")
	fmt.Println("Binary say Hello World4")
	fmt.Fprintf(os.Stderr, "Std err 4\n")
	fmt.Println("Binary say Hello World5")
	fmt.Fprintf(os.Stderr, "Std err 5\n")
	fmt.Println("Binary say Hello World6")
	fmt.Fprintf(os.Stderr, "Std err 6\n")
	fmt.Println("Binary say Hello World7")
	fmt.Fprintf(os.Stderr, "Std err 7\n")
	fmt.Println("Binary say Hello World8")
	fmt.Fprintf(os.Stderr, "Std err 8\n")
	fmt.Println("Binary say Hello World9")
	fmt.Fprintf(os.Stderr, "Std err 9\n")
}
