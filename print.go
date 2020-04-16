package bencode

import "fmt"

// Print B
func (b *B) Print() {
	b.b.print(0)
}

func (bs *bstruct) print(ident int) {
	switch bs.typ {
	case INT:
		fmt.Print(bs.n, " ")
	case STRING:
		fmt.Print(bs.str, " ")
	case LIST:
		for k := range bs.list {
			for i := 0; i < ident; i++ {
				fmt.Print(" ")
			}
			bs.list[k].value.print(ident + 2)
			fmt.Print("\n")
		}
	case DICT:
		for k := range bs.list {
			for i := 0; i < ident; i++ {
				fmt.Print(" ")
			}
			fmt.Print(*bs.list[k].key, ":")
			bs.list[k].value.print(ident + 2)
			fmt.Print("\n")
		}
	}
}
