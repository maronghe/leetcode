package s

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(IsValid("()"))
	slice := [10]int{}
	s1 := slice[2:3]
	fmt.Println(len(s1)) // 1
	fmt.Println(cap(s1)) // 8

	s := []int{1, 2}
	s = append(s, 3, 4, 5)

	// step 1  3 * 2 = 6
	// 5 > 4 = 5
	// 5 * 8 = 40 <= 48
	// 48 / 8 = 6
	fmt.Println(len(s))   // 5
	fmt.Println(cap(s))   // 6
	fmt.Printf("%p\n", s) // 0xc00001c180
	s = append(s, 6)
	fmt.Printf("%p\n", s) // 0xc00001c180

	fmt.Println(len(s)) // 6
	fmt.Println(cap(s)) // 6

	a := []string{"my", "name", "is"} // 64位 16字节
	a = append(a, "Logan")
	fmt.Println(len(a)) // 4
	fmt.Println(cap(a)) // 4

	b := make([]string, 1024)
	b = append(b, "123")
	fmt.Println(cap(b))
	fmt.Println(1024 * 1.25)

	fmt.Println(ValidPalindrome("abca"))

}
