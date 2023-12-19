package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// EXAMPLE 05

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)
		t := strings.Join(s, new)

		fmt.Println(t)
	}
}

// EXAMPLE 04
// GO STRINGS ARE UTF-8 ENCODING OF UNICODE CHARACTERS

// func main() {
// 	s := "ělite"
// 	fmt.Printf("%T %[1]v %d\n", s, len(s))
// 	fmt.Printf("%T %[1]v %d\n", []rune(s), len(s))
// 	fmt.Printf("%T %[1]v %d\n", []byte(s), len(s))
// }

// EXAMPLE 03

// func main() {
// 	var sum float64
// 	var n int

// 	for {
// 		var val float64

// 		_, err := fmt.Fscanln(os.Stdin, &val)
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}

// 		sum += val
// 		n++
// 	}

// 	if n == 0 {
// 		fmt.Fprintln(os.Stderr, "no values")
// 		os.Exit(-1)
// 	}

// 	fmt.Println("The average is", sum/float64(n))
// }

// EXAMPLE 02

// func main() {
// 	a := 10
// 	b := 3.4

// 	fmt.Printf("a: %8T %v\n", a, a)
// 	fmt.Printf("b: %8T %v\n", b, b)

// 	fmt.Printf("a: %8T %[1]v\n", a)
// 	fmt.Printf("b: %8T %[1]v\n", b)

// 	a = int(b)
// 	b = float64(a)

// 	fmt.Printf("a: %8T %[1]v\n", a)
// 	fmt.Printf("b: %8T %[1]v\n", b)
// }

// EXAMPLE 01

// func main() {
// 	fmt.Println("Hello World!")
// }
