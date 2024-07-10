package main

import (
	"fmt"
	"go-ds/consistent_hash"
)

func main() {
	hr := consistent_hash.NewHashRing(100000)
	hr.AddNode("Node11")
	hr.AddNode("No123de12")
	hr.AddNode("Node13")
	hr.AddNode("Node14")
	hr.AddNode("#!$#@%")
	cnt := make(map[string]int, 4)
	for i := 0; i < 2000000; i++ {
		node := hr.GetNode(fmt.Sprintf("key%d", i))
		cnt[node]++
	}
	fmt.Println(cnt)
}
