package main

import (
	"github.com/dan-and-dna/sensitive"
	"log"
)

func main() {
	tree := sensitive.NewTrie()

	tree.Add("我妈妈妈妈妈妈妈妈", "我妈妈妈妈妈妈妈", "我妈妈妈妈妈妈", "我妈妈妈妈", "我妈妈", "我妈妈妈", "我妈妈妈妈妈", "我妈", "妈")

	runes := []rune("你们好，我怕吵醒我妈妈,dd,我妈妈妈妈妈妈妈妈dd妈的你们哈")
	leftStr, pos := tree.AddPlaceholderLongest(string(runes))
	log.Println(pos)
	log.Println(leftStr)
}
