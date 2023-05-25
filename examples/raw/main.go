package main

import (
	"fmt"
	"github.com/dan-and-dna/sensitive"
)

func main() {
	tree := sensitive.NewTrie()

	tree.Add("习近平", "习大大")
	//fmt.Println(tree.Replace("你好吗 我支持习大大， 他的名字叫做习近平", '*'))
	//fmt.Println(tree.Filter("你好吗 我支持习大大， 他的名字叫做习近平"))

	left, word := tree.Hide("你好吗 我支持习大大， 他的名字叫做习近平")
	fmt.Println(left == "你好吗 我支持<<||#|0|-|=>>， 他的名字叫做<<||#|1|-|=>>")
	fmt.Println(word)
	for k, v := range word {
		fmt.Println(k)
		fmt.Println(v)
	}
}
