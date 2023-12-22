package main

// EXAMPLE 14

// var raw = `
// <html>
// <body>
//     <h1>First Heading</h1>
//     <p>First Paragraph</p>
//     <p>HTML images are defined with the img tag:</p>
//     <img src="abc.jpg" width="100" height="100">
// </body>
// </html>
// `

// func visit(node *html.Node, words *int, pics *int) {

// 	fmt.Println("Node:-", node.Data)
// 	if node.Type == html.TextNode {
// 		(*words) += len(strings.Fields(node.Data))
// 	} else if node.Type == html.ElementNode && node.Data == "img" {
// 		(*pics)++
// 	}
// 	for c := node.FirstChild; c != nil; c = c.NextSibling {
// 		visit(c, words, pics)
// 	}
// }

// func countWordsAndImages(doc *html.Node) (int, int) {
// 	var words, pics int
// 	visit(doc, &words, &pics)
// 	return words, pics
// }

// func main() {
// 	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "parse failed: %s", err)
// 		os.Exit(-1)
// 	}
// 	words, pics := countWordsAndImages(doc)
// 	fmt.Printf("#words:%d & #pics:%d\n", words, pics)
// }

// EXAMPLE 13

// func main() {
// 	a := [3]int{1, 2, 3}

// 	b := a[:2]
// 	c := a[:2:2]
// 	fmt.Printf("A:%d %d %p %v\n", len(a), cap(a), &a, a)
// 	fmt.Printf("B:%d %d %p %v\n", len(b), cap(b), b, b)
// 	fmt.Printf("C:%d %d %p %v\n", len(c), cap(c), c, c)

// 	b = append(b, -10)
// 	a[0] = -78
// 	c = append(c, 99)
// 	fmt.Printf("A:%d %d %p %v\n", len(a), cap(a), &a, a)
// 	fmt.Printf("B:%d %d %p %v\n", len(b), cap(b), b, b)
// 	fmt.Printf("C:%d %d %p %v\n", len(c), cap(c), c, c)
// }

// EXAMPLE 12
// Closures

// func main() {
// 	s := make([]func(), 4)

// 	for i := 0; i < 4; i++ {
// 		i2 := i // closure capture
// 		s[i] = func() {
// 			fmt.Printf("%d %p\n", i2, &i2)
// 		}
// 	}

// 	for i := 0; i < 4; i++ {
// 		s[i]()
// 	}
// }

// func fib() func() int {
// 	a, b := 0, 1

// 	return func() int {
// 		a, b = b, a+b
// 		return b
// 	}
// }

// func main() {
// 	f, g := fib(), fib()
// 	fmt.Println(f(), f(), f(), f(), f())
// 	fmt.Println(g(), g(), g(), g(), g())
// }

// EXAMPLE 11
// defer GOTCHAS

// func main() {
// 	a := 10
// 	defer fmt.Println(&a)
// 	a = a * a * a * a
// 	fmt.Println(&a)
// }

// EXAMPLE 10

// func do(a1 *[2]int) {
// 	a1[0] = 99
// 	a1[1] = -100
// }

// func main() {
// 	a := [2]int{1, 2}
// 	fmt.Println(a)
// 	do(&a)
// 	fmt.Println(a)
// }

// func do(m1 *map[int]int) {
// 	(*m1)[9] = 99
// 	(*m1) = make(map[int]int)
// }

// func main() {
// 	m := map[int]int{
// 		1: 11,
// 		2: 22,
// 		3: 33,
// 	}
// 	fmt.Println(m)
// 	do(&m)
// 	fmt.Println(m)
// }

// func do(m1 map[int]int) {
// 	m1[9] = 99
// 	m1 = make(map[int]int)
// }

// func main() {
// 	m := map[int]int{
// 		1: 11,
// 		2: 22,
// 		3: 33,
// 		4: 44,
// 	}
// 	fmt.Println(m)
// 	do(m)
// 	fmt.Println(m)
// }

// EXAMPLE 09
// Line count, word count, char count for a file

// func main() {
// 	for _, fname := range os.Args[1:] {
// 		var lc, wc, cc int
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		scan := bufio.NewScanner(file)

// 		for scan.Scan() {
// 			s := scan.Text()

// 			wc += len(strings.Fields(s))
// 			cc += len(s)
// 			lc++
// 		}
// 		fmt.Printf("%5d %5d %5d %s\n", lc, wc, cc, fname)
// 	}
// }

// EXAMPLE 09
// Concat file content to the stdout

// func main() {
// 	for _, fname := range os.Args[1:] {
// 		file, err := os.Open(fname)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		// if _, err := io.Copy(os.Stdout, file); err != nil {
// 		// 	fmt.Fprintln(os.Stderr, err)
// 		// 	continue
// 		// }

// 		data, err := io.ReadAll(file)

// 		if err != nil {
// 			fmt.Fprintln(os.Stderr, err)
// 			continue
// 		}

// 		fmt.Println(string(data)) // convert byte slice to string
// 		fmt.Println("The file has", len(data), "bytes")
// 		file.Close()
// 	}
// }

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
