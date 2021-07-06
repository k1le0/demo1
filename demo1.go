package main

type Demo1 struct {
}

func New() *Demo1 {
	return &Demo1{}
}

func Print(demo1 *Demo1) string {
	return "Hello World"
}
