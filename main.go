package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ********************************Example 4 🏂********************************

func main() {

	// str := "Ashiq.Hussain.Kumar"
	// arr := strings.Split(str, ".")
	// str = strings.Join(arr, "#")
	// fmt.Println(str, len(arr))

	if len(os.Args) < 3 {
		fmt.Fprint(os.Stderr, "not enough args")
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

// ********************************Example 3 🏂********************************

// func main() {
// 	elite := "ělite"
// 	fmt.Printf("%T %[1]v %d\n", elite, len(elite))
// 	fmt.Printf("%T %[1]v\n", []rune(elite))
// 	fmt.Printf("%T %[1]v\n", []byte(elite))
// }

// ********************************Example 2 🏂********************************

// func main() {
// 	var sum float64
// 	var n int

// 	for {
// 		var val float64

// 		_, err := fmt.Fscanln(os.Stdin, &val)

// 		if err != nil {
// 			break
// 		}

// 		sum += val
// 		n++
// 	}

// 	if n == 0 {
// 		fmt.Fprintln(os.Stderr, "no values")
// 		os.Exit(-1)
// 	}

// 	fmt.Println("The average is:", sum/float64(n))
// }

// ********************************Example 1 🏂********************************

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
