package main

import (
	"fmt"
	"os"
)

// ********************************Example 2********************************

func main() {
	var sum float64
	var n int

	for {
		var val float64

		_, err := fmt.Fscanln(os.Stdin, &val)

		if err != nil {
			break
		}

		sum += val
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "no values")
		os.Exit(-1)
	}

	fmt.Println("The average is:", sum/float64(n))
}

// ********************************Example 1********************************

// import (
// 	"fmt"
// 	"os"
// 	"test/hello"
// )

// func main() {
// 	// fmt.Println(hello.Say(os.Args[1:]))
// 	a := 22
// 	b := 3.5

// 	fmt.Printf("a: %T %v\n", a, a)
// 	fmt.Printf("b: %T %v\n", b, b)

// 	fmt.Printf("a: %8T %[1]v\n", a)
// 	fmt.Printf("b: %8T %[1]v\n", b)

// 	a = int(b)
// 	b = float64(a)
// 	fmt.Printf("a: %8T %[1]v\n", a)
// 	fmt.Printf("a: %8T %[1]v\n", b)
// }
