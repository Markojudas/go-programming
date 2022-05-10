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

}
