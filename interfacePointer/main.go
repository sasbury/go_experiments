package main

import "fmt"

type Hello interface {
	SayHello()
}

type MyPointerHello struct {
	Name string
}

func (self *MyPointerHello) SayHello() {
	fmt.Printf("Hello %v\n", self.Name)
}

type MyStackHello struct {
	Name string
}

func (self MyStackHello) SayHello() {
	fmt.Printf("Hello %v\n", self.Name)
}

func (self *MyStackHello) SayHello() {
	fmt.Printf("Hello %v\n", self.Name)
}

func TellThemHello(hello Hello) {
	hello.SayHello()
}

func main() {

	var helloOnStack MyStackHello
	var helloPointer = new(MyPointerHello)

	helloOnStack.Name = "From The Stack"
	helloPointer.Name = "From A Pointer"

	TellThemHello(helloOnStack)
	TellThemHello(helloPointer)
}
