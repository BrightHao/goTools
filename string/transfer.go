package main

import "fmt"

func main() {
	s := "中文123nice"

	index := 0
	runeS := []rune(s)
	byteS := []byte(s)

	fmt.Printf("%s index:%d \n", s, index)
	fmt.Printf("==========\n %v \nrune length: %d value:%v, string:%s \n", runeS, len(runeS), runeS[index], string(runeS[index]))
	fmt.Printf("==========\n %v \nbyte length:%d, value:%v, string:%s \n", byteS, len(byteS), byteS[index], string(byteS[index]))

}
