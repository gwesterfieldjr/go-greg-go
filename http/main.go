package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type shape interface {
	getArea() float64
}

type triange struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func (t triange) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

type logWriter struct{}

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//bs := make([]byte, 99999)
	//res.Body.Read(bs)
	//fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, res.Body)

	s := square{sideLength: 2.0}
	t := triange{base: 2.0, height: 3.0}

	printArea(s)
	printArea(t)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
