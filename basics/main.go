package main

// EXAMPLE 08

// func main() {
// 	hashMap := make(map[string]int)
// 	names := []string{"ashiq", "amir", "zahid", "yawer", "bilal", "owais", "muzi"}

// 	for index, value := range names {
// 		hashMap[value] = index*index + index
// 	}
// 	for key, value := range hashMap {
// 		fmt.Printf("%v => %v\n", key, value)
// 	}
// }

// func main() {
// 	var arr [3]int
// 	for index, value := range arr {
// 		arr[index] = (value+1)*(value+2) + index*index
// 	}
// 	for index, value := range arr {
// 		fmt.Println(index, value)
// 	}
// }

// func main() {
// 	n := 10

// 	for i := 0; i < 4; i++ {
// 		for j := 0; j < 4; j++ {
// 			fmt.Printf("(%v, %v)\n", i, j)
// 		}
// 	}

// 	price, n := 4, n*10
// 	fmt.Println(n, price)
// }

// EXAMPLE 07
// program to count the top 3 words in a file

// func main() {
// 	scan := bufio.NewScanner(os.Stdin)
// 	words := make(map[string]int)

// 	scan.Split(bufio.ScanWords)

// 	for scan.Scan() {
// 		words[scan.Text()]++
// 	}
// 	fmt.Println("unique words in file:-", len(words))

// 	type kv struct {
// 		key string
// 		val int
// 	}

// 	var ss []kv

// 	for k, v := range words {
// 		ss = append(ss, kv{k, v})
// 	}

// 	sort.Slice(ss, func(i int, j int) bool {
// 		return ss[i].val > ss[j].val
// 	})
// 	for _, s := range ss[:3] {
// 		fmt.Println(s.key, "appears", s.val, "times")
// 	}
// }

// EXAMPLE 06
// i.ARRAYS ARE GETTING COPIED IN GO (Chunk of Memory Copied)
// ii.SLICES STRINGS and MAPS are refrenced (even when passing to functions also), their descriptor is going to be copied but the memory chunk they are
// pointing is getting refrenced (shared)

// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	b := a

// 	for i := 0; i < 4; i++ {
// 		b[i] += 10
// 	}
// 	fmt.Println(a, len(a))
// 	fmt.Println(b, len(b))
// }

// EXAMPLE 05

// func main() {
// 	if len(os.Args) < 3 {
// 		fmt.Fprintln(os.Stderr, "not enough args")
// 		os.Exit(-1)
// 	}

// 	old, new := os.Args[1], os.Args[2]
// 	scan := bufio.NewScanner(os.Stdin)

// 	for scan.Scan() {
// 		s := strings.Split(scan.Text(), old)
// 		t := strings.Join(s, new)

// 		fmt.Println(t)
// 	}
// }

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
