package main

import "fmt"

const (
	a = iota //0
	b        //1
	c        //2
)

const (
	d = iota //0
	e        //1
	f = iota //2
	g        //3
)

const (
	h = iota + 1 //1
	i            //2
	j            //3
	k            //4
)

const (
	l = iota      //0
	m             //1
	n             //2
	o = 1 << iota // iota == 3; o == 8 - move 0001 (1) shift to the left 'iota' bits, in this case iota iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Printf("bit shift 1 three bits to the left: %d\n", 1<<3)

}
