package main

import "fmt"

func main() {

	ch := make(chan int, 5)
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
	hashTable := map[int]int{}
	hashTable2 := make(map[int]int, 2)
	fmt.Println(hashTable)
	fmt.Println(len(hashTable))
	fmt.Println(len(hashTable2))
	hashTable[0] = 1
	hashTable2[0] = 1
	hashTable2[1] = 1
	hashTable2[3] = 1
	hashTable2[4] = 1
	fmt.Println(hashTable)
	fmt.Println(len(hashTable))
	fmt.Println(len(hashTable2))
	hashTable[0] = 1
}
