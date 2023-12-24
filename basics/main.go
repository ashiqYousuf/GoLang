package main

/*
We can't take address of a MAP entry.
Also if I have a map of structs & i wanna do something to a value inside that struct, I can't do that.
So you are always gonna see map of key to struct pointers
Ex:- map[string]*Employee

We can't take address of a MAP or SLICE entry, why?
Because slices can get reallocated & we are then storing some stale address, in case of maps, maps keep on changing
(rearranging themselves here and there i.e no order), so we too might keep pointing to some stale pointer variable.
AVOID:- &slice[0] or %map[key]

Do not capture refrence to a Loop variable

// OOPS
// IN GO WE CAN PUT METHODS ON ANY USER DECLARED TYPE
// NOTE THAT I CAN ASSIGN ANYTHING TO THE INTERFACE THAT SATISFIES THE INTERFACE (ANY TYPE THAT HAS ALL INTERFACE METHODS)
// A METHOD IS A FUNCTION ASSOCIATED WITH A TYPE
*/

// GO OOPS START

// EXAMPLE 06

// type Text string

// func (t *Text) Write(p []byte) (int, error) {
// 	*t = Text(string(p))
// 	return len(p), nil
// }

// func main() {
// 	var t Text
// 	if len(os.Args) > 1 {
// 		file, _ := os.Open(os.Args[1])
// 		io.Copy(&t, file)
// 	}
// 	fmt.Printf("%v\n", string(t))
// 	fmt.Printf("%d bytes written\n", len(t))
// }

// EXAMPLE 05

// type Point struct {
// 	X, Y float64
// }

// type Line struct {
// 	Begin, End Point
// }

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
// }

// func (l Line) ScaleBy(f float64) Line {
// 	l.End.X += (f - 1) * (l.End.X - l.Begin.X)
// 	l.End.Y += (f - 1) * (l.End.Y - l.Begin.Y)
// 	return Line{l.Begin, Point{l.End.X, l.End.Y}}
// }

// func main() {
// 	l1 := Line{Point{1, 2}, Point{4, 6}}
// 	l2 := l1.ScaleBy(2.5)
// 	fmt.Println(l1.Distance())
// 	fmt.Println(l2.Distance())
// 	// Method chaining, only possible as ScaleBy() is returning Line
// 	fmt.Println(Line{Point{1, 2}, Point{4, 6}}.ScaleBy(10).Distance())
// }

// EXAMPLE 04 (IMPORTANT!)

// type Point struct {
// 	X, Y float64
// }

// type Line struct {
// 	Begin, End Point
// }

// type Path []Point

// func (l Line) Distance() float64 {
// 	return math.Hypot(l.End.X-l.Begin.X, l.End.Y-l.Begin.Y)
// }

// func (p Path) Distance() (sum float64) {
// 	for i := 1; i < len(p); i++ {
// 		sum += Line{p[i-1], p[i]}.Distance()
// 	}
// 	return sum
// }

// type Distancer interface {
// 	Distance() float64
// }

// // Line and Path both UDT satisfy the Distancer interface as both have the Distance() method
// func PrintDistance(d Distancer) {
// 	fmt.Println(d.Distance())
// }

// func main() {
// 	l := Line{Point{1, 2}, Point{4, 6}}
// 	path := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
// 	PrintDistance(l)    // l.DIstance()
// 	PrintDistance(path) // path.Distance()
// }

// EXAMPLE 03 (IMPORTANT EXAMPLE)

// io.Writer is an interface with Write(p []byte) method
// io.Reader is an interface with Read(p []byte) method
// in io.Copy(dest, src), anything with the Read() method can be the src
// and anything with the Write() method can be the dest

// type ByteCounter int // concrete type, providing behaviour of a Writer

// func (b *ByteCounter) Write(p []byte) (int, error) {
// 	// I need a pointer to the receiver (ByteCounter) as i need to modify it
// 	l := len(p)
// 	*b += ByteCounter(l)
// 	return l, nil
// }

// func main() {
// 	var c ByteCounter
// 	f1, _ := os.Open("ip.txt")
// 	// f2, _ := os.Create("out.txt")
// 	f2 := &c

// 	n, _ := io.Copy(f2, f1) // io.Copy(dest io.Writer, interface src io.Reader interface)

// 	fmt.Println("copied", n, "bytes")
// 	fmt.Println(c)
// }

// EXAMPLE 02
//  A method is a function associated with a type
// Both Offset and Move are methods because they have a receiver of type Point
// and *Point, respectively. They operate on instances of the Point type.

// type Point struct {
// 	x float64
// 	y float64
// }

// func (p Point) Offset(x float64, y float64) Point {
// 	return Point{p.x + x, p.y + y}
// }

// func (p *Point) Move(x float64, y float64) {
// 	p.x += x
// 	p.y += y
// }

// func main() {
// 	p := Point{
// 		x: 0,
// 		y: 0,
// 	}
// 	fmt.Println(p)
// 	p1 := p.Offset(13, 13)
// 	fmt.Println(p1)
// 	fmt.Println(p)
// 	p.Move(4, 5)
// 	fmt.Println(p)
// }

// EXAMPLE 01

// type IntSlice []int

// func (is IntSlice) String() string {
// 	var strs []string

// 	for _, v := range is {
// 		strs = append(strs, strconv.Itoa(v))
// 	}
// 	return "[" + strings.Join(strs, ";") + "]"
// }

// func main() {
// 	var v IntSlice = []int{1, 2, 3}

// 	// s is an interface, what can I assign to an interface?
// 	// I can assign to it anything that satisfies the interface (any actual type that has a String method here!)
// 	var s fmt.Stringer = v

// 	for i, v := range v {
// 		fmt.Printf("v[%d]: %d\n", i, v)
// 	}
// 	fmt.Printf("%T %[1]v\n", v)
// 	fmt.Printf("%T %[1]v\n", s)
// }

// OOPS END

// *************************************************************************************************************************
// *************************************************************************************************************************
// *************************************************************************************************************************
// *************************************************************************************************************************
// *************************************************************************************************************************

// EXAMPLE 24
// Client Server

// type todo struct {
// 	UserID    int    `json:"userID"`
// 	ID        int    `json:"id"`
// 	Title     string `json:"title"`
// 	Completed bool   `json:"completed"`
// }

// var form = `
// <h1>Todo #{{.ID}}</h1>
// <div style="color:green">{{.UserID}}</div>
// <div>{{.Title}}</div>
// <h3>Completed {{.Completed}}</h3>
// `

// func handler(w http.ResponseWriter, r *http.Request) {
// 	const base = "https://jsonplaceholder.typicode.com/"
// 	resp, err := http.Get(base + r.URL.Path)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	// body, err := io.ReadAll(resp.Body)

// 	var item todo
// 	// err = json.Unmarshal(body, &item)

// 	if json.NewDecoder(resp.Body).Decode(&item); err != nil {
// 		http.Error(w, err.Error(), http.StatusServiceUnavailable)
// 		return
// 	}
// 	tmp1 := template.New("mine")
// 	tmp1.Parse(form)
// 	tmp1.Execute(w, item)
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// EXAMPLE 23
// Http Server

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Body: ", r.Body)
// 	fmt.Println("Header: ", r.Header)
// 	fmt.Println("Form: ", r.Form)
// 	fmt.Println("Method: ", r.Method)
// 	fmt.Println("Host", r.Host)
// 	fmt.Println("Query", r.URL.Query())
// 	fmt.Println("Raw Query", r.URL.RawQuery)
// 	fmt.Fprintf(w, "Hello, World from %s\n", r.URL.Path)
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// EXAMPLE 22
// REGEX

// func main() {
// 	text := "a aba abba abbba abbbba"
// 	re := regexp.MustCompile("b+")

// 	res := re.FindAllString(text, -1)
// 	for _, v := range res {
// 		fmt.Println(v)
// 	}
// 	fmt.Println(res)
// }

// EXAMPLE 21
// Slices gotcha & fix like closures

// func main() {
// 	a := [3]int{1, 2, 3}
// 	b := a[:1]
// 	c := a[:2]

// 	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
// 	fmt.Printf("%d %d %v\n", len(b), cap(b), b)
// 	fmt.Printf("%d %d %v\n", len(c), cap(c), c)
// 	fmt.Println("------------------------------------------------------")
// 	c = append(c, 9)
// 	fmt.Printf("%d %d %v\n", len(a), cap(a), a)
// 	fmt.Printf("%d %d %v\n", len(b), cap(b), b)
// 	fmt.Printf("%d %d %v\n", len(c), cap(c), c)

// 	for i, _ := range a {
// 		fmt.Printf("%d %p\n", a[i], &a[i])
// 	}
// 	fmt.Println("***************************************")
// 	for i, _ := range b {
// 		fmt.Printf("%d %p\n", b, &b[i])
// 	}
// 	fmt.Println("***************************************")
// 	for i, _ := range c {
// 		fmt.Printf("%d %p\n", c[i], &c[i])
// 	}
// }

// func main() {
// 	a := [][2]int{{1, 2}, {3, 4}, {5, 6}}
// 	b := make([][]int, 0, 3)

// 	for _, item := range a {
// 		i := make([]int, len(item))
// 		copy(i, item[:]) // make unique
// 		b = append(b, i[:])
// 	}
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// EXAMPLE 21

// func fib() func() int {
// 	a, b := 0, 1

// 	return func() int {
// 		a, b = b, a+b
// 		return b
// 	}
// }

// func main() {
// 	f := fib()
// 	g := fib()
// 	fmt.Printf("%d %d %d\n", f(), f(), f())
// 	fmt.Printf("%d %d %d\n", g(), g(), g())
// 	fmt.Println(f(), g())
// }

// func print(arr []*int) {
// 	for i, _ := range arr {
// 		fmt.Printf("%d ", *arr[i])
// 	}
// 	fmt.Println()
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	b := []*int{}

// 	for _, val := range a {
// 		b = append(b, &val)
// 	}
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	print(b)
// }

// EXAMPLE 20

// Stale slice pointer problem

// func main() {
// 	a := []int{1}
// 	p := &a[0]
// 	// *(p) = -90
// 	fmt.Printf("%d %d %v %[3]p\n", len(a), cap(a), a)
// 	a = append(a, 2)
// 	*(p) = -90
// 	fmt.Printf("%d %d %v %[3]p\n", len(a), cap(a), a)
// }

// EXAMPLE 19
// JSON & Structs

// type Response struct {
// 	Page  int      `json:"page"`
// 	Words []string `json:"words,omitempty"`
// }

// func main() {
// 	r1 := Response{
// 		1,
// 		[]string{"solo", "mountains", "green-lands"},
// 	}
// 	j, err := json.Marshal(r1)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(-1)
// 	}
// 	fmt.Println(string(j))

// 	var r2 Response

// 	err = json.Unmarshal(j, &r2)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(-1)
// 	}
// 	fmt.Println(r1)
// 	fmt.Println(r2)
// }

// EXAMPLE 18

// type Album struct {
// 	title string
// }

// func changeTitle(album *Album) {
// 	album.title = "UNDEFINED"
// }

// func main() {
// 	a1 := Album{"The Green Fields"}
// 	fmt.Println(a1)
// 	changeTitle(&a1)
// 	fmt.Println(a1)
// }

// EXAMPLE 17

// VALID
// func main() {
// 	m := map[int]*int{}
// 	a := 10
// 	m[10] = &a
// 	fmt.Println(m[10])
// }

// EXAMPLE 16

// func main() {
// 	type Album struct {
// 		title string
// 	}

// 	a := Album{"The Black Album"}
// 	b := &Album{"The Green Album"}

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	a.title = "Black Changed Changed!"
// 	b.title = "Hey Green Grass Changed"
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// EXAMPLE 15

// type Employee struct {
// 	Name  string
// 	Email string
// 	Age   int
// 	Boss  *Employee
// 	Hired time.Time
// }

// func main() {
// 	emps := map[string]*Employee{}

// 	emps["matt"] = &Employee{
// 		Name:  "matt",
// 		Email: "holiday@matt.com",
// 		Age:   53,
// 		Boss:  nil,
// 		Hired: time.Now(),
// 	}

// 	emps["ashiq"] = &Employee{
// 		Name:  "ashiq",
// 		Email: "a@a.com",
// 		Age:   24,
// 		Boss:  emps["matt"],
// 		Hired: time.Now(),
// 	}

// 	emps["ashiq"].Age += 10
// 	fmt.Println(*emps["matt"])
// 	fmt.Println(*emps["ashiq"])
// }

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
