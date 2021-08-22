package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	nonce int
	// trx   []string
	phash []byte
	chash []byte
	trx   []int
}

func createblock(nonce int, trx []int) block {
	b := block{nonce: nonce, trx: trx}
	return b
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b1 := createblock(10, a[:]) //block1

	s := fmt.Sprintf("%v", b1)
	sum := sha256.Sum256([]byte(s))
	b1.chash = sum[:] //hash of current block

	// fmt.Println("BLOCK-1 HASH")
	// fmt.Println(sum)

	fmt.Println("BLOCK-genesis")
	fmt.Println(b1) //, reflect.TypeOf(b1))
	blocks := [4]block{}
	blocks[0] = b1
	for i := 1; i < 4; i++ {
		a := [5]int{10, 20, 30, 40, 50}

		blocks[i] = createblock(blocks[i-1].nonce+1, a[:]) //block2
		blocks[i].phash = blocks[i-1].chash                //append prev hash to current block

		s2 := fmt.Sprintf("%v", blocks[i])
		sum2 := sha256.Sum256([]byte(s2))
		blocks[i].chash = sum2[:]
		// fmt.Println(i+1,"-BLOCK HASH")	//hash of a block
		// fmt.Println(sum2)
		fmt.Println("------------------------------")
		fmt.Println("BLOCK-", i+1)
		fmt.Println(blocks[i])
	}

}
