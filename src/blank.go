package main

import "fmt"

func main() {
	v, _ := hoge()

	fmt.Println(v)
}

func hoge() (string, string) {
	return "hogehoge", "fugafuga"
}