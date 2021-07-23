package main

import "fmt"

func main() {
	fmt.Println(validate("00000000000"))    // false
	fmt.Println(validate("86446422799"))    // false
	fmt.Println(validate("86446422784"))    // true
	fmt.Println(validate("864.464.227-84")) // true
	fmt.Println(validate("91720489726"))    // true
	fmt.Println(validate("a1720489726"))    // false
}
