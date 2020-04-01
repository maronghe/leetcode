package main

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	test2()
}

type TI struct { // Stuct Interface
	Data interface{} `json:"data"`
}

type SA struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr"` // diff
}

type SB struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school"` // diff
}

func test1() {

	str := `{"name":"logan","age":18,"addr":"Beijing"}`

	var sa SA
	err := json.Unmarshal([]byte(str), &sa)
	if err != nil {
		panic("err")
	}

	fmt.Printf("sa %v \n", sa)

}

func test2() {

	str := `{
			"data": {"name":"logan","age":18,"addr":"Beijing"}
			}`

	var ti2 TI
	//ti2.Data = &SA{} // #1
	ti2.Data = &SB{} // #2
	err := json.Unmarshal([]byte(str), &ti2)
	if err != nil {
		panic("err")
	}

	fmt.Printf(" it2 data type is %T\n", ti2.Data)

	// type assertions

	ti3, ok := ti2.Data.(*SA)
	if ok {
		fmt.Printf("ti3 is %v \n", ti3)
	} else {
		fmt.Printf("ti3 type assertion error, it isn't type SA \n")
	}

	ti4, ok := ti2.Data.(*SB)
	if ok {
		fmt.Printf("ti4 is %v \n", ti4)
	} else {
		fmt.Printf("ti4 type assertion error, it isn't type SB \n")
	}

	// type assertions
	ti5, ok := ti2.Data.(map[string]interface{})
	if ok {
		fmt.Printf("ti5 is %v \n", ti5)
	} else {
		fmt.Printf("ti5 type assertion error, it isn't type map[string]interface{} \n")
	}

}
