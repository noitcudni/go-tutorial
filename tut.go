package main

import (
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(arg int) (x, y int) {
	x = arg * 4 / 9
	y = arg - x
	return
}

func short_var() int {
	k := 3
	return k
}

func conv() {
	x, y := 1, 4
	f := math.Sqrt(float64(x*x + y*y))
	fmt.Println("sqrt: ", f)
	fmt.Println("int(f): ", int(f))
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func my_for_loop() {
	for i := 0; i < 10; i++ {
		fmt.Println("looping..", i)
	}
}

func my_while_loop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

///////////////// BEGIN: my sqrt /////////////////

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "Dude, you can be passing me " + strconv.FormatFloat(float64(e), 'f', 6, 64)
}

func my_sqrt(x float64) (float64, error) {
	if x < 0 {
		//return my_sqrt(-x) + "i", x
		return 0, ErrNegativeSqrt(x)
	}
	//return fmt.Sprint(math.Sqrt(x)), nil
	return math.Sqrt(x), nil
}

///////////////// END: my sqrt /////////////////

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func newton_sqrt(x float64) (float64, error) {
	// Todo: Run this.
	z := 1.0

	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z, nil
}

func switch_stmt() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Println("%s", os)
	}
}

func switch_stmt_fallthrough() {
	test := "foo"
	switch test {
	case "foo":
		fmt.Println("foo")
		fallthrough
	case "bar":
		fmt.Println("fall through to bar")
	}
}

func switch_no_condition() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening")
	}
}

func defer_stmt() {
	defer fmt.Println("defer: world")
	fmt.Println("non-defer: world")
}

func defer_stack() {
	for i := 0; i < 9; i++ {
		defer fmt.Println("defer:", i)
	}
	fmt.Println("non-defer stack")
}

func pointer() {
	fmt.Println("----- pointer -----")
	v := 1
	fmt.Println("the address of v: ", &v)
	p := &v
	fmt.Println("the value of p: ", *p)
	*p = 2
	fmt.Println("the new value of v: ", v)
}

type Vertex struct {
	x int
	y int
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(float64(v.x*v.x) + float64(v.y*v.y))
}

func (v *Vertex) Scale(f int) {
	v.x *= f
	v.y *= f
}

func struct_func() {
	fmt.Println("----- struct -----")
	v := Vertex{1, 2}
	fmt.Println(v)
	p := &v
	fmt.Println("access vertex struct via pointer:", p.x)

	v2 := Vertex{y: 4, x: 3}
	fmt.Println("v2", v2)
}

func array() {
	var arry [2]int
	arry[0] = 1
	arry[1] = 2
	fmt.Println(arry)

	arry2 := []int{
		1, 2, 3, 4, 5,
		6, 7,
	}

	fmt.Println("arry2:", arry2)
	fmt.Println(arry2[4:])

	a3 := make([]int, 5)
	fmt.Println("zeroed array: ", a3)
	var a4 [5]int
	fmt.Println("zeroed array: ", a4)

	a5 := make([]int, 0, 3)
	fmt.Println(">>>>> zeroed array with capacity: a5", a5)
	fmt.Println("capacity a5: ", cap(a5))
	fmt.Println("len a5: ", len(a5))
	a5 = append(a5, 1, 2, 3, 4)
	fmt.Println(">>>> a5: ", a5)
	fmt.Println("capacity a5: ", cap(a5))

	a6 := []int{}
	fmt.Println(a6, len(a6), cap(a6))
	fmt.Println(a6 == nil)

	var a7 []int
	fmt.Println(a7 == nil)

}

func range_func() {
	a := []int{1, 2, 3, 4, 5}
	for i, v := range a {
		fmt.Println("index: ", i, " | value: ", v)
	}
	for _, v := range a {
		fmt.Println("dropping index:", v)
	}
	for i := range a {
		fmt.Println("dropping value:", a[i])
	}
}

func map_func() {
	fmt.Println("---- map ----")
	m := make(map[string]Vertex)
	m["bell labs"] = Vertex{2, 2}
	m["sandia"] = Vertex{3, 2}
	fmt.Println(m)

	fmt.Println("---- map literals ----")
	m1 := map[string]Vertex{
		"foo": {0, 0},
		"bar": {1, 1},
	}
	fmt.Println(m1)

	m1["foo"] = Vertex{2, 2}
	fmt.Println(m1)
	delete(m1, "bar")
	fmt.Println(m1)

	v, ok := m1["foo"]
	fmt.Println(v, ok)

}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func closure() {
	fmt.Println("----- closure -----")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func fibonacci() func() int {
	curr := 1
	prev := 0

	return func() int {
		r := curr
		curr = curr + prev
		prev = r
		return r
	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {

	s_array := make([]string, 4)
	for i, x := range ip {
		fmt.Println("raw: ", x)
		s_array[i] = strconv.Itoa(int(x))
	}
	return strings.Join(s_array, ".")
}

func ipaddr_stringers() {
	m := map[string]IPAddr{
		"home": {127, 0, 0, 1},
		"work": {130, 2, 3, 123},
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

//////// begin: Reader //////

func reader_func() {
	r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	cnt := len(b)
	for i := 0; i < cnt; i++ {
		b[i] = 'A'
	}
	return cnt, nil
}

//////// end: Reader   //////

/////// start: rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (rotr rot13Reader) Read(b []byte) (int, error) {

	n, err := rotr.r.Read(b)
	cnt := len(b)
	for i := 0; i < cnt; i++ {
		b[i] = uint8(math.Mod(float64(b[i]-13), float64('A')))
	}

	return n, err
}

/////// end: rot13Reader

/////// start: http handers ///////
type String string
type Struct struct {
	Greeting string
	Punt     string
	Who      string
}

func (h String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "string handler")
}

func (h Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "struct handler")
}

/////// end: http handers ///////

///// Go routine /////
func go_routine() {
	c := make(chan int, 2)
	c <- 1
	c <- 2

	fmt.Println(<-c)
	fmt.Println(<-c)
}

///// end: Go routine /////

///// start : Equivalent Binary Trees
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func check_tree_impl(t *Tree, c chan int) {
	if t.Left != nil {
		check_tree_impl(t.Left, c)
	}

	c <- t.Value

	if t.Right != nil {
		check_tree_impl(t.Right, c)
	}
}

func check_tree(t *Tree, c chan int) {
	check_tree_impl(t, c)
	close(c)
}

func check_trees(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go check_tree(t1, c1)
	go check_tree(t2, c2)

	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2

		if ok1 == false && ok2 == false {
			return true
		} else if ok1 == false && ok2 == true {
			return false
		} else if ok2 == false && ok1 == true {
			return false
		} else {
			if v1 != v2 {
				return true
			}
		}
	} // for
}

///// end : Equivalent Binary Trees

func main() {

	var c, python, java = true, false, "no!"
	fmt.Println(c, python, java)

	fmt.Println("hello, world")
	//fmt.Println("The time is", time.Now())
	fmt.Println("My number is", math.Pi)
	fmt.Println(add(100, 200))
	fmt.Println(swap("hello", "world"))
	var r1, r2 = split(17)
	fmt.Println(r1, r2)

	//Basic Types
	const f string = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)

	conv()
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	my_for_loop()
	my_while_loop()

	//fmt.Println(my_sqrt(2), my_sqrt(-4))
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	fmt.Println(newton_sqrt(25))
	switch_stmt()
	switch_stmt_fallthrough()
	switch_no_condition()
	defer_stmt()
	defer_stack()
	pointer()
	struct_func()
	array()
	range_func()
	map_func()
	closure()
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}

	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	v.Scale(2)
	fmt.Println(v)

	// Stringer Exercise
	ipaddr_stringers()

	fmt.Println(my_sqrt(-2))
	fmt.Println(my_sqrt(9))

	reader_func()
	//http.Handle("/string", String("I'm a frayed knot."))
	//http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	//log.Fatal(http.ListenAndServe("localhost:4000", nil))

	go_routine()

	//Build myself trees
	t01 := &Tree{nil, 1, nil}
	t02 := &Tree{nil, 2, nil}
	t03 := &Tree{t01, 1, t02}
	t04 := &Tree{nil, 5, nil}
	t05 := &Tree{nil, 13, nil}
	t06 := &Tree{t04, 8, t05}
	t07 := &Tree{t03, 3, t06}

	t11 := &Tree{nil, 1, nil}
	t12 := &Tree{nil, 2, nil}
	t13 := &Tree{t11, 1, t12}
	t14 := &Tree{nil, 5, nil}
	t15 := &Tree{t13, 3, t14}
	t16 := &Tree{nil, 13, nil}
	t17 := &Tree{t15, 8, t16}

	fmt.Println("check_trees: ", check_trees(t07, t17))
}
