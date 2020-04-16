package main

import (
	"fmt"
	"os"
	"github.com/capric98/bencode"
)

func main() {
	tlmc, _ := os.Open("../test/TLMC.torrent")
	b, e := bencode.NewDecoder(tlmc).Decode()
	if e != nil {
		fmt.Println(e)
		return
	}
	b.Print()
}
