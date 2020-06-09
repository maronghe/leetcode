package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	//	foo := &Foo{Name: "biezhi"}
	//	fmt.Printf("foo: %v address: %p \n", foo, &foo) // 1
	//
	//	modifyNameByPoint(foo)
	//	fmt.Printf("foo: %v address: %p \n", foo, &foo) // 3

	//foo := Foo{Name: "biezhi"}
	//fmt.Printf("foo: %v address: %p \n", foo, &foo) // 1

	//str := nameToUpper(foo)
	//fmt.Printf("foo: %v address: %p str:%s\n", foo, &foo, str) // 3
	var list []*Stu
	list = append(list, &Stu{"maronghe", 123})
	list = append(list, &Stu{"maronghe", 123})
	fmt.Printf("%+v\n", list)
	for _, v := range list {
		fmt.Println(*v)
	}

}

type Stu struct {
	name string
	age  int
}

func nameToUpper(foo Foo) string {
	fmt.Printf("1-nameToUpper foo: %v address: %p \n", foo, &foo) // 2
	foo.Name = strings.ToUpper(foo.Name)
	fmt.Printf("2-nameToUpper foo: %v address: %p \n", foo, &foo) // 2
	return foo.Name
}

type Foo struct {
	Name string
}

func modifyNameByPoint(foo *Foo) {
	fmt.Printf("1-modifyNameByPoint foo: %v address: %p \n", foo, &foo) // 2
	foo.Name = "emmmm " + foo.Name
	fmt.Printf("2-modifyNameByPoint foo: %v address: %p \n", foo, &foo) // 2
}
