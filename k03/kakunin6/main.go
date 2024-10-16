package main

import "fmt"

func PanicRun() {
	panic("ERROR")
}

func main() {
	fmt.Println("start")
	PanicRun()
	fmt.Println("end")
}
