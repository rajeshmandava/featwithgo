package coretypes

import (
	"fmt"
	"os"
)

func TestCoreTypes() string {
	return "This is test package function calling in Go\n"
}

func DefineArray() [10]int {
	var arr [10]int
	return arr
}

func CompArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0,0,0,0,0}
	arr4 := [5]int{0,0,0,0,9}

	return arr1 == arr2, arr1 == arr3, arr1 ==arr4
}

func CompArrays2() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9:0}
	arr3 := [10]int{1,9: 10, 4: 5}
	return arr1 == arr2, arr1 == arr3, arr3
}

func Message() string {
	var m string = ""
	arr := [4]int{1,2,3,4}
	for i:=0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m +=fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}

func FillArray(arr [10]int) [10]int {
	for i:=0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func OpArray(arr [10]int) [10]int {
	for i :=0; i < len(arr); i++ {
		arr[i] = arr[i]* arr[i]
	}
	return arr
}

func GetPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}

	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func FindLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
	
		}
	}
	return longest
}

var users = map[string]string{
	"305" : "Sue",
	"204" : "Bob",
	"631" : "Jake",
	"073" : "Tracy",
}

func FindName(key string) (string,bool){
	name, exists := users[key]
	return name, exists
}
