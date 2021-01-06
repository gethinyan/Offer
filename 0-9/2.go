package main

import "fmt"

func ReplaceSpace(s string) string {
	//write your code here
	sNew := ""
	for _, str := range s {
		if " " == string(str) {
			sNew += "%20"
		} else {
			sNew += string(str)
		}
	}
	return sNew
}

func main() {
	fmt.Println(ReplaceSpace("We Are Happy."))
}
