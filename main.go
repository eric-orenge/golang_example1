package main

import "fmt"

type Person struct {
	fname string
	lname string
}

func (pers Person) Speak() {
	fmt.Println("My name is ", pers.fname)
}

type secretAgent struct { //oop inherit but not oop so its composition
	Person
	licenceToKill bool
}

func (s1 secretAgent) Speak() {
	fmt.Println(s1.fname, s1.lname, "Got licence to kill:", s1.licenceToKill)
}

type Human interface { //defines behaviour//polymophism
	Speak()
}

func saySomething(h Human) {
	h.Speak()
}

func main() {
	p1 := Person{
		"eric",
		"orenge",
	}
	s1 := secretAgent{
		Person{"James",
			"Bond"},
		true,
	}
	// fmt.Println(p1)
	saySomething(p1)
	saySomething(s1)

}
